package view_controller

import (
	"github.com/astaxie/beego"
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
	c.Data["IsLogin"] = true
	c.Data["Username"] = username

	c.Data["Recommend"] = api.GetUserRecommend(username)

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

	flag := api.CheckUserPassword(username, password)
	if flag {
		c.setUsername(username)
		c.Redirect("/", 302)
	} else {
		flash.Error("用户名或密码错误")
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
	} else {
		api.AddUser(username, password)

		c.setUsername(username)
		c.Redirect("/", 302)
	}
}

//登出
func (c *ViewController) Logout() {
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

	c.Data["PageTitle"] = "GitStar - 关于"
	c.Layout = "layout/layout.tpl"
	c.TplName = "about.tpl"
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
	c.Data["Hitter"] = api.GetUserHitter(username)
	c.Data["Repos"] = api.GetUserRepos(username)

	c.Data["PageTitle"] = "GitStar - 用户设置"
	c.Layout = "layout/layout.tpl"
	c.TplName = "user/setting.tpl"
}

func (c *ViewController) Setting() {
	flash := beego.NewFlash()

	username := c.getUsername()
	if username == "" {
		//flash.Error("请先登录")
		//flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	hitter := c.Input().Get("hitter")
	api.UpdateUserHitter(username, hitter)

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

	c.Data["IsLogin"] = true
	c.Data["Username"] = username

	c.Data["PageTitle"] = "GitStar - 添加项目"
	c.Layout = "layout/layout.tpl"
	c.TplName = "repo/add.tpl"
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

	api.AddUserRepo(username, repo)

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

	api.DeleteUserRepo(username, repo)

	flash.Success("删除项目成功")
	flash.Store(&c.Controller)
	c.Redirect("/user/setting", 302)
}
