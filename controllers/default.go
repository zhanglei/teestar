package controllers

import (
	"github.com/astaxie/beego"
	"github.com/hsluoyz/gitstar/api"
	"strings"
)

type MainController struct {
	beego.Controller
}

var adapter *Adapter

func init() {
	adapter = NewAdapter("mysql", "root:@tcp(127.0.0.1:3306)/")
}

func (c *MainController) GetUserAllRepos() {
	user := c.GetString(":user")

	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"

	repos := api.ListRepos(user)
	c.Data["json"] = repos
	c.ServeJSON()
}

func (c *MainController) GetUserRepos() {
	user := c.GetString(":user")

	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"

	userRepos := UserRepos{User:user}
	has, err := adapter.engine.Get(&userRepos)
	if err != nil {
		panic(err)
	}

	if !has {
		c.Data["json"] = []string{}
		c.ServeJSON()
	}

	c.Data["json"] = strings.Split(userRepos.Repos, ",")
	c.ServeJSON()
}
