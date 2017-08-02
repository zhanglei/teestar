package util

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

func LogInfo(ctx *context.Context, f string, v ...interface{}) {
	logs.Info("(" + ctx.Request.RemoteAddr + ") " + f, v...)
}
