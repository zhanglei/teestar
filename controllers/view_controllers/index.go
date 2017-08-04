package view_controllers

import (
	"html/template"
	"strings"

	"github.com/astaxie/beego"
	"github.com/hsluoyz/gitstar/api"
	"github.com/hsluoyz/gitstar/util"
)

type ViewController struct {
	beego.Controller
}

func (c *ViewController) getUsername() string {
	username := c.GetSession("username")
	if username == nil {
		return ""
	}

	return username.(string)
}

func (c *ViewController) setUsername(username string) {
	c.SetSession("username", username)
}

//首页
func (c *ViewController) Index() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()

	username := c.getUsername()
	if username == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	objUser := api.GetExtendedUser(username)

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

	c.Data["Recommend"] = api.GetUserRecommend(username)

	util.LogInfo(c.Ctx, "[%s] viewed homepage", username)

	c.Data["PageTitle"] = "GitStar - GitHub项目点赞"
	c.Layout = "layout/layout.tpl"
	c.TplName = "index.tpl"
}

func (c *ViewController) OwePage() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()

	username := c.getUsername()
	if username == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	c.Data["IsLogin"] = true
	c.Data["UserInfo"] = api.GetExtendedUser(username)

	c.Data["Owe"] = api.GetUserOwe(username)

	util.LogInfo(c.Ctx, "[%s] viewed owe page", username)

	c.Data["PageTitle"] = "GitStar - 欠我赞的人"
	c.Layout = "layout/layout.tpl"
	c.TplName = "owe.tpl"
}

func (c *ViewController) OwesPage() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()

	username := c.getUsername()
	if username == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	c.Data["IsLogin"] = true
	c.Data["UserInfo"] = api.GetUser(username)

	c.Data["Owe"] = api.GetOwe()

	util.LogInfo(c.Ctx, "[%s] viewed owe ranking", username)

	c.Data["PageTitle"] = "GitStar - 欠赞排行"
	c.Layout = "layout/layout.tpl"
	c.TplName = "owes.tpl"
}

func (c *ViewController) Update() {
	flash := beego.NewFlash()

	username := c.getUsername()
	if username == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	api.UpdateUserStarringRepos(username)

	util.LogInfo(c.Ctx, "[%s] updated his stars", username)
	c.Redirect("/", 302)
}

//登录页
func (c *ViewController) LoginPage() {
	username := c.getUsername()
	if username != "" {
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
	username, password := c.Input().Get("username"), c.Input().Get("password")

	msg := api.CheckUserLogin(username, password)
	if msg != "" {
		flash.Error(msg)
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
	} else {
		c.setUsername(username)

		util.LogInfo(c.Ctx, "[%s] logged in", username)
		c.Redirect("/", 302)
	}
}

//注册页
func (c *ViewController) RegisterPage() {
	username := c.getUsername()
	if username != "" {
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
	username, password := c.Input().Get("username"), c.Input().Get("password")

	msg := api.CheckUserRegister(username, password)
	if msg != "" {
		flash.Error(msg)
		flash.Store(&c.Controller)
		c.Redirect("/register", 302)
	} else {
		api.AddUser(username, password)

		c.setUsername(username)

		util.LogInfo(c.Ctx, "[%s] is registered as new user", username)
		c.Redirect("/user/setting", 302)
	}
}

//登出
func (c *ViewController) Logout() {
	username := c.getUsername()
	util.LogInfo(c.Ctx, "[%s] logged off", username)

	c.setUsername("")
	c.Redirect("/", 302)
}

//关于
func (c *ViewController) About() {
	username := c.getUsername()
	if username != "" {
		c.Data["IsLogin"] = true
		c.Data["UserInfo"] = api.GetUser(username)
	}

	util.LogInfo(c.Ctx, "[%s] viewed about", username)

	c.Data["PageTitle"] = "GitStar - 关于"
	c.Layout = "layout/layout.tpl"
	c.TplName = "about.tpl"
}

type EscapedRepo struct {
	Repo        string
	RepoEscaped string
}

func (c *ViewController) SettingPage() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()

	username := c.getUsername()
	if username == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	c.Data["IsLogin"] = true
	c.Data["UserInfo"] = api.GetUser(username)
	repos := api.GetUserRepos(username)
	c.Data["Repos"] = repos

	escapedRepos := []EscapedRepo{}
	for _, repo := range repos {
		escaped := strings.Replace(repo, "/", ".", -1)
		escapedRepos = append(escapedRepos, EscapedRepo{Repo: repo, RepoEscaped: escaped})
	}
	c.Data["EscapedRepos"] = escapedRepos

	util.LogInfo(c.Ctx, "[%s] viewed setting", username)

	c.Data["PageTitle"] = "GitStar - 用户设置"
	c.Layout = "layout/layout.tpl"
	c.TplName = "user/setting.tpl"
}

func (c *ViewController) Setting() {
	flash := beego.NewFlash()

	username := c.getUsername()
	if username == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	hitter := c.Input().Get("hitter")

	msg := api.CheckUserUpdateHitter(username, hitter)
	if msg != "" {
		flash.Error(msg)
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}

	api.UpdateUserHitter(username, hitter)

	qq := c.Input().Get("qq")

	if qq == "" {
		flash.Error("请填写QQ号")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}

	if qq != "" && !api.HasQQUser(qq) {
		flash.Error("QQ号不是合法的、已存在的QQ号码")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}

	api.UpdateUserQQ(username, qq)

	nickname := c.Input().Get("nickname")
	api.UpdateUserNickname(username, nickname)

	email := c.Input().Get("email")
	api.UpdateUserEmail(username, email)

	util.LogInfo(c.Ctx, "[%s] updated his setting", username)

	flash.Success("更新资料成功")
	flash.Store(&c.Controller)
	c.Redirect("/user/setting", 302)
}

func (c *ViewController) ChangeUserPassword() {
	flash := beego.NewFlash()

	username := c.getUsername()
	if username == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	oldPassword := c.Input().Get("oldpassword")
	newPassword := c.Input().Get("newpassword")

	msg := api.CheckUserChangePassword(username, oldPassword, newPassword)
	if msg != "" {
		flash.Error(msg)
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
	} else {
		api.ChangeUserPassword(username, newPassword)

		util.LogInfo(c.Ctx, "[%s] changed his password", username)

		flash.Success("修改密码成功")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
	}
}

func (c *ViewController) AddRepoPage() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()

	username := c.getUsername()
	if username == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	objUser := api.GetUser(username)
	if objUser.QQ == "" {
		flash.Error("填写QQ号后才能添加项目")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}

	c.Data["IsLogin"] = true
	c.Data["UserInfo"] = api.GetUser(username)

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

	username := c.getUsername()
	if username == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	repo := c.GetString("name")
	repo = formatRepoAddress(repo)

	msg := api.CheckAddRepo(username, repo)
	if msg != "" {
		flash.Error(msg)
		flash.Store(&c.Controller)
		c.Redirect("/repo/add", 302)
		return
	}

	api.AddUserRepo(username, repo)

	util.LogInfo(c.Ctx, "[%s] added repo: [%s]", username, repo)

	flash.Success("添加项目成功")
	flash.Store(&c.Controller)
	c.Redirect("/user/setting", 302)
}

func (c *ViewController) DeleteRepo() {
	flash := beego.NewFlash()

	username := c.getUsername()
	if username == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	repo := c.GetString(":repo")
	repo = strings.Replace(repo, ".", "/", -1)

	msg := api.CheckDeleteRepo(username, repo)
	if msg != "" {
		flash.Error(msg)
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}

	api.DeleteUserRepo(username, repo)

	util.LogInfo(c.Ctx, "[%s] deleted repo: [%s]", username, repo)

	flash.Success("删除项目成功")
	flash.Store(&c.Controller)
	c.Redirect("/user/setting", 302)
}

func (c *ViewController) UserPage() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()

	username := c.getUsername()
	if username == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	target := c.GetString(":user")

	util.LogInfo(c.Ctx, "[%s] viewed [%s]'s profile", username, target)

	c.Data["IsLogin"] = true
	c.Data["UserInfo"] = api.GetUser(username)

	c.Data["TargetInfo"] = api.GetExtendedUser(target)

	c.Data["PageTitle"] = "GitStar - 用户：" + target
	c.Layout = "layout/layout.tpl"
	c.TplName = "user/detail.tpl"
}

func (c *ViewController) UsersPage() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()

	username := c.getUsername()
	if username == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	users := api.GetUserObjects()

	c.Data["UserInfos"] = users

	util.LogInfo(c.Ctx, "[%s] viewed user list", username)

	c.Data["IsLogin"] = true
	c.Data["UserInfo"] = api.GetUser(username)

	c.Data["PageTitle"] = "GitStar - 用户列表"
	c.Layout = "layout/layout.tpl"
	c.TplName = "user/list.tpl"
}

func (c *ViewController) CountPage() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()

	username := c.getUsername()
	if username == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	users := api.GetExtendedUserObjects()

	c.Data["UserInfos"] = users

	util.LogInfo(c.Ctx, "[%s] viewed count page", username)

	c.Data["IsLogin"] = true
	c.Data["UserInfo"] = api.GetUser(username)

	c.Data["PageTitle"] = "GitStar - 用户统计数据"
	c.Layout = "layout/layout.tpl"
	c.TplName = "user/count.tpl"
}

func (c *ViewController) RepoPage() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()

	username := c.getUsername()
	if username == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	c.Data["IsLogin"] = true
	c.Data["UserInfo"] = api.GetExtendedUser(username)

	repos := api.GetRepoObjects(username)
	c.Data["Repos"] = repos

	util.LogInfo(c.Ctx, "[%s] viewed repo page", username)

	c.Data["PageTitle"] = "GitStar - 我的项目"
	c.Layout = "layout/layout.tpl"
	c.TplName = "repo.tpl"
}

func (c *ViewController) ReferrerPage() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()

	username := c.getUsername()
	if username == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	c.Data["IsLogin"] = true
	c.Data["UserInfo"] = api.GetUser(username)

	util.LogInfo(c.Ctx, "[%s] viewed referrer page", username)

	c.Data["PageTitle"] = "GitStar - Referrer测试"
	c.Layout = "layout/layout.tpl"
	c.TplName = "referrer.tpl"
}

func (c *ViewController) LogPage() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()

	username := c.getUsername()
	if username == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	c.Data["IsLogin"] = true
	c.Data["UserInfo"] = api.GetUser(username)
	c.Data["Log"] = util.ReadLog()

	util.LogInfo(c.Ctx, "[%s] viewed log page", username)

	c.Data["PageTitle"] = "GitStar - 系统日志"
	c.Layout = "layout/layout.tpl"
	c.TplName = "log.tpl"
}
