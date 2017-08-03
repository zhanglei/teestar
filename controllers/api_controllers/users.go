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
// @Description get user list
// @Success 200 {[]string}
// @router / [get]
func (c *UsersController) GetUsers() {
	users := api.GetUsers()
	c.Data["json"] = users
	c.ServeJSON()
}

// @Title GetUser
// @Description get user
// @Param   user     path    string  true        "The username"
// @Success 200 {object} api.User
// @router /:user [get]
func (c *UsersController) GetUser() {
	user := c.GetString(":user")

	objUser := api.GetUser(user)
	objUser.Password = ""
	c.Data["json"] = objUser
	c.ServeJSON()
}
