package view_controllers

import (
	"html/template"
	"strings"

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

	messages := api.GetSystemMessages()
	for _, message := range messages {
		if !message.IsHTML {
			flash.Data[message.Type] = message.Text
			c.Data["flash"] = flash.Data
		} else {
			flash.Data[message.Type] = " "
			c.Data["flash"] = flash.Data
			c.Data["flash_data"] = template.HTML(message.Text)
		}
	}

	c.Data["IsLogin"] = true
	c.Data["UserInfo"] = objUser

	c.Data["Recommend"] = api.GetUserRecommend(user)

	util.LogInfo(c.Ctx, "[%s] viewed homepage", user)

	c.Data["PageTitle"] = "GitStar - GitHub项目点赞"
	c.Layout = "layout/layout.tpl"
	c.TplName = "index.tpl"
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

	c.Data["IsLogin"] = true
	c.Data["UserInfo"] = api.GetExtendedUser(user)

	c.Data["Owe"] = api.GetUserOwe(user)

	util.LogInfo(c.Ctx, "[%s] viewed owe page", user)

	c.Data["PageTitle"] = "GitStar - 欠我赞的人"
	c.Layout = "layout/layout.tpl"
	c.TplName = "owe.tpl"
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
	c.TplName = "owes.tpl"
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
		escaped := strings.Replace(repo, "/", ".", -1)
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
	repo = strings.Replace(repo, ".", "/", -1)

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

	c.Data["PageTitle"] = "GitStar - 用户：" + target
	c.Layout = "layout/layout.tpl"
	c.TplName = "user/detail.tpl"
}

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

	c.Data["IsLogin"] = true
	c.Data["UserInfo"] = api.GetExtendedUser(user)

	c.Data["Repos"] = api.GetUserRepoObjects(user)

	util.LogInfo(c.Ctx, "[%s] viewed repo page", user)

	c.Data["PageTitle"] = "GitStar - 我的项目"
	c.Layout = "layout/layout.tpl"
	c.TplName = "repo.tpl"
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
