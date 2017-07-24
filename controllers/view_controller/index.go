package view_controller

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/session"
	"github.com/hsluoyz/gitstar/api"
)

type ViewController struct {
	beego.Controller
}

var globalSessions *session.Manager

func init() {
	config := `{"cookieName":"gosessionid","gclifetime":9999999, "enableSetCookie":true}`
	conf := new(session.ManagerConfig)
	if err := json.Unmarshal([]byte(config), conf); err != nil {
		panic(err)
	}
	globalSessions, _ = session.NewManager("memory", conf)
	go globalSessions.GC()
}

func getUsername(ctx *context.Context) string {
	sess, err := globalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
	if err != nil {
		panic(err)
	}

	username := sess.Get("username")
	if username == nil {
		return ""
	}

	return sess.Get("username").(string)
}

func setUsername(ctx *context.Context, username string) {
	sess, err := globalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
	if err != nil {
		panic(err)
	}

	err = sess.Set("username", username)
	if err != nil {
		panic(err)
	}
}

//首页
func (c *ViewController) Index() {
	c.Data["PageTitle"] = "GitStar - GitHub项目点赞"

	username := getUsername(c.Ctx)
	if username != "" {
		c.Data["IsLogin"] = true
		c.Data["Username"] = username
	}

	c.Layout = "layout/layout.tpl"
	c.TplName = "index.tpl"
}

//登录页
func (c *ViewController) LoginPage() {
	username := getUsername(c.Ctx)
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
	username, _ := c.Input().Get("username"), c.Input().Get("password")

	flag := api.HasUser(username)
	if flag {
		setUsername(c.Ctx, username)
		c.Redirect("/", 302)
	} else {
		flash.Error("用户名或密码错误")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
	}
}

//注册页
func (c *ViewController) RegisterPage() {
	username := getUsername(c.Ctx)
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
		api.AddUser(username)

		setUsername(c.Ctx, username)
		c.Redirect("/", 302)
	}
}

//登出
func (c *ViewController) Logout() {
	setUsername(c.Ctx, "")
	c.Redirect("/", 302)
}

//关于
func (c *ViewController) About() {
	username := getUsername(c.Ctx)
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

	username := getUsername(c.Ctx)
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

	username := getUsername(c.Ctx)
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

	username := getUsername(c.Ctx)
	if username == "" {
		flash.Error("请先登录")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	c.Data["PageTitle"] = "GitStar - 添加项目"
	c.Layout = "layout/layout.tpl"
	c.TplName = "repo/add.tpl"
}

func (c *ViewController) AddRepo() {
	flash := beego.NewFlash()

	username := getUsername(c.Ctx)
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

	username := getUsername(c.Ctx)
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
