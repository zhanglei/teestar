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

	userRepos := getUserRepos(target)
	userStarringRepos := getUserStarringRepos(user)
	c.Data["json"] = getIntersect(userRepos, userStarringRepos)
	c.ServeJSON()
}
