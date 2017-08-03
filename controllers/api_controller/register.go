package api_controller

import (
	"github.com/hsluoyz/gitstar/api"
	"github.com/hsluoyz/gitstar/util"
)

type Response struct {
	Code int
	Msg  string
	Data string
}

func (c *APIController) Register() {
	var resp Response
	username, password := c.Input().Get("username"), c.Input().Get("password")

	msg := api.CheckUserRegister(username, password)
	if msg == "" {
		api.AddUser(username, password)

		util.LogInfo(c.Ctx, "API: [%s] is registered as new user", username)
		resp = Response{Code: 200, Msg: "注册成功", Data: username}
	} else {
		resp = Response{Code: 0, Msg: msg, Data: ""}
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *APIController) Login() {
	var resp Response
	username, password := c.Input().Get("username"), c.Input().Get("password")

	if !api.HasUser(username) {
		resp = Response{Code: 0, Msg: "用户名不存在，请先注册", Data: ""}
	} else if !api.CheckUserPassword(username, password) {
		resp = Response{Code: 0, Msg: "密码错误", Data: ""}
	} else {
		util.LogInfo(c.Ctx, "API: [%s] logged in", username)
		resp = Response{Code: 200, Msg: "登录成功", Data: username}
	}

	c.Data["json"] = resp
	c.ServeJSON()
}
