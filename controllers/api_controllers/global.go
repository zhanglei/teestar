package api_controllers

import (
	"github.com/astaxie/beego"
	"github.com/hsluoyz/gitstar/api"
)

type GlobalController struct {
	beego.Controller
}

func (c *GlobalController) UpdateStarringRepos() {
	affected := api.UpdateStarringRepos()

	if affected {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = "not affected"
	}
	c.ServeJSON()
}

func (c *GlobalController) GetRecommend() {
	c.Data["json"] = api.GetRecommend()
	c.ServeJSON()
}

func (c *GlobalController) GetOwe() {
	c.Data["json"] = api.GetOwe()
	c.ServeJSON()
}
