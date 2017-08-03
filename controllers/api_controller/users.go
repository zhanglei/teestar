package api_controller

import "github.com/hsluoyz/gitstar/api"

func (c *UsersController) GetUsers() {
	users := api.GetUsers()
	c.Data["json"] = users
	c.ServeJSON()
}

func (c *UsersController) GetUser() {
	user := c.GetString(":user")

	objUser := api.GetUser(user)
	objUser.Password = ""
	c.Data["json"] = objUser
	c.ServeJSON()
}
