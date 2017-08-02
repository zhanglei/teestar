package main

import (
	"github.com/astaxie/beego"
	_ "github.com/hsluoyz/gitstar/routers"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionProvider="file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600 * 24 * 365

	beego.SetLogger("file", `{"filename":"logs/gitstar.log","maxdays":99999}`)
	beego.SetLogFuncCall(false)

	beego.Run()
}

