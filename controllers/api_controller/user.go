package api_controller

import "github.com/hsluoyz/gitstar/api"

func (c *APIController) GetUsers() {
	users := api.GetUsers()
	c.Data["json"] = users
	c.ServeJSON()
}
