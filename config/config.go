package config

 import (
 	"os"
 	"io/ioutil"
 	"encoding/json"
 	"path/filepath"

 	gonLogger "github.com/mellowarex/gonc/logger"
 )

 var Conf = struct {
	WatchExts          []string  `json:"watch_ext"`
	WatchExtsStatic    []string  `json:"watch_ext_static" `
	GoInstall          bool      `json:"go_install"` // Indicates whether execute "go install" before "go build".
	DirStruct          dirStruct `json:"dir_structure"`
	CmdArgs            []string  `json:"cmd_args"`
	Envs               []string
	// Bale               bale
	Database           database
	EnableReload       bool              `json:"enable_reload"`
	EnableNotification bool              `json:"enable_notification"`
	Scripts            map[string]string `json:"scripts"`
}{
	WatchExts:       []string{".go"},
	WatchExtsStatic: []string{".html", ".tpl", ".js", ".css"},
	GoInstall:       true,
	DirStruct: dirStruct{
		Others: []string{},
	},
	CmdArgs: []string{},
	Envs:    []string{},
	// Bale: bale{
	// 	Dirs:   []string{},
	// 	IngExt: []string{},
	// },
	Database: database{
		Driver: "postgres",
	},
	EnableNotification: true,
	Scripts:            map[string]string{},
}

// dirStruct describes the application's directory structure
type dirStruct struct {
	WatchAll    bool `json:"watch_all"`
	Controllers string
	Models      string
	Others      []string // Other directories
}

// bale
type bale struct {
	Import string
	Dirs   []string
	IngExt []string `json:"ignore_ext"`
}

// database holds the database connection information
type database struct {
	Driver string
	Conn   string
	Dir    string
}

// LoadConfig loads gon tool configuration
// looks for Hunterfile or gon.json in the current path-webapp path
// if config file not found default to default configuration
func LoadConfig() {
	currentPath, err := os.Getwd()
	if err != nil {
		gonLogger.Log.Error(err.Error())
	}

	dir, err := os.Open(currentPath)
	if err != nil {
		gonLogger.Log.Error(err.Error())
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		gonLogger.Log.Error(err.Error())
	}

	for _, file := range files {
		if file.Name() == "gonc.json" {
			err = parseJSON(filepath.Join(currentPath, file.Name()), &Conf)
			if err != nil {
				gonLogger.Log.Errorf("Failed to parse gonc.json: %s, using fallback config", err)
			}
			break
		}
	}

	// Set variables
	if len(Conf.DirStruct.Controllers) == 0 {
		Conf.DirStruct.Controllers = "controllers"
	}

	if len(Conf.DirStruct.Models) == 0 {
		Conf.DirStruct.Models = "models"
	}
}

func parseJSON(path string, v interface{}) error {
	var (
		data []byte
		err  error
	)
	data, err = ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, v)
	return err
}