package api_controllers

import "github.com/hsluoyz/gitstar/api"

// @Title GetUserHitter
// @Description Get the hitter for a user
// @Param   user     path    string  true        "The username"
// @Success 200 {string}
// @router /:user/hitter [get]
func (c *UsersController) GetUserHitter() {
	user := c.GetString(":user")

	hitter := api.GetUserHitter(user)

	c.Data["json"] = hitter
	c.ServeJSON()
}

// @Title UpdateUserHitter
// @Description Get the hitter for a user
// @Param   user     path    string  true        "The username"
// @Param   hitter     path    string  true        "The GitHub account the user uses to star other repos"
// @Success 200 {object} controllers.api_controller.Response The response object
// @router /:user/hitter/update/:hitter [get]
func (c *UsersController) UpdateUserHitter() {
	var resp Response
	user := c.GetString(":user")
	hitter := c.GetString(":hitter")

	affected := api.UpdateUserHitter(user, hitter)

	if affected {
		resp = Response{Code: 200, Msg: "ok", Data: ""}
	} else {
		resp = Response{Code: 200, Msg: "not affected", Data: ""}
	}

	c.Data["json"] = resp
	c.ServeJSON()
}
