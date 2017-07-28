package main

import (
	"github.com/astaxie/beego"
	_ "github.com/hsluoyz/gitstar/routers"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionProvider="file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"

	beego.SetLogger("file", `{"filename":"logs/gitstar.log"}`)
	beego.SetLogFuncCall(false)

	beego.Run()
}

