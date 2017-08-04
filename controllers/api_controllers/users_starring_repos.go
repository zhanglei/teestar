package api_controllers

import (
	"github.com/hsluoyz/gitstar/api"
	"github.com/hsluoyz/gitstar/util"
)

// @Title GetUserStarringRepos
// @Description Get all the repos starred by the user in GitStar cache
// @Param   user     path    string  true        "The username"
// @Success 200 {[]string}
// @router /:user/starring-repos [get]
func (c *UsersController) GetUserStarringRepos() {
	user := c.GetString(":user")

	c.Data["json"] = api.GetUserStarringRepos(user)
	c.ServeJSON()
}

// @Title UpdateUserStarringRepos
// @Description update all the repos starred by the user into GitStar cache
// @Param   user     path    string  true        "The username"
// @Success 200 {object} controllers.api_controller.Response The Response object
// @router /:user/starring-repos/update [get]
func (c *UsersController) UpdateUserStarringRepos() {
	var resp Response
	user := c.GetString(":user")

	api.UpdateUserStarringRepos(user)

	util.LogInfo(c.Ctx, "API: [%s] updated his stars", user)

	resp = Response{Code: 200, Msg: "更新点赞缓存成功", Data: ""}

	c.Data["json"] = resp
	c.ServeJSON()
}
