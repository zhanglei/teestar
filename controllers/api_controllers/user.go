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

var CookieSecret string = "gitstar1qaz2wsx"
var CookieKey    string = "gitstar_username"

// User API
type UserController struct {
	beego.Controller
}

func (c *UserController) getSessionUser() string {
	user, ok := c.GetSecureCookie(CookieSecret, CookieKey)
	if !ok {
		return ""
	}

	return user
}

func (c *UserController) setSessionUser(user string) {
	c.SetSecureCookie(CookieSecret, CookieKey, user)
}

// @Title Register
// @Description register a new user
// @Param   username     formData    string  true        "The username to register"
// @Param   password     formData    string  true        "The password"
// @Success 200 {object} controllers.api_controller.Response The Response object
// @router /register [post]
func (c *UserController) Register() {
	var resp Response
	user, password := c.Input().Get("username"), c.Input().Get("password")

	msg := api.CheckUserRegister(user, password)
	if msg != "" {
		resp = Response{Code: 0, Msg: msg, Data: ""}
	} else {
		api.AddUser(user, password)

		c.setSessionUser(user)

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
	user, password := c.Input().Get("username"), c.Input().Get("password")

	msg := api.CheckUserLogin(user, password)
	if msg != "" {
		resp = Response{Code: 0, Msg: msg, Data: ""}
	} else {
		c.setSessionUser(user)

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

	user := c.getSessionUser()
	util.LogInfo(c.Ctx, "API: [%s] logged off", user)

	c.setSessionUser("")

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
