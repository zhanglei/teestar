package controllers

import (
	"errors"
	"strings"

	"github.com/hsluoyz/gitstar/api"
)

func (c *MainController) GetUserTarget() {
	user := c.GetString(":user")
	target := c.GetString(":target")

	targetRepos := getUserRepos(target)
	targetStarredReposPotential := api.ListStarringRepos(user)
	res, ok := Intersect(targetRepos, targetStarredReposPotential)
	if !ok {
		panic(errors.New("cannot find intersect"))
	}
	targetStarredRepos := res.Interface().([]string)

	c.Data["json"] = targetStarredRepos
	c.ServeJSON()
}

type UserTargetStatus struct {
	StarringRepos []string
	StarredRepos  []string
}

func getAllUserAndOrganRepos(user string) []string {
	var repos []string

	tokens := strings.Split(user, ",")
	for _, token := range tokens {
		repos = append(repos, api.ListRepos(token)...)
	}

	return repos
}

func getRealUser(user string) string {
	tokens := strings.Split(user, ",")
	if len(tokens) == 0 {
		panic(errors.New("invalid user"))
	}

	return tokens[0]
}

func (c *MainController) GetUserTargetStatus() {
	user := c.GetString(":user")
	target := c.GetString(":target")

	targetRepos := getAllUserAndOrganRepos(target)
	targetStarredReposPotential := api.ListStarringRepos(getRealUser(user))
	targetRes, ok := Intersect(targetRepos, targetStarredReposPotential)
	if !ok {
		panic(errors.New("cannot find intersect"))
	}
	targetStarredRepos := targetRes.Interface().([]string)

	userRepos := getAllUserAndOrganRepos(user)
	userStarredReposPotential := api.ListStarringRepos(getRealUser(target))
	userRes, ok := Intersect(userRepos, userStarredReposPotential)
	if !ok {
		panic(errors.New("cannot find intersect"))
	}
	userStarredRepos := userRes.Interface().([]string)

	c.Data["json"] = UserTargetStatus{StarringRepos: targetStarredRepos, StarredRepos: userStarredRepos}
	c.ServeJSON()
}
