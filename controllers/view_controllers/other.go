package view_controllers

import (
	"github.com/astaxie/beego"
	"github.com/hsluoyz/gitstar/api"
	"github.com/hsluoyz/gitstar/util"
)

//关于
func (c *ViewController) About() {
	user := c.GetSessionUser()
	if user != "" {
		c.Data["IsLogin"] = true
		c.Data["UserInfo"] = api.GetUser(user)
	}

	util.LogInfo(c.Ctx, "[%s] viewed about", user)

	c.Data["PageTitle"] = "GitStar - 关于"
	c.Layout = "layout/layout.tpl"
	c.TplName = "about.tpl"
}

func (c *ViewController) QuestionAndAnswer() {
	user := c.GetSessionUser()
	if user != "" {
		c.Data["IsLogin"] = true
		c.Data["UserInfo"] = api.GetUser(user)
	}

	util.LogInfo(c.Ctx, "[%s] viewed qa", user)

	c.Data["PageTitle"] = "GitStar - 常见问题"
	c.Layout = "layout/layout.tpl"
	c.TplName = "qa.tpl"
}

func (c *ViewController) UserPage() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()

	user := c.GetSessionUser()
	if user == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	target := c.GetString(":user")

	util.LogInfo(c.Ctx, "[%s] viewed [%s]'s profile", user, target)

	c.Data["IsLogin"] = true
	c.Data["UserInfo"] = api.GetUser(user)

	c.Data["TargetInfo"] = api.GetExtendedUser(target)
	c.Data["TargetRepos"] = api.GetUserRepoObjects(target)
	c.Data["TargetFollowedStatus"] = api.GetUserFollowedStatus(target)

	c.Data["PageTitle"] = "GitStar - 用户：" + target
	c.Layout = "layout/layout.tpl"
	c.TplName = "user/detail.tpl"
}

func (c *ViewController) ReferrerPage() {
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

	util.LogInfo(c.Ctx, "[%s] viewed referrer page", user)

	c.Data["PageTitle"] = "GitStar - Referrer测试"
	c.Layout = "layout/layout.tpl"
	c.TplName = "referrer.tpl"
}
