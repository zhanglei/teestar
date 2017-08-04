package api_controllers

import (
	"github.com/astaxie/beego"
	"github.com/hsluoyz/gitstar/api"
)

// Users API
type UsersController struct {
	beego.Controller
}

// @Title GetUsers
// @Description Get user list
// @Success 200 {[]string}
// @router / [get]
func (c *UsersController) GetUsers() {
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
	user := c.GetString(":user")

	objUser := api.GetExtendedUser(user)
	c.Data["json"] = objUser
	c.ServeJSON()
}
