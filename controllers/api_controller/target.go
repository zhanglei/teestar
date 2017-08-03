package api_controller

import "github.com/hsluoyz/gitstar/api"

func (c *UsersController) GetUserTarget() {
	user := c.GetString(":user")
	target := c.GetString(":target")

	targetRepos := api.GetUserRepos(target)
	userStarringRepos := api.GetUserStarringRepos(user)
	c.Data["json"] = api.GetIntersect(targetRepos, userStarringRepos)
	c.ServeJSON()
}

func (c *UsersController) GetUserTargetStatus() {
	user := c.GetString(":user")
	target := c.GetString(":target")

	c.Data["json"] = api.GetUserTargetStatus(user, target)
	c.ServeJSON()
}

func (c *UsersController) GetUserTargetPool() {
	user := c.GetString(":user")
	target := c.GetString(":target")

	c.Data["json"] = api.GetUserTargetPool(user, target)
	c.ServeJSON()
}

func (c *UsersController) GetUserStatus() {
	user := c.GetString(":user")

	statusList := api.GetUserStatus(user)

	c.Data["json"] = statusList
	c.ServeJSON()
}

func (c *UsersController) GetUserRecommend() {
	user := c.GetString(":user")

	c.Data["json"] = api.GetUserRecommend(user)
	c.ServeJSON()
}
