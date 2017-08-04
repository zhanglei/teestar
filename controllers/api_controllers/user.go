package api_controllers

import (
	"github.com/astaxie/beego"

	"github.com/hsluoyz/gitstar/api"
	"github.com/hsluoyz/gitstar/util"
)

type Response struct {
	Code int
	Msg  string
	Data string
}

// User API
type UserController struct {
	beego.Controller
}

// @Title Register
// @Description register a new user
// @Param   username     formData    string  true        "The username to register"
// @Param   password     formData    string  true        "The password"
// @Success 200 {object} controllers.api_controller.Response The response object
// @router /register [post]
func (c *UserController) Register() {
	var resp Response
	username, password := c.Input().Get("username"), c.Input().Get("password")

	msg := api.CheckUserRegister(username, password)
	if msg != "" {
		resp = Response{Code: 0, Msg: msg, Data: ""}
	} else {
		api.AddUser(username, password)

		util.LogInfo(c.Ctx, "API: [%s] is registered as new user", username)
		resp = Response{Code: 200, Msg: "注册成功", Data: username}
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

// @Title Login
// @Description login as a user
// @Param   username     formData    string  true        "The username to login"
// @Param   password     formData    string  true        "The password"
// @Success 200 {object} controllers.api_controller.Response The response object
// @router /login [post]
func (c *UserController) Login() {
	var resp Response
	username, password := c.Input().Get("username"), c.Input().Get("password")

	msg := api.CheckUserLogin(username, password)
	if msg != "" {
		resp = Response{Code: 0, Msg: msg, Data: ""}
	} else {
		util.LogInfo(c.Ctx, "API: [%s] logged in", username)
		resp = Response{Code: 200, Msg: "登录成功", Data: username}
	}

	c.Data["json"] = resp
	c.ServeJSON()
}
