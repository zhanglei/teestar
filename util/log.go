package util

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

func LogInfo(ctx *context.Context, f interface{}, v ...interface{}) {
	logs.Info(f, v...)
}
