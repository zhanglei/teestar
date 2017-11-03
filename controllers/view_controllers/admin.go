package view_controllers

import (
	"github.com/astaxie/beego"
	"github.com/weilaihui/teestar/api"
	"github.com/weilaihui/teestar/util"
)

func (c *ViewController) UsersPage() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()

	user := c.GetSessionUser()
	if user == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	users := api.GetUserObjects()

	c.Data["UserInfos"] = users

	util.LogInfo(c.Ctx, "[%s] viewed user list", user)

	c.Data["IsLogin"] = true
	c.Data["UserInfo"] = api.GetUser(user)

	c.Data["PageTitle"] = "GitStar - 用户列表"
	c.Layout = "layout/layout.tpl"
	c.TplName = "user/list.tpl"
}

func (c *ViewController) CountPage() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()

	user := c.GetSessionUser()
	if user == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	users := api.GetExtendedUserObjects()

	c.Data["UserInfos"] = users

	util.LogInfo(c.Ctx, "[%s] viewed count page", user)

	c.Data["IsLogin"] = true
	c.Data["UserInfo"] = api.GetUser(user)

	c.Data["PageTitle"] = "GitStar - 用户统计数据"
	c.Layout = "layout/layout.tpl"
	c.TplName = "user/count.tpl"
}

func (c *ViewController) LogPage() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()

	user := c.GetSessionUser()
	if user == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	c.Data["IsLogin"] = true
	c.Data["UserInfo"] = api.GetUser(user)
	c.Data["Log"] = util.ReadLog()

	util.LogInfo(c.Ctx, "[%s] viewed log page", user)

	c.Data["PageTitle"] = "GitStar - 系统日志"
	c.Layout = "layout/layout.tpl"
	c.TplName = "log.tpl"
}
