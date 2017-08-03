package api_controller

import "github.com/hsluoyz/gitstar/api"

func (c *APIController) GetUsers() {
	users := api.GetUsers()
	c.Data["json"] = users
	c.ServeJSON()
}

func (c *APIController) GetUser() {
	user := c.GetString(":user")

	objUser := api.GetUser(user)
	c.Data["json"] = objUser
	c.ServeJSON()
}
