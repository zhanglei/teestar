package api_controller

import (
	"github.com/hsluoyz/gitstar/api"
	"github.com/hsluoyz/gitstar/util"
)

func (c *APIController) Login() {
	msg := "ok"
	username, password := c.Input().Get("username"), c.Input().Get("password")

	if !api.HasUser(username) {
		msg = "用户名不存在，请先注册"
	}

	ok := api.CheckUserPassword(username, password)
	if ok {
		util.LogInfo(c.Ctx, "API: [%s] logged in", username)
	} else {
		msg = "密码错误"
	}

	c.Data["json"] = msg
	c.ServeJSON()
}
