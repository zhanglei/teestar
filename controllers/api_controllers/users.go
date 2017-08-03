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
// @Success 200 {object} api.User The user object
// @router /:user [get]
func (c *UsersController) GetUser() {
	user := c.GetString(":user")

	objUser := api.GetUser(user)
	objUser.Password = ""
	c.Data["json"] = objUser
	c.ServeJSON()
}
