package api_controllers

import (
	"github.com/hsluoyz/gitstar/api"
	"github.com/hsluoyz/gitstar/util"
)

// Users API
type UsersController struct {
	BaseController
}

// @Title GetUsers
// @Description Get user list
// @Success 200 {[]string}
// @router / [get]
func (c *UsersController) GetUsers() {
	if c.RequireAdmin() {
		return
	}

	users := api.GetUsers()

	c.Data["json"] = users
	c.ServeJSON()
}

// @Title GetUser
// @Description Get a user
// @Param   user     path    string  true        "The username"
// @Success 200 {object} api.User The User object
// @router /:user [get]
func (c *UsersController) GetUser() {
	if c.RequireLogin() {
		return
	}

	user := c.GetString(":user")

	objUser := api.GetUser(user)
	objUser.Password = ""
	c.Data["json"] = objUser
	c.ServeJSON()
}

// @Title GetExtendedUser
// @Description Get a user with extended information
// @Param   user     path    string  true        "The username"
// @Success 200 {object} api.ExtendedUser The ExtendedUser object
// @router /:user/extended [get]
func (c *UsersController) GetExtendedUser() {
	if c.RequireLogin() {
		return
	}

	user := c.GetString(":user")

	objUser := api.GetExtendedUser(user)
	c.Data["json"] = objUser
	c.ServeJSON()
}

// @Title UpdateUser
// @Description Update a user
// @Param   user     path      string  true    "The user that needs to be updated"
// @Param   hitter   formData  string  true    "The GitHub account the user uses to star other repos"
// @Param   qq       formData  string  true    "The Tencent QQ account"
// @Param   nickname formData  string  true    "The Tencent QQ nickname"
// @Param   email    formData  string  true    "The Email address"
// @Success 200 {object} controllers.api_controller.Response The Response object
// @router /:user/update [post]
func (c *UsersController) UpdateUser() {
	user := c.GetString(":user")
	if c.RequireUser(user) {
		return
	}

	var resp Response
	hitter := c.Input().Get("hitter")
	qq := c.Input().Get("qq")
	nickname := c.Input().Get("nickname")
	email := c.Input().Get("email")

	msg := api.CheckUserUpdateHitter(user, hitter)
	if msg != "" {
		resp = Response{Code: 0, Msg: msg, Data: ""}
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}

	api.UpdateUserHitter(user, hitter)

	msg = api.CheckUserUpdateQQ(qq)
	if msg != "" {
		resp = Response{Code: 0, Msg: msg, Data: ""}
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}

	api.UpdateUserQQ(user, qq)

	api.UpdateUserNickname(user, nickname)

	api.UpdateUserEmail(user, email)

	util.LogInfo(c.Ctx, "API: [%s] updated his setting", user)

	resp = Response{Code: 200, Msg: "更新资料成功", Data: ""}

	c.Data["json"] = resp
	c.ServeJSON()
}

