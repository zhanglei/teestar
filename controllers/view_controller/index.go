package view_controller

import (
	"html/template"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/hsluoyz/gitstar/api"
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

	flash.Notice("管理员消息：")
	c.Data["flash"] = flash.Data
	c.Data["flash_data"] = template.HTML("目前可以采用以下API访问谁欠了我的Star：<a target=\"_blank\" href=\"/api/users/" + username + "/owe\">https://gitstar.cn/api/users/" + username + "/owe</a>，其中Target是对方的名字，QQ和Nickname是对方QQ号和昵称，Hitter是对方用的GitHub小号，Score表示对方欠我多少个Star，CanBeStarredRepos表示对方还能Star我的项目，可以通过QQ通知对方Star自己。相关UI界面还在开发中..")

	c.Data["IsLogin"] = true
	c.Data["Username"] = username

	c.Data["Recommend"] = api.GetUserRecommend(username)

	logs.Info("[%s] viewed homepage", username)

	c.Data["PageTitle"] = "GitStar - GitHub项目点赞"
	c.Layout = "layout/layout.tpl"
	c.TplName = "index.tpl"
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

	logs.Info("[%s] updated his stars", username)
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

	if !api.HasUser(username) {
		flash.Error("用户名不存在，请先注册")
		flash.Store(&c.Controller)
		c.Redirect("/register", 302)
		return
	}

	flag := api.CheckUserPassword(username, password)
	if flag {
		c.setUsername(username)

		logs.Info("[%s] logged in", username)
		c.Redirect("/", 302)
	} else {
		flash.Error("密码错误")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
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
	if len(username) == 0 || len(password) == 0 {
		flash.Error("用户名或密码不能为空")
		flash.Store(&c.Controller)
		c.Redirect("/register", 302)
	} else if api.HasUser(username) {
		flash.Error("用户名已被注册")
		flash.Store(&c.Controller)
		c.Redirect("/register", 302)
	} else if api.HasHitter("", username) {
		flash.Error("用户名已被其他用户注册为点赞小号")
		flash.Store(&c.Controller)
		c.Redirect("/register", 302)
	} else if strings.Contains(username, "@") {
		flash.Error("请不要使用邮箱，GitHub profile（如https://github.com/abc）中，abc是用户名")
		flash.Store(&c.Controller)
		c.Redirect("/register", 302)
	} else if !api.HasGitHubUser(username) {
		flash.Error("用户名不是合法的、已存在的GitHub用户名")
		flash.Store(&c.Controller)
		c.Redirect("/register", 302)
	} else {
		api.AddUser(username, password)

		c.setUsername(username)

		logs.Info("[%s] is registered as new user", username)
		c.Redirect("/user/setting", 302)
	}
}

//登出
func (c *ViewController) Logout() {
	username := c.getUsername()
	logs.Info("[%s] logged off", username)

	c.setUsername("")
	c.Redirect("/", 302)
}

//关于
func (c *ViewController) About() {
	username := c.getUsername()
	if username != "" {
		c.Data["IsLogin"] = true
		c.Data["Username"] = username
	}

	logs.Info("[%s] viewed about", username)

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
	c.Data["Username"] = username
	userObj := api.GetUser(username)
	c.Data["Hitter"] = userObj.Hitter
	c.Data["QQ"] = userObj.QQ
	c.Data["Nickname"] = userObj.Nickname
	c.Data["Email"] = userObj.Email
	repos := api.GetUserRepos(username)
	c.Data["Repos"] = repos

	escapedRepos := []EscapedRepo{}
	for _, repo := range repos {
		escaped := strings.Replace(repo, "/", ".", -1)
		escapedRepos = append(escapedRepos, EscapedRepo{Repo: repo, RepoEscaped: escaped})
	}
	c.Data["EscapedRepos"] = escapedRepos

	logs.Info("[%s] viewed setting", username)

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

	if hitter != "" && hitter == username {
		flash.Error("不需要把点赞账号（小号）设置为与用户名（大号）一致，留空即表示用大号点赞")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}

	if hitter != "" && api.HasUser(hitter) {
		flash.Error("点赞账号与其他用户的用户名（大号）重复，无法使用")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}

	if hitter != "" && api.HasHitter(username, hitter) {
		flash.Error("点赞账号与其他用户的点赞账号（小号）重复，无法使用")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}

	if hitter != "" && !api.HasGitHubUser(hitter) {
		flash.Error("点赞账号不是合法的、已存在的GitHub用户名")
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

	logs.Info("[%s] updated his setting", username)

	flash.Success("更新资料成功")
	flash.Store(&c.Controller)
	c.Redirect("/user/setting", 302)
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
	c.Data["Username"] = username

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
	if len(repo) == 0 {
		flash.Error("项目不能为空")
		flash.Store(&c.Controller)
		c.Redirect("/repo/add", 302)
		return
	}

	repo = formatRepoAddress(repo)

	if api.HasUserRepo(username, repo) {
		flash.Error("该项目已经存在")
		flash.Store(&c.Controller)
		c.Redirect("/repo/add", 302)
		return
	}

	if !api.HasGitHubRepo(repo) {
		flash.Error("项目地址不是合法的、已存在的GitHub项目地址")
		flash.Store(&c.Controller)
		c.Redirect("/repo/add", 302)
		return
	}

	affected := api.AddUserRepo(username, repo)

	if !affected {
		flash.Error("该项目已经存在")
		flash.Store(&c.Controller)
		c.Redirect("/repo/add", 302)
		return
	}

	logs.Info("[%s] added repo: [%s]", username, repo)

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
	if len(repo) == 0 {
		flash.Error("项目不能为空")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}

	repo = strings.Replace(repo, ".", "/", -1)
	affected := api.DeleteUserRepo(username, repo)

	if !affected {
		flash.Error("该项目不存在")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}

	logs.Info("[%s] deleted repo: [%s]", username, repo)

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
	objTarget := api.GetUser(target)
	if objTarget != nil {
		c.Data["TargetInfo"] = objTarget
	}

	logs.Info("[%s] viewed [%s]'s profile", username, target)

	c.Data["IsLogin"] = true
	c.Data["Username"] = username

	c.Data["PageTitle"] = "GitStar - 用户：" + target
	c.Layout = "layout/layout.tpl"
	c.TplName = "user/detail.tpl"
}
