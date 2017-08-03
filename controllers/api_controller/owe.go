package api_controller

import "github.com/hsluoyz/gitstar/api"

func (c *UsersController) GetUserOwe() {
	user := c.GetString(":user")

	statusList := api.GetUserOwe(user)

	c.Data["json"] = statusList
	c.ServeJSON()
}
