package view_controllers

import (
	"github.com/astaxie/beego"
	"github.com/weilaihui/teestar/api"
	"github.com/weilaihui/teestar/util"
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
	qq := c.Input().Get("qq")

	msg := api.CheckUserRegister(user, password)
	if msg != "" {
		flash.Error(msg)
		flash.Store(&c.Controller)
		c.Redirect("/register", 302)
		return
	}

	msg = api.CheckUserUpdateQQ(qq)
	if msg != "" {
		flash.Error(msg)
		flash.Store(&c.Controller)
		c.Redirect("/register", 302)
		return
	}

	api.AddUser(user, password)
	api.UpdateUserQQ(user, qq)

	c.SetSessionUser(user)

	util.LogInfo(c.Ctx, "[%s] is registered as new user", user)
	flash.Success("注册成功！请点击下方“我需要被别人点赞的项目”右侧的“添加项目”，添加你需要被其他人点赞的GitHub项目地址")
	flash.Store(&c.Controller)
	c.Redirect("/user/setting", 302)
}

//登出
func (c *ViewController) Logout() {
	user := c.GetSessionUser()
	util.LogInfo(c.Ctx, "[%s] logged off", user)

	c.SetSessionUser("")
	c.Redirect("/", 302)
}
