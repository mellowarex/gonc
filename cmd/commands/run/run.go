package run

import (
	"os"
	"io/ioutil"
	// "strings"
	"runtime"
	path "path/filepath"

	gonLogger "github.com/mellowarex/gonc/logger"
	"github.com/mellowarex/gonc/cmd/commands"
	"github.com/mellowarex/gonc/config"
	"github.com/mellowarex/gonc/utils"
)

var CmdRun = &commands.Command{
	UsageLine: 	"run",
	Use: 			 	"run appname",
	Args:				[]string{},
	Short:			"Run the application by starting a server",
	Long:				`Run the web application by starting the server.
	Watches the application filesystem for changes and recompile/restart the server.`,
	PreRun:			func(cmd *commands.Command, args []string) {},
	Run:				RunApp,
}

var (
	// WebApp path
	currentPath string
	// WebApp name
	appname string
	// current user workspace
	currentGoPath	string
	// Channel to signal an Exit
	exit chan bool
	// Pass through to -tags arg of "go build"
	buildTags string
	// Extra args to run application
	runargs string
	// Pass through to -ldflags arg of "go build"
	buildLDFlags string
)

var started = make(chan bool)

func init() {
	exit = make(chan bool)
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

// RunApp locates files to watch, and start gon app
func RunApp(cmd *commands.Command, args []string) int {
	// default app path is current working directory
	appPath, _ := os.Getwd()

	appname = path.Base(appPath)
	currentGoPath = appPath

	gonLogger.Log.Infof("Gon WebApp: %s", appname)
	gonLogger.Log.Debugf("Current path: %s", utils.FILE(), utils.LINE(), appPath)

	var paths []string
	readAppDirectories(appPath, &paths)

	// watch current directory and ignore non-go files
	// for _, p := range config.Conf.DirStruct.Others {
	// 	paths = append(paths, strings.Replace(p, "$GOPATH", currentGoPath), -1)
	// }

	files := []string{}

	// Start the reload server
	if config.Conf.EnableReload {
		startReloadServer()
	}

	NewWatcher(paths, files, false)
	AutoBuild(files, false)

	for {
		<-exit
		runtime.Goexit()
	}
}

func readAppDirectories(directory string, paths *[]string) {
	fileInfos, err := ioutil.ReadDir(directory)
	if err != nil {
		return
	}
	useDirectory := false
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() && fileInfo.Name()[0] != '.' {
			readAppDirectories(directory+"/"+fileInfo.Name(), paths)
			continue
		}

		if useDirectory {
			continue
		}

		if path.Ext(fileInfo.Name()) == ".go" || (ifStaticFile(fileInfo.Name()) && config.Conf.EnableReload) {
			*paths = append(*paths, directory)
			useDirectory = true
		}
	}
}