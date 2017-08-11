package view_controllers

import (
	"github.com/astaxie/beego"
	"github.com/hsluoyz/gitstar/api"
	"github.com/hsluoyz/gitstar/util"
)

func (c *ViewController) FollowIndex() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()

	user := c.GetSessionUser()
	if user == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	objUser := api.GetExtendedUser(user)

	c.Data["IsLogin"] = true
	c.Data["UserInfo"] = objUser

	c.Data["CanFollow"] = api.GetUserCanFollowStatus(user)

	util.LogInfo(c.Ctx, "[%s] viewed follow index page", user)

	c.Data["PageTitle"] = "GitStar - 互粉主页"
	c.Layout = "layout/layout.tpl"
	c.TplName = "index/follow_index.tpl"
}

func (c *ViewController) FollowerPage() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()

	user := c.GetSessionUser()
	if user == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	objUser := api.GetExtendedUser(user)

	c.Data["IsLogin"] = true
	c.Data["UserInfo"] = objUser

	c.Data["Followers"] = api.GetUserFollowedStatus(user)

	util.LogInfo(c.Ctx, "[%s] viewed follower page", user)

	c.Data["PageTitle"] = "GitStar - 我的粉丝"
	c.Layout = "layout/layout.tpl"
	c.TplName = "index/follower.tpl"
}

func (c *ViewController) FollowOwePage() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()

	user := c.GetSessionUser()
	if user == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	objUser := api.GetExtendedUser(user)

	c.Data["IsLogin"] = true
	c.Data["UserInfo"] = objUser

	c.Data["Following"] = api.GetUserFollowingStatus(user)

	util.LogInfo(c.Ctx, "[%s] viewed follow owe page", user)

	c.Data["PageTitle"] = "GitStar - 欠我粉的人"
	c.Layout = "layout/layout.tpl"
	c.TplName = "index/follow_owe.tpl"
}

func (c *ViewController) FollowUpdate() {
	flash := beego.NewFlash()

	user := c.GetSessionUser()
	if user == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	api.UpdateUserFollowingTargets(user)

	util.LogInfo(c.Ctx, "[%s] updated his following", user)
	c.Redirect("/follow", 302)
}
