package api_controllers

import (
	"github.com/weilaihui/teestar/api"
	"github.com/weilaihui/teestar/util"
)

// @Title GetUserTargetStatus
// @Description Get the status between the user and the target
// @Param   user     path    string  true        "The username"
// @Param   target     path    string  true        "Another user"
// @Success 200 {object} api.UserTargetStatus The Status object
// @router /:user/status/targets/:target [get]
func (c *UsersController) GetUserTargetStatus() {
	if c.RequireLogin() {
		return
	}

	user := c.GetString(":user")
	target := c.GetString(":target")

	c.Data["json"] = api.GetUserTargetStatus(user, target)
	c.ServeJSON()
}

// @Title GetUserStatus
// @Description Get the status between the user and all other users
// @Param   user     path    string  true        "The username"
// @Success 200 {object} []*api.UserTargetStatus The list of Status objects
// @router /:user/status [get]
func (c *UsersController) GetUserStatus() {
	if c.RequireLogin() {
		return
	}

	user := c.GetString(":user")

	util.LogInfo(c.Ctx, "API: [%s] viewed status", user)

	c.Data["json"] = api.GetUserStatus(user)
	c.ServeJSON()
}

// @Title GetUserRecommend
// @Description Get the recommend repos for the user to star
// @Param   user     path    string  true        "The username"
// @Success 200 {object} []api.Entry The list of Entry objects
// @router /:user/status/recommend [get]
func (c *UsersController) GetUserRecommend() {
	if c.RequireLogin() {
		return
	}

	user := c.GetString(":user")

	util.LogInfo(c.Ctx, "API: [%s] viewed recommend", user)

	c.Data["json"] = api.GetUserRecommend(user)
	c.ServeJSON()
}

// @Title GetUserOwe
// @Description Get the details that owes user stars
// @Param   user     path    string  true        "The username"
// @Success 200 {object} []*api.UserTargetStatus The list of Status objects
// @router /:user/status/owe [get]
func (c *UsersController) GetUserOwe() {
	if c.RequireLogin() {
		return
	}

	user := c.GetString(":user")

	util.LogInfo(c.Ctx, "API: [%s] viewed owe", user)

	c.Data["json"] = api.GetUserOwe(user)
	c.ServeJSON()
}
