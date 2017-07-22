package controllers

import (
	"github.com/astaxie/beego"
	"github.com/hsluoyz/gitstar/api"
)

type MainController struct {
	beego.Controller
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
