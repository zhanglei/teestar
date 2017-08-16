package view_controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/hsluoyz/gitstar/api"
	"github.com/hsluoyz/gitstar/util"
)

type EscapedRepo struct {
	Repo        string
	RepoEscaped string
}

func (c *ViewController) SettingPage() {
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
	repos := api.GetUserRepos(user)
	c.Data["Repos"] = repos

	escapedRepos := []EscapedRepo{}
	for _, repo := range repos {
		escaped := strings.Replace(repo, "/", "~", -1)
		escapedRepos = append(escapedRepos, EscapedRepo{Repo: repo, RepoEscaped: escaped})
	}
	c.Data["EscapedRepos"] = escapedRepos

	util.LogInfo(c.Ctx, "[%s] viewed setting", user)

	c.Data["PageTitle"] = "GitStar - 用户设置"
	c.Layout = "layout/layout.tpl"
	c.TplName = "user/setting.tpl"
}

func (c *ViewController) Setting() {
	flash := beego.NewFlash()

	user := c.GetSessionUser()
	if user == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	hitter := c.Input().Get("hitter")

	msg := api.CheckUserUpdateHitter(user, hitter)
	if msg != "" {
		flash.Error(msg)
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}

	api.UpdateUserHitter(user, hitter)

	qq := c.Input().Get("qq")

	msg = api.CheckUserUpdateQQ(qq)
	if msg != "" {
		flash.Error(msg)
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}

	api.UpdateUserQQ(user, qq)

	nickname := c.Input().Get("nickname")
	api.UpdateUserNickname(user, nickname)

	email := c.Input().Get("email")
	api.UpdateUserEmail(user, email)

	var followable bool
	if c.Input().Get("followable") == "on" {
		followable = true
	} else {
		followable = false
	}
	api.UpdateUserFollowable(user, followable)

	util.LogInfo(c.Ctx, "[%s] updated his setting", user)

	flash.Success("更新资料成功")
	flash.Store(&c.Controller)
	c.Redirect("/user/setting", 302)
}

func (c *ViewController) ChangeUserPassword() {
	flash := beego.NewFlash()

	user := c.GetSessionUser()
	if user == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	oldPassword := c.Input().Get("oldpassword")
	newPassword := c.Input().Get("newpassword")

	msg := api.CheckUserChangePassword(user, oldPassword, newPassword)
	if msg != "" {
		flash.Error(msg)
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
	} else {
		api.ChangeUserPassword(user, newPassword)

		util.LogInfo(c.Ctx, "[%s] changed his password", user)

		flash.Success("修改密码成功")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
	}
}

func (c *ViewController) AddRepoPage() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()

	user := c.GetSessionUser()
	if user == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	objUser := api.GetUser(user)
	if objUser.QQ == "" {
		flash.Error("填写QQ号后才能添加项目")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}

	c.Data["IsLogin"] = true
	c.Data["UserInfo"] = api.GetUser(user)

	c.Data["PageTitle"] = "GitStar - 添加项目"
	c.Layout = "layout/layout.tpl"
	c.TplName = "repo/add.tpl"
}

func formatRepoAddress(repo string) string {
	pos := strings.Index(repo, "github.com/")
	if pos != -1 {
		repo = repo[pos + len("github.com/"):]
	}

	return repo
}

func (c *ViewController) AddRepo() {
	flash := beego.NewFlash()

	user := c.GetSessionUser()
	if user == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	repo := c.GetString("name")
	repo = formatRepoAddress(repo)

	msg := api.CheckAddRepo(user, repo)
	if msg != "" {
		flash.Error(msg)
		flash.Store(&c.Controller)
		c.Redirect("/repo/add", 302)
		return
	}

	api.AddUserRepo(user, repo)

	util.LogInfo(c.Ctx, "[%s] added repo: [%s]", user, repo)

	flash.Success("添加项目成功")
	flash.Store(&c.Controller)
	c.Redirect("/user/setting", 302)
}

func (c *ViewController) DeleteRepo() {
	flash := beego.NewFlash()

	user := c.GetSessionUser()
	if user == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	repo := c.GetString(":repo")
	repo = strings.Replace(repo, "~", "/", -1)

	msg := api.CheckDeleteRepo(user, repo)
	if msg != "" {
		flash.Error(msg)
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}

	api.DeleteUserRepo(user, repo)

	util.LogInfo(c.Ctx, "[%s] deleted repo: [%s]", user, repo)

	flash.Success("删除项目成功")
	flash.Store(&c.Controller)
	c.Redirect("/user/setting", 302)
}
