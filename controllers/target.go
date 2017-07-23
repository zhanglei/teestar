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

func getSubtract(a []string, b []string) []string {
	res := []string{}
	for _, ia := range a {
		found := false
		for _, ib := range b {
			if ia == ib {
				found = true
				break
			}
		}

		if !found {
			res = append(res, ia)
		}
	}

	return res
}

func Get(a []string, b []string) []string {
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

type UserTargetStatus struct {
	StarringRepos []string
	StarredRepos  []string
	Score         int
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

	score := len(starredRepos) - len(starringRepos)

	c.Data["json"] = UserTargetStatus{StarringRepos: starringRepos, StarredRepos: starredRepos, Score: score}
	c.ServeJSON()
}

func (c *MainController) GetUserTargetPool() {
	user := c.GetString(":user")
	target := c.GetString(":target")

	targetRepos := getUserRepos(target)
	userStarringRepos := getUserStarringRepos(user)
	poolRepos := getSubtract(targetRepos, userStarringRepos)

	c.Data["json"] = poolRepos
	c.ServeJSON()
}
