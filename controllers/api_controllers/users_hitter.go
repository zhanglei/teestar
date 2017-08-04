package api_controllers

import (
	"github.com/hsluoyz/gitstar/api"
	"github.com/hsluoyz/gitstar/util"
)

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

	msg := api.CheckUserUpdateHitter(user, hitter)
	if msg != "" {
		resp = Response{Code: 0, Msg: msg, Data: ""}
	} else {
		affected := api.UpdateUserHitter(user, hitter)
		util.LogInfo(c.Ctx, "API: [%s] updated his hitter", user)

		if affected {
			resp = Response{Code: 200, Msg: "ok", Data: ""}
		} else {
			resp = Response{Code: 200, Msg: "not affected", Data: ""}
		}
	}

	c.Data["json"] = resp
	c.ServeJSON()
}
