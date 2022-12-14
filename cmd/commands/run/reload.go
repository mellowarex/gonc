package run

import (
	"bytes"
	"net/http"
	"time"

	gonLogger "github.com/mellowarex/gonc/logger"
	"github.com/gorilla/websocket"
)

// wsBroker maintains the set of active clients and broadcasts messages to the clients.
type wsBroker struct {
	clients    map[*wsClient]bool // Registered clients.
	broadcast  chan []byte        // Inbound messages from the clients.
	register   chan *wsClient     // Register requests from the clients.
	unregister chan *wsClient     // Unregister requests from clients.
}

func (br *wsBroker) run() {
	for {
		select {
		case client := <-br.register:
			br.clients[client] = true
		case client := <-br.unregister:
			if _, ok := br.clients[client]; ok {
				delete(br.clients, client)
				close(client.send)
			}
		case message := <-br.broadcast:
			for client := range br.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(br.clients, client)
				}
			}
		}
	}
}

// wsClient represents the end-client.
type wsClient struct {
	broker *wsBroker       // The broker.
	conn   *websocket.Conn // The websocket connection.
	send   chan []byte     // Buffered channel of outbound messages.
}

// readPump pumps messages from the websocket connection to the broker.
func (c *wsClient) readPump() {
	defer func() {
		c.broker.unregister <- c
		c.conn.Close()
	}()

	for {
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				gonLogger.Log.Errorf("An error happened when reading from the Websocket client: %v", err)
			}
			break
		}
	}
}

// write writes a message with the given message type and payload.
func (c *wsClient) write(mt int, payload []byte) error {
	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	return c.conn.WriteMessage(mt, payload)
}

// writePump pumps messages from the broker to the websocket connection.
func (c *wsClient) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				// The broker closed the channel.
				c.write(websocket.CloseMessage, []byte{})
				return
			}

			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write([]byte("/n"))
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

var (
	broker        *wsBroker  // The broker.
	reloadAddress = ":12450" // The port on which the reload server will listen to.

	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
)

const (
	writeWait  = 10 * time.Second    // Time allowed to write a message to the peer.
	pongWait   = 60 * time.Second    // Time allowed to read the next pong message from the peer.
	pingPeriod = (pongWait * 9) / 10 // Send pings to peer with this period. Must be less than pongWait.
)

func startReloadServer() {
	broker = &wsBroker{
		broadcast:  make(chan []byte),
		register:   make(chan *wsClient),
		unregister: make(chan *wsClient),
		clients:    make(map[*wsClient]bool),
	}

	go broker.run()
	http.HandleFunc("/reload", func(w http.ResponseWriter, r *http.Request) {
		handleWsRequest(broker, w, r)
	})

	go startServer()
	gonLogger.Log.Infof("Reload server listening at %s", reloadAddress)
}

func startServer() {
	err := http.ListenAndServe(reloadAddress, nil)
	if err != nil {
		gonLogger.Log.Errorf("Failed to start up the Reload server: %v", err)
		return
	}
}

func sendReload(payload string) {
	message := bytes.TrimSpace([]byte(payload))
	broker.broadcast <- message
}

// handleWsRequest handles websocket requests from the peer.
func handleWsRequest(broker *wsBroker, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		gonLogger.Log.Errorf("error while upgrading server connection: %v", err)
		return
	}

	client := &wsClient{
		broker: broker,
		conn:   conn,
		send:   make(chan []byte, 256),
	}
	client.broker.register <- client

	go client.writePump()
	client.readPump()
}
