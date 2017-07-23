package controllers

import (
	"errors"
)

func (c *MainController) GetUserTarget() {
	user := c.GetString(":user")
	target := c.GetString(":target")

	targetRepos := getUserRepos(target)
	targetStarredReposPotential := getUserStarringRepos(user)
	res, ok := Intersect(targetRepos, targetStarredReposPotential)
	if !ok {
		panic(errors.New("cannot find intersect"))
	}
	targetStarredRepos := res.Interface().([]string)

	c.Data["json"] = targetStarredRepos
	c.ServeJSON()
}
