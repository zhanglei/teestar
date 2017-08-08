package api_controllers

import (
	"github.com/hsluoyz/gitstar/api"
	"github.com/hsluoyz/gitstar/controllers"
	"github.com/hsluoyz/gitstar/util"
)

type Response struct {
	Code int
	Msg  string
	Data string
}

// User API
type UserController struct {
	controllers.BaseController
}

// @Title Register
// @Description register a new user
// @Param   username     formData    string  true        "The username to register"
// @Param   password     formData    string  true        "The password"
// @Success 200 {object} controllers.api_controller.Response The Response object
// @router /register [post]
func (c *UserController) Register() {
	var resp Response

	if c.GetSessionUser() != "" {
		resp = Response{Code: 0, Msg: "请先注销当前用户后再注册", Data: c.GetSessionUser()}
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}

	user, password := c.Input().Get("username"), c.Input().Get("password")

	msg := api.CheckUserRegister(user, password)
	if msg != "" {
		resp = Response{Code: 0, Msg: msg, Data: ""}
	} else {
		api.AddUser(user, password)

		c.SetSessionUser(user)

		util.LogInfo(c.Ctx, "API: [%s] is registered as new user", user)
		resp = Response{Code: 200, Msg: "注册成功", Data: user}
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

// @Title Login
// @Description login as a user
// @Param   username     formData    string  true        "The username to login"
// @Param   password     formData    string  true        "The password"
// @Success 200 {object} controllers.api_controller.Response The Response object
// @router /login [post]
func (c *UserController) Login() {
	var resp Response

	if c.GetSessionUser() != "" {
		resp = Response{Code: 0, Msg: "请先注销当前用户后再登录", Data: c.GetSessionUser()}
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}

	user, password := c.Input().Get("username"), c.Input().Get("password")

	msg := api.CheckUserLogin(user, password)
	if msg != "" {
		resp = Response{Code: 0, Msg: msg, Data: ""}
	} else {
		c.SetSessionUser(user)

		util.LogInfo(c.Ctx, "API: [%s] logged in", user)
		resp = Response{Code: 200, Msg: "登录成功", Data: user}
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

// @Title Logout
// @Description logout the current user
// @Success 200 {object} controllers.api_controller.Response The Response object
// @router /logout [post]
func (c *UserController) Logout() {
	var resp Response

	user := c.GetSessionUser()
	util.LogInfo(c.Ctx, "API: [%s] logged off", user)

	c.SetSessionUser("")

	resp = Response{Code: 200, Msg: "注销成功", Data: user}

	c.Data["json"] = resp
	c.ServeJSON()
}

// @Title ChangePassword
// @Description change password
// @Param   username     formData    string  true        "The username"
// @Param   oldpassword     formData    string  true        "Old password"
// @Param   newpassword     formData    string  true        "New password"
// @Success 200 {object} controllers.api_controller.Response The Response object
// @router /changepwd [post]
func (c *UserController) ChangePassword() {
	var resp Response
	user := c.Input().Get("username")
	if c.RequireUser(user) {
		return
	}

	oldPassword := c.Input().Get("oldpassword")
	newPassword := c.Input().Get("newpassword")

	msg := api.CheckUserChangePassword(user, oldPassword, newPassword)
	if msg != "" {
		resp = Response{Code: 0, Msg: msg, Data: ""}
	} else {
		api.ChangeUserPassword(user, newPassword)

		util.LogInfo(c.Ctx, "API: [%s] changed his password", user)
		resp = Response{Code: 200, Msg: "修改密码成功", Data: user}
	}

	c.Data["json"] = resp
	c.ServeJSON()
}
