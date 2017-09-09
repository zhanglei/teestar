package view_controllers

import (
	"html/template"

	"github.com/astaxie/beego"
	"github.com/hsluoyz/gitstar/api"
	"github.com/hsluoyz/gitstar/controllers"
	"github.com/hsluoyz/gitstar/util"
)

type ViewController struct {
	controllers.BaseController
}

//首页
func (c *ViewController) Index() {
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

	if objUser.IsDisabled {
		flash.Error("该账户已被管理员禁用，项目已处于隐藏状态。有问题请联系管理员，QQ群：646373152")
	}

	messages := api.GetSystemMessages(user)
	for _, message := range messages {
		if !message.IsHTML {
			flash.Data[message.Type] = message.Text
			c.Data["flash"] = flash.Data
		} else {
			flash.Data[message.Type] = " "
			c.Data["flash"] = flash.Data
			c.Data["flash_html_" + message.Type] = template.HTML(message.Text)
		}
	}

	c.Data["IsLogin"] = true
	c.Data["UserInfo"] = objUser

	c.Data["Recommend"] = api.GetUserRecommend(user)

	util.LogInfo(c.Ctx, "[%s] viewed homepage", user)

	c.Data["PageTitle"] = "GitStar - GitHub项目点赞"
	c.Layout = "layout/layout.tpl"
	c.TplName = "index/index.tpl"
}

func (c *ViewController) RepoPage() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()

	user := c.GetSessionUser()
	if user == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	messages := api.GetSystemMessages(user)
	for _, message := range messages {
		if !message.IsHTML {
			flash.Data[message.Type] = message.Text
			c.Data["flash"] = flash.Data
		} else {
			flash.Data[message.Type] = " "
			c.Data["flash"] = flash.Data
			c.Data["flash_html_" + message.Type] = template.HTML(message.Text)
		}
	}

	c.Data["IsLogin"] = true
	c.Data["UserInfo"] = api.GetExtendedUser(user)

	c.Data["Repos"] = api.GetExtendedUserRepoObjects(user)

	util.LogInfo(c.Ctx, "[%s] viewed repo page", user)

	c.Data["PageTitle"] = "GitStar - 我的项目"
	c.Layout = "layout/layout.tpl"
	c.TplName = "index/repo.tpl"
}

func (c *ViewController) OwePage() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()

	user := c.GetSessionUser()
	if user == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	messages := api.GetSystemMessages(user)
	for _, message := range messages {
		if !message.IsHTML {
			flash.Data[message.Type] = message.Text
			c.Data["flash"] = flash.Data
		} else {
			flash.Data[message.Type] = " "
			c.Data["flash"] = flash.Data
			c.Data["flash_html_" + message.Type] = template.HTML(message.Text)
		}
	}

	c.Data["IsLogin"] = true
	c.Data["UserInfo"] = api.GetExtendedUser(user)

	c.Data["Owe"] = api.GetUserOwe(user)

	util.LogInfo(c.Ctx, "[%s] viewed owe page", user)

	c.Data["PageTitle"] = "GitStar - 欠我赞的人"
	c.Layout = "layout/layout.tpl"
	c.TplName = "index/owe.tpl"
}

func (c *ViewController) OwesPage() {
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

	c.Data["Owe"] = api.GetOwe()

	util.LogInfo(c.Ctx, "[%s] viewed owe ranking", user)

	c.Data["PageTitle"] = "GitStar - 欠赞排行"
	c.Layout = "layout/layout.tpl"
	c.TplName = "index/owes.tpl"
}

func (c *ViewController) Update() {
	flash := beego.NewFlash()

	user := c.GetSessionUser()
	if user == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	api.UpdateUserStarringRepos(user)

	util.LogInfo(c.Ctx, "[%s] updated his stars", user)
	c.Redirect("/", 302)
}
