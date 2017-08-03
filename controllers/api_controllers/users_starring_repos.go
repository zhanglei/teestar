package api_controllers

import "github.com/hsluoyz/gitstar/api"

func (c *UsersController) GetUserStarringRepos() {
	user := c.GetString(":user")

	c.Data["json"] = api.GetUserStarringRepos(user)
	c.ServeJSON()
}

func (c *UsersController) UpdateUserStarringRepos() {
	user := c.GetString(":user")

	affected := api.UpdateUserStarringRepos(user)

	if affected {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = "not affected"
	}
	c.ServeJSON()
}
