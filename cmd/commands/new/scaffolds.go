package new

var maingo = `package main

import (
	_ "%s/routers"
	"github.com/mellowarex/gon"
)

func main() {
	gon.Run()
}

`

var developmentJson = `{
  "Listen": {
    "ServerTimeOut": 15,
    "ListenTCP4": false,
    "Domains": [],
    "EnableHTTP": true,
    "HTTPAddr": "127.0.0.1",
    "HTTPPort": 8000,
    "TLSCacheDir": ".",
    "AutoTLS": false,
    "EnableHTTPS": false,
    "EnableMutualHTTPS": false,
    "HTTPSAddr": "127.0.0.1",
    "HTTPSPort": 10443,
    "HTTPSCertFile": "",
    "HTTPSKeyFile": "",
    "TrustCaFile": "",
    "ClientAuth": 0
  },
  "WebConfig": {
    "FlashName": "GON_FLASH",
    "FlashSeparator": "GONFLASH",
    "EnableXSRF": true,
    "XSRFKey": "gonxsrf",
    "XSRFExpire": 0,
    "Session": {
      "SessionOn": true,
      "SessionProvider": "cookie",
      "SessionName": "Acuity",
      "SessionGCMaxLifetime": 315360000,
      "SessionProviderConfig": "cookie.json",
      "SessionCookieLifeTime": 315360000 ,
      "SessionAutoSetCookie": true,
      "SessionDomain": "",
      "SessionDisableHTTPOnly": false,
      "SessionEnableSidInHTTPHeader": false,
      "SessionNameInHTTPHeader": "Gonsessionid",
      "SessionEnableSidInURLQuery": false
    }
  },
  "Log": {
    "DateLog": true,
    "AccessLogs": true,
    "EnableStaticLogs": true,
    "AccessLogsFormat": "APACHE_FORMAT",
    "LMail": {
      "SendMail": false,
      "Env": "development",
      "Username": "",
      "Password": "",
      "Host": "",
      "Subject": "",
      "FromAddress": "",
      "SendTo": [""],
      "Level": 2,
      "HttpStatus": {
        "2xx": false,
        "4xx": false,
        "5xx": false
      }
    }
  }
}`

var applicationJson = `{
  "app": "%s",
  "envmode": "development",
  "server": "nen",
  "recoverPanic": true,
  "copyRequestBody": false,
  "enableGzip": false,
  "maxMemory": 1048576,
  "maxUploadSize": 5242880,

  "directoryIndex": false,
  "staticDir": {"/static":  "assets"},
  "staticExtensionsToGzip": [".css",".js"],
  "staticCacheFileSize": 102400,
  "staticCacheFileNum": 1000,
  "templateLeft": "{{",
  "templateRight": "}}",
  "viewsPath": "views"
}
`

var cookieJson = `{
	"securityKey": "Specialist0001",
	"blockKey": "WaleIsland000001",
	"securityName": "GingFreecsspecialist",
	"cookieName": "Acuity",
	"secure": false,
	"maxage": 315360000
}
`

var indexController = `package controllers

import (
	"github.com/mellocraft/gon/ctrl"
)

type Index struct {
	ctrl.Controller
}

func (this *Index) BeforeAction() {
	this.ActionName = "index"
}

func (this *Index) Get() {
	
}
`

var xx4Html = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <title>{{.Title}}</title>
    <style type="text/css">
      * {
        margin:0;
        padding:0;
      }

      body {
        background-color:#EFEFEF;
        font: .9em "Lucida Sans Unicode", "Lucida Grande", sans-serif;
      }

      #wrapper{
        width:600px;
        margin:40px auto 0;
        text-align:center;
        -moz-box-shadow: 5px 5px 10px rgba(0,0,0,0.3);
        -webkit-box-shadow: 5px 5px 10px rgba(0,0,0,0.3);
        box-shadow: 5px 5px 10px rgba(0,0,0,0.3);
      }

      #wrapper h1{
        color:#FFF;
        text-align:center;
        margin-bottom:20px;
      }

      #wrapper a{
        display:block;
        font-size:.9em;
        padding-top:20px;
        color:#FFF;
        text-decoration:none;
        text-align:center;
      }

      #container {
        width:600px;
        padding-bottom:15px;
        background-color:#FFFFFF;
      }

      .navtop{
        height:40px;
        background-color:#24B2EB;
        padding:13px;
      }

      .content {
        padding:10px 10px 25px;
        background: #FFFFFF;
        margin:;
        color:#333;
      }

      a.button{
        color:white;
        padding:15px 20px;
        text-shadow:1px 1px 0 #00A5FF;
        font-weight:bold;
        text-align:center;
        border:1px solid #24B2EB;
        margin:0px 200px;
        clear:both;
        background-color: #24B2EB;
        border-radius:100px;
        -moz-border-radius:100px;
        -webkit-border-radius:100px;
      }

      a.button:hover{
        text-decoration:none;
        background-color: #24B2EB;
      }

    </style>
  </head>
  <body>
    <div id="wrapper">
      <div id="container">
        <div class="navtop">
          <h1>{{.Title}}</h1>
        </div>
        <div id="content">
          {{.Content}}
          <a href="/" title="Home" class="button">Go Home</a><br />

          <br>Powered by Gon
        </div>
      </div>
    </div>
  </body>
</html>
`
var xx5Html = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <title>{{.Title}}</title>
    <style type="text/css">
      * {
        margin:0;
        padding:0;
      }

      body {
        background-color:#EFEFEF;
        font: .9em "Lucida Sans Unicode", "Lucida Grande", sans-serif;
      }

      #wrapper{
        width:600px;
        margin:40px auto 0;
        text-align:center;
        -moz-box-shadow: 5px 5px 10px rgba(0,0,0,0.3);
        -webkit-box-shadow: 5px 5px 10px rgba(0,0,0,0.3);
        box-shadow: 5px 5px 10px rgba(0,0,0,0.3);
      }

      #wrapper h1{
        color:#FFF;
        text-align:center;
        margin-bottom:20px;
      }

      #wrapper a{
        display:block;
        font-size:.9em;
        padding-top:20px;
        color:#FFF;
        text-decoration:none;
        text-align:center;
      }

      #container {
        width:600px;
        padding-bottom:15px;
        background-color:#FFFFFF;
      }

      .navtop{
        height:40px;
        background-color:#24B2EB;
        padding:13px;
      }

      .content {
        padding:10px 10px 25px;
        background: #FFFFFF;
        margin:;
        color:#333;
      }

      a.button{
        color:white;
        padding:15px 20px;
        text-shadow:1px 1px 0 #00A5FF;
        font-weight:bold;
        text-align:center;
        border:1px solid #24B2EB;
        margin:0px 200px;
        clear:both;
        background-color: #24B2EB;
        border-radius:100px;
        -moz-border-radius:100px;
        -webkit-border-radius:100px;
      }

      a.button:hover{
        text-decoration:none;
        background-color: #24B2EB;
      }

    </style>
  </head>
  <body>
    <div id="wrapper">
      <div id="container">
        <div class="navtop">
          <h1>{{.Title}}</h1>
        </div>
        <div id="content">
          {{.Content}}
          <a href="/" title="Home" class="button">Go Home</a><br />

          <br>Gon
        </div>
      </div>
    </div>
  </body>
</html>
`
var robots = ``
var sitemap = ``

var routers = `package routers

import (
	"github.com/mellocraft/gon"
	"%s/controllers"
)

var (
	Mux = gon.Mux
)

func init() {
	Mux = gon.NewMux()

	Mux.Route("/", &controllers.Index{})
}
`

var test = "package test"

// var test = `package test

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"runtime"
// 	"path/filepath"
// 	_ "%s/routers"

// 	"github.com/mellocraft/gon"
// 	. "github.com/smartystreets/goconvey/convey"
// )

// func init() {
// 	_, file, _, _ := runtime.Caller(0)
// 	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
// 	gon.TestGonInit(apppath)
// }


// // TestBeego is a sample to run an endpoint test
// func TestGon(t *testing.T) {
// 	r, _ := http.NewRequest("GET", "/", nil)
// 	w := httptest.NewRecorder()
// 	gon.GonApp.Handlers.ServeHTTP(w, r)

// 	gon.Trace("testing", "TestGon", "Code[%d]\n%s", w.Code, w.Body.String())

// 	Convey("Subject: Test Station Endpoint\n", t, func() {
// 	        Convey("Status Code Should Be 200", func() {
// 	                So(w.Code, ShouldEqual, 200)
// 	        })
// 	        Convey("The Result Should Not Be Empty", func() {
// 	                So(w.Body.Len(), ShouldBeGreaterThan, 0)
// 	        })
// 	})
// }

// `



var indextpl = `<!DOCTYPE html>
<html>
  <head>
    <meta charset='utf-8'/>
    <meta http-equiv='X-UA-Compatible' content='IE=edge'/>
    <meta name='viewport' content='width=device-width, initial-scale=1'/>
    <meta name='keywords' content='Web Design, Web Design Company, Web Development, Website Design'/>
    <meta name='description' content='<?= $description ?>'/>
    <meta name='geo.region' content='KE-300'/>
    <meta name='geo.placename' content='Mombasa'/>
    <meta name='geo.position' content='-4.050000;39.600000'/>
    <meta name='ICBM' content='-4.050000;39.600000'/>
    <title>{{.Title}}</title>
    <link rel='shortcut icon' href='logo.png'/>
    <link rel='stylesheet' href='/static/css/material-icons.css?v=11'/>
    <link rel='stylesheet' href='/static/css/select2.min.css'/>
    <link rel='stylesheet' href='/static/css/wbbtheme.css'/>
    <link rel='stylesheet' href='/static/css/crimson.css?v=2ddd32ed3'/>
    <link rel='stylesheet' href='/static/css/crimson-darks.css'/>
    <!--[if lt IE 9]>
    <script src="https://oss.maxcdn.com/html5shiv/3.7.2/html5shiv.min.js"></script>
    <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
    <![endif]-->
    <script src="/static/js/loader.js"></script>
    <script src="/static/js/InputFilter.js"></script>
    <script src="/static/js/FormValidation.js?v=1e"></script>
    <script src="/static/js/jquery-3.5.1.min.js"></script>
    <script src="/static/js/bootstrap.bundle.min.js"></script>
  </head>
  <body>
  <script src="/static/js/material.min.js"></script>
  <script src="/static/js/jquery.wysibb.js?v=3dd"></script>
  <script src="/static/js/select2.full.min.js"></script>
  <script src="/static/js/master.js?v=33dsdd"></script>
  <script src="/static/js/papers.js?V=1Wqw"></script>
  <script src="/static/js/user.js?v=1221qE31er323"></script>
</body>
</html>`