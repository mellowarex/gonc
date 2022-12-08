package new

import (
	"fmt"
	"os"
	path "path/filepath"
	// "strings"

	"github.com/mellowarex/gonc/cmd/commands"
	// "github.com/beego/bee/cmd/commands/version"
	gonLogger "github.com/mellowarex/gonc/logger"
	// "github.com/beego/bee/logger/colors"
	"github.com/mellowarex/gonc/utils"
)

var gonVersion utils.DocValue

var CmdNew = &commands.Command{
	UsageLine: "new", 
	Use:				"new appname [-v=1.0.0]",
	Args: 			[]string{"[-v=1.0.0]"},
	Short:     "Creates a Gon application",
	Long: `Creates a Gon web application for the given app name in the current directory.
	The web app is a go module project
	The command 'new' creates a folder named appname and generates the following structure:

            ├── main.go
            ├── go.mod
            ├── assets
            │     └── css
            │     └── img
            │     └── js
            ├── config
            │     └── environments
            │           └── development.json
            │           └── production.json
            │     └── application.json
            │     └── cookie.json
            ├── controllers
            │     └── index.go
            ├── log
            │     └── development
            │     └── production
            ├── models
            ├── public
            │     └── 4xx.html
            │     └── 5xx.html
            │     └── favicon.ico
            │     └── robots.txt
            │     └── sitemap.xml
            ├── routers
            │     └── public.go
            ├── tests
            │     └── default_test.go
            └── views
                  └── index.tpl

`,
	PreRun: nil,
	Run:    CreateApp,
}

// var goMod = `module %s

// go %s

// require github.com/mellocraft/gon %s
// require github.com/smartystreets/goconvey v1.6.4
// `
var goMod = `module %s

go %s

replace github.com/mellocraft/gon => ../gon

require github.com/mellocraft/gon v0.0.0-00010101000000-000000000000
`

var reloadJsClient = `function b(a){var c=new WebSocket(a);c.onclose=function(){setTimeout(function(){b(a)},2E3)};c.onmessage=function(){location.reload()}}try{if(window.WebSocket)try{b("ws://localhost:12450/reload")}catch(a){console.error(a)}else console.log("Your browser does not support WebSockets.")}catch(a){console.error("Exception during connecting to Reload:",a)};
`

func init() {
	CmdNew.Flag.Var(&gonVersion, "v", "set gon version to run the app, version must be installed to work")
	commands.AvailableCommands = append(commands.AvailableCommands, CmdNew)
}

func CreateApp(cmd *commands.Command, args []string) int {
	output := cmd.StdOut()
	if len(args) == 0 {
		gonLogger.Log.Fatal("Argument appname is missing")
		return 2
	}

	if len(args) >= 2 {
		err := cmd.Flag.Parse(args[1:])
		if err != nil {
			gonLogger.Log.Fatal("Parse args err " + err.Error())
		}
	}
	var appPath string
	var packPath string
	var appName string
	// var err error

	// 	beeLogger.Log.Info("generate new project support go modules.")
	// 	appPath = path.Join(utils.GetBeeWorkPath(), args[0])
	// 	packPath = args[0]
	if gonVersion.String() == `` {
		gonVersion.Set(`v1.0.0`)
	}
	appPath = path.Join(utils.GetWorkPath(), args[0])
	packPath = args[0]
	appName = args[0]

	if utils.IsExist(appPath) {
		gonLogger.Log.Errorf("Application '%s' already exists", appPath)
		gonLogger.Log.Warn("Do you want to overwrite it? [Yes|No] ")
		if !utils.AskForConfirmation() {
			os.Exit(2)
		}
	}

	gonLogger.Log.Info("Creating application...")

	// If it is the current directory, select the current folder name to package path
	if packPath == "." {
		packPath = path.Base(appPath)
	}

	os.MkdirAll(appPath, 0755)
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "go.mod"), "\x1b[0m")
	utils.WriteToFile(path.Join(appPath, "go.mod"), fmt.Sprintf(goMod, packPath, utils.GetGoVersionSkipMinor()))
	// utils.WriteToFile(path.Join(appPath, "go.mod"), fmt.Sprintf(goMod, packPath, utils.GetGoVersionSkipMinor(), gonVersion.String()))

	// create assets folder: css, js, img
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "assets")+string(path.Separator), "\x1b[0m")
	os.Mkdir(path.Join(appPath, "assets"), 0755)
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "assets", "js")+string(path.Separator), "\x1b[0m")
	os.Mkdir(path.Join(appPath, "assets", "css"), 0755)
	
	os.Mkdir(path.Join(appPath, "assets", "js"), 0755)
	// utils.WriteToFile(path.Join(appPath, "static", "js", "reload.min.js"), reloadJsClient)
	
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "assets", "css")+string(path.Separator), "\x1b[0m")
	os.Mkdir(path.Join(appPath, "assets", "img"), 0755)
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "assets", "img")+string(path.Separator), "\x1b[0m")


	// create config
	// create config folder
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "config", "environments")+string(path.Separator), "\x1b[0m")
	os.Mkdir(path.Join(appPath, "config"), 0755)
	// create environment folder
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "config", "environments")+string(path.Separator), "\x1b[0m")
	os.Mkdir(path.Join(appPath, "config", "environments"), 0755)
	// create development.json
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "config", "environments", "development.json"), "\x1b[0m")
	utils.WriteToFile(path.Join(appPath, "config", "environments", "development.json"), developmentJson)
	// create production.json
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "config", "environments", "production.json"), "\x1b[0m")
	utils.WriteToFile(path.Join(appPath, "config", "environments", "production.json"), developmentJson)
	// create application.json
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "config", "application.json"), "\x1b[0m")
	utils.WriteToFile(path.Join(appPath, "config", "application.json"), fmt.Sprintf(applicationJson, appName))
	// create application.json
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "config", "cookie.json"), "\x1b[0m")
	utils.WriteToFile(path.Join(appPath, "config", "cookie.json"), cookieJson)



	// controllers
	// create controller folder
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "controllers")+string(path.Separator), "\x1b[0m")
	os.Mkdir(path.Join(appPath, "controllers"), 0755)
	// create controller/index.go
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "controllers", "index.go"), "\x1b[0m")
	utils.WriteToFile(path.Join(appPath, "controllers", "index.go"), indexController)



	// log
	// create log
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "log")+string(path.Separator), "\x1b[0m")
	os.Mkdir(path.Join(appPath, "log"), 0755)
	// create log/development
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "log", "development")+string(path.Separator), "\x1b[0m")
	os.Mkdir(path.Join(appPath, "log", "development"), 0755)
	// create log/production
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "log", "production")+string(path.Separator), "\x1b[0m")
	os.Mkdir(path.Join(appPath, "log", "production"), 0755)



	// models
	// create empty models
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "models")+string(path.Separator), "\x1b[0m")
	os.Mkdir(path.Join(appPath, "models"), 0755)


	// public
	// create public/
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "public")+string(path.Separator), "\x1b[0m")
	os.Mkdir(path.Join(appPath, "public"), 0755)
	// create public/favicon.ico
	// fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "public", "5xx.html"), "\x1b[0m")
	// utils.WriteToFile(path.Join(appPath, "public", "5xx.html"), xx5Html)
	// create public/robots.txt
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "public", "robots.txt"), "\x1b[0m")
	utils.WriteToFile(path.Join(appPath, "public", "robots.txt"), robots)
	// create public/sitemap.xml
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "public", "sitemap.xml"), "\x1b[0m")
	utils.WriteToFile(path.Join(appPath, "public", "sitemap.xml"), sitemap)


	// create routers
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "routers")+string(path.Separator), "\x1b[0m")
	os.Mkdir(path.Join(appPath, "routers"), 0755)
	// create routers/routers.go
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "routers", "routers"), "\x1b[0m")
	utils.WriteToFile(path.Join(appPath, "routers", "routers.go"), fmt.Sprintf(routers, appName))


	// create tests
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "tests")+string(path.Separator), "\x1b[0m")
	os.Mkdir(path.Join(appPath, "tests"), 0755)
	// create tests/default_test.go
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "tests", "default_test.go"), "\x1b[0m")
	utils.WriteToFile(path.Join(appPath, "tests","default_test.go"), test)
	// utils.WriteToFile(path.Join(appPath, "tests","default_test.go"), fmt.Sprintf(test, appName))



	// create views
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "views")+string(path.Separator), "\x1b[0m")
	os.Mkdir(path.Join(appPath, "views"), 0755)
	// create views/index.tpl
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "views", "index.tpl"), "\x1b[0m")
	utils.WriteToFile(path.Join(appPath, "views", "index.tpl"), indextpl)
	// create view/errors
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "views", "errors")+string(path.Separator), "\x1b[0m")
	os.Mkdir(path.Join(appPath, "views", "errors"), 0755)
	// create view/errors/4xx.tpl
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "public", "4xx.tpl"), "\x1b[0m")
	utils.WriteToFile(path.Join(appPath, "views", "errors", "4xx.tpl"), xx4Html)
	// create view/errors/5xx.tpl
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "public", "5xx.tpl"), "\x1b[0m")
	utils.WriteToFile(path.Join(appPath, "views", "errors", "5xx.tpl"), xx5Html)

	// create main.go
	fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "main.go"), "\x1b[0m")
	utils.WriteToFile(path.Join(appPath, "main.go"), fmt.Sprintf(maingo, appName))

	gonLogger.Log.Success("New application successfully created!")
	return 0
}
