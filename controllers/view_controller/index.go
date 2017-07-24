package view_controller

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/session"
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

////首页
//func (c *ViewController) Index() {
//	c.Data["PageTitle"] = "首页"
//	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)
//	p, _ := strconv.Atoi(c.Ctx.Input.Query("p"))
//	if p == 0 {
//		p = 1
//	}
//	size, _ := beego.AppConfig.Int("page.size")
//	s, _ := strconv.Atoi(c.Ctx.Input.Query("s"))
//	c.Data["S"] = s
//	section := models.Section{Id: s}
//	c.Data["Page"] = models.PageTopic(p, size, &section)
//	c.Data["Sections"] = models.FindAllSection()
//	c.Layout = "layout/layout.tpl"
//	c.TplName = "index.tpl"
//}
//
////登录页
//func (c *ViewController) LoginPage() {
//	IsLogin, _ := filters.IsLogin(c.Ctx)
//	if IsLogin {
//		c.Redirect("/", 302)
//	} else {
//		beego.ReadFromRequest(&c.Controller)
//		c.Data["PageTitle"] = "登录"
//		c.Layout = "layout/layout.tpl"
//		c.TplName = "login.tpl"
//	}
//}
//
////验证登录
//func (c *ViewController) Login() {
//	flash := beego.NewFlash()
//	username, password := c.Input().Get("username"), c.Input().Get("password")
//	if flag, user := models.Login(username, password); flag {
//		c.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), user.Token, 30 * 24 * 60 * 60, "/", beego.AppConfig.String("cookie.domain"), false, true)
//		c.Redirect("/", 302)
//	} else {
//		flash.Error("用户名或密码错误")
//		flash.Store(&c.Controller)
//		c.Redirect("/login", 302)
//	}
//}

//注册页
func (c *ViewController) RegisterPage() {
	username := getUsername(c.Ctx)
	if username != "" {
		c.Redirect("/", 302)
	} else {
		beego.ReadFromRequest(&c.Controller)
		c.Data["PageTitle"] = "注册"
		c.Layout = "layout/layout.tpl"
		c.TplName = "register.tpl"
	}
}

////验证注册
//func (c *ViewController) Register() {
//	flash := beego.NewFlash()
//	username, password := c.Input().Get("username"), c.Input().Get("password")
//	if len(username) == 0 || len(password) == 0 {
//		flash.Error("用户名或密码不能为空")
//		flash.Store(&c.Controller)
//		c.Redirect("/register", 302)
//	} else if api.HasUser() {
//		flash.Error("用户名已被注册")
//		flash.Store(&c.Controller)
//		c.Redirect("/register", 302)
//	} else {
//		var token = uuid.Rand().Hex()
//		user := models.User{Username: username, Password: password, Avatar: "/static/imgs/avatar.png", Token: token}
//		models.SaveUser(&user)
//		// others are ordered as cookie's max age time, path,domain, secure and httponly.
//		c.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), token, 30 * 24 * 60 * 60, "/", beego.AppConfig.String("cookie.domain"), false, true)
//		c.Redirect("/", 302)
//	}
//}

////登出
//func (c *ViewController) Logout() {
//	c.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), "", -1, "/", beego.AppConfig.String("cookie.domain"), false, true)
//	c.Redirect("/", 302)
//}
//
