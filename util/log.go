package util

import (
	"fmt"

	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/hsluoyz/gitstar/api"
)

func LogInfo(ctx *context.Context, f string, v ...interface{}) {
	var ipString string
	clientIP := ctx.Request.Header.Get("x-forwarded-for")
	if clientIP != "" {
		desc := api.GetDescFromIP(clientIP)
		ipString = fmt.Sprintf("(%s: %s) ", clientIP, desc)
	} else {
		ipString = "() "
	}

	logs.Info(ipString + f, v...)
}
