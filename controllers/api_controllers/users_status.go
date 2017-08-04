package api_controllers

import "github.com/hsluoyz/gitstar/api"

// @Title GetUserTargetStatus
// @Description Get the status between the user and the target
// @Param   user     path    string  true        "The username"
// @Param   target     path    string  true        "Another user"
// @Success 200 {object} api.UserTargetStatus The Status object
// @router /:user/status/targets/:target [get]
func (c *UsersController) GetUserTargetStatus() {
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
	user := c.GetString(":user")

	statusList := api.GetUserStatus(user)

	c.Data["json"] = statusList
	c.ServeJSON()
}

// @Title GetUserRecommend
// @Description Get the recommend repos for the user to star
// @Param   user     path    string  true        "The username"
// @Success 200 {object} []api.Entry The list of Entry objects
// @router /:user/status/recommend [get]
func (c *UsersController) GetUserRecommend() {
	user := c.GetString(":user")

	c.Data["json"] = api.GetUserRecommend(user)
	c.ServeJSON()
}

// @Title GetUserOwe
// @Description Get the details that owes user stars
// @Param   user     path    string  true        "The username"
// @Success 200 {object} []*api.UserTargetStatus The list of status objects
// @router /:user/status/owe [get]
func (c *UsersController) GetUserOwe() {
	user := c.GetString(":user")

	statusList := api.GetUserOwe(user)

	c.Data["json"] = statusList
	c.ServeJSON()
}
