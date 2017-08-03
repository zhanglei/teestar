package api_controller

import "github.com/hsluoyz/gitstar/api"

func (c *APIController) GetUserStarringRepos() {
	user := c.GetString(":user")

	c.Data["json"] = api.GetUserStarringRepos(user)
	c.ServeJSON()
}

func (c *APIController) UpdateUserStarringRepos() {
	user := c.GetString(":user")

	affected := api.UpdateUserStarringRepos(user)

	if affected {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = "not affected"
	}
	c.ServeJSON()
}

func (c *APIController) GetUserHitter() {
	user := c.GetString(":user")

	hitter := api.GetUserHitter(user)

	c.Data["json"] = hitter
	c.ServeJSON()
}

func (c *APIController) UpdateUserHitter() {
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
