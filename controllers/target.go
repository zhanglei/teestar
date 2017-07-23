package controllers

import (
	"errors"
)

func getIntersect(a []string, b []string) []string {
	if len(a) == 0 || len(b) == 0 {
		return []string{}
	}

	res, ok := Intersect(a, b)
	if !ok {
		panic(errors.New("cannot find intersect"))
	}

	return res.Interface().([]string)
}

func (c *MainController) GetUserTarget() {
	user := c.GetString(":user")
	target := c.GetString(":target")

	targetRepos := getUserRepos(target)
	userStarringRepos := getUserStarringRepos(user)
	c.Data["json"] = getIntersect(targetRepos, userStarringRepos)
	c.ServeJSON()
}

func (c *MainController) GetUserTargetStatus() {
	user := c.GetString(":user")
	target := c.GetString(":target")

	targetRepos := getUserRepos(target)
	userStarringRepos := getUserStarringRepos(user)
	starringRepos := getIntersect(targetRepos, userStarringRepos)

	userRepos := getUserRepos(user)
	targetStarringRepos := getUserStarringRepos(target)
	starredRepos := getIntersect(userRepos, targetStarringRepos)

	c.Data["json"] = UserTargetStatus{StarringRepos: starringRepos, StarredRepos: starredRepos}
	c.ServeJSON()
}
