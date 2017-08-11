package view_controllers

import (
	"github.com/astaxie/beego"
	"github.com/hsluoyz/gitstar/api"
	"github.com/hsluoyz/gitstar/util"
)

//登录页
func (c *ViewController) LoginPage() {
	user := c.GetSessionUser()
	if user != "" {
		c.Redirect("/", 302)
	} else {
		beego.ReadFromRequest(&c.Controller)
		c.Data["PageTitle"] = "GitStar - 登录"
		c.Layout = "layout/layout.tpl"
		c.TplName = "login.tpl"
	}
}

//验证登录
func (c *ViewController) Login() {
	flash := beego.NewFlash()
	user, password := c.Input().Get("username"), c.Input().Get("password")

	msg := api.CheckUserLogin(user, password)
	if msg != "" {
		flash.Error(msg)
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
	} else {
		c.SetSessionUser(user)

		util.LogInfo(c.Ctx, "[%s] logged in", user)
		c.Redirect("/", 302)
	}
}

//注册页
func (c *ViewController) RegisterPage() {
	user := c.GetSessionUser()
	if user != "" {
		c.Redirect("/", 302)
	} else {
		beego.ReadFromRequest(&c.Controller)
		c.Data["PageTitle"] = "GitStar - 注册"
		c.Layout = "layout/layout.tpl"
		c.TplName = "register.tpl"
	}
}

//验证注册
func (c *ViewController) Register() {
	flash := beego.NewFlash()
	user, password := c.Input().Get("username"), c.Input().Get("password")

	msg := api.CheckUserRegister(user, password)
	if msg != "" {
		flash.Error(msg)
		flash.Store(&c.Controller)
		c.Redirect("/register", 302)
	} else {
		api.AddUser(user, password)

		c.SetSessionUser(user)

		util.LogInfo(c.Ctx, "[%s] is registered as new user", user)
		c.Redirect("/user/setting", 302)
	}
}

//登出
func (c *ViewController) Logout() {
	user := c.GetSessionUser()
	util.LogInfo(c.Ctx, "[%s] logged off", user)

	c.SetSessionUser("")
	c.Redirect("/", 302)
}
