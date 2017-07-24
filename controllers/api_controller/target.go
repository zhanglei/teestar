package api_controller

import "github.com/hsluoyz/gitstar/api"

func (c *APIController) GetUserTarget() {
	user := c.GetString(":user")
	target := c.GetString(":target")

	targetRepos := api.GetUserRepos(target)
	userStarringRepos := api.GetUserStarringRepos(user)
	c.Data["json"] = api.GetIntersect(targetRepos, userStarringRepos)
	c.ServeJSON()
}

func (c *APIController) GetUserTargetStatus() {
	user := c.GetString(":user")
	target := c.GetString(":target")

	c.Data["json"] = api.GetUserTargetStatus(user, target)
	c.ServeJSON()
}

func (c *APIController) GetUserTargetPool() {
	user := c.GetString(":user")
	target := c.GetString(":target")

	c.Data["json"] = api.GetUserTargetPool(user, target)
	c.ServeJSON()
}

func (c *APIController) GetUserStatus() {
	user := c.GetString(":user")

	statusList := api.GetUserStatus(user)

	c.Data["json"] = statusList
	c.ServeJSON()
}

func (c *APIController) GetUserRecommend() {
	user := c.GetString(":user")

	repos := []string{}
	statusList := api.GetUserStatus(user)
	for _, status := range statusList {
		repos = append(repos, status.CanStarRepos...)
	}

	c.Data["json"] = repos
	c.ServeJSON()
}
