package api_controllers

import (
	"github.com/astaxie/beego"
	"github.com/hsluoyz/gitstar/api"
)

// Global API
type GlobalController struct {
	beego.Controller
}

// @Title UpdateStarringRepos
// @Description update all the repos starred by each user into GitStar cache
// @Success 200 {object} controllers.api_controller.Response The response object
// @router /starring-repos/update [get]
func (c *GlobalController) UpdateStarringRepos() {
	var resp Response

	affected := api.UpdateStarringRepos()

	if affected {
		resp = Response{Code: 200, Msg: "ok", Data: ""}
	} else {
		resp = Response{Code: 200, Msg: "not affected", Data: ""}
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

// @Title GetRecommend
// @Description Get the recommend repos for each user
// @Success 200 {object} []api.Entry2 The list of Entry2 objects
// @router /recommend [get]
func (c *GlobalController) GetRecommend() {
	c.Data["json"] = api.GetRecommend()
	c.ServeJSON()
}

// @Title GetOwe
// @Description Get the details that each user owes another user stars
// @Success 200 {object} []*api.UserTargetStatus The list of status objects
// @router /owe [get]
func (c *GlobalController) GetOwe() {
	c.Data["json"] = api.GetOwe()
	c.ServeJSON()
}