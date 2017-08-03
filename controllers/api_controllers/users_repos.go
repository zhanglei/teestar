package api_controllers

import (
	"strings"

	"github.com/hsluoyz/gitstar/api"
)

func getAllUserAndOrganRepos(user string) []string {
	var repos []string

	tokens := strings.Split(user, ",")
	for _, token := range tokens {
		repos = append(repos, api.ListRepos(token)...)
	}

	return repos
}

func (c *UsersController) GetUserAllRepos() {
	user := c.GetString(":user")

	repos := getAllUserAndOrganRepos(user)
	c.Data["json"] = repos
	c.ServeJSON()
}

func (c *UsersController) GetUserRepos() {
	user := c.GetString(":user")

	c.Data["json"] = api.GetUserRepos(user)
	c.ServeJSON()
}

func (c *UsersController) AddUserRepo() {
	user := c.GetString(":user")
	repo := c.GetString(":repo")
	repo = strings.Replace(repo, ".", "/", -1)

	affected := api.AddUserRepo(user, repo)

	if affected {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = "not affected"
	}
	c.ServeJSON()
}

func (c *UsersController) DeleteUserRepo() {
	user := c.GetString(":user")
	repo := c.GetString(":repo")
	repo = strings.Replace(repo, ".", "/", -1)

	affected := api.DeleteUserRepo(user, repo)

	if affected {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = "not affected"
	}
	c.ServeJSON()
}
