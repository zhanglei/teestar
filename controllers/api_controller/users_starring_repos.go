package api_controller

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

func (c *UsersController) GetUserHitter() {
	user := c.GetString(":user")

	hitter := api.GetUserHitter(user)

	c.Data["json"] = hitter
	c.ServeJSON()
}

func (c *UsersController) UpdateUserHitter() {
	user := c.GetString(":user")
	hitter := c.GetString(":hitter")

	affected := api.UpdateUserHitter(user, hitter)

	if affected {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = "not affected"
	}
	c.ServeJSON()
}
