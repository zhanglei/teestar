package routers

import (
	"github.com/astaxie/beego/context"
	"github.com/weilaihui/teestar/api"
	"github.com/weilaihui/teestar/util"
)

var FilterIP = func(ctx *context.Context) {
	clientIP := ctx.Request.Header.Get("x-forwarded-for")

	if clientIP != "" && !api.IsMainland(clientIP) {
		util.LogInfo(ctx, "request is denied to access")

		w := ctx.ResponseWriter
		w.WriteHeader(403)
		w.Write([]byte("您所在的地区无法访问本网站, 如果需要添加白名单授权访问，请联系管理员，加入QQ群：646373152\n"))
	}
}
