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

// @Title GetUserAllRepos
// @Description Get all the repos the user has in GitHub, this API is actually not used yet
// @Param   user     path    string  true        "The username"
// @Success 200 {[]string}
// @router /:user/repos/all [get]
func (c *UsersController) GetUserAllRepos() {
	user := c.GetString(":user")

	repos := getAllUserAndOrganRepos(user)
	c.Data["json"] = repos
	c.ServeJSON()
}

// @Title GetUserRepos
// @Description Get all the repos the user added in GitStar
// @Param   user     path    string  true        "The username"
// @Success 200 {[]string}
// @router /:user/repos [get]
func (c *UsersController) GetUserRepos() {
	user := c.GetString(":user")

	c.Data["json"] = api.GetUserRepos(user)
	c.ServeJSON()
}

// @Title AddUserRepo
// @Description Add a repo for the user
// @Param   user     path    string  true        "The username"
// @Param   repo     path    string  true        "The repository name, like user_name.repo_name"
// @Success 200 {object} controllers.api_controller.Response The response object
// @router /:user/repos/add/:repo [get]
func (c *UsersController) AddUserRepo() {
	var resp Response
	user := c.GetString(":user")
	repo := c.GetString(":repo")
	repo = strings.Replace(repo, ".", "/", -1)

	msg := api.CheckAddRepo(user, repo)
	if msg != "" {
		resp = Response{Code: 0, Msg: msg, Data: ""}
	} else {
		affected := api.AddUserRepo(user, repo)

		if affected {
			resp = Response{Code: 200, Msg: "ok", Data: ""}
		} else {
			resp = Response{Code: 200, Msg: "not affected", Data: ""}
		}
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

// @Title DeleteUserRepo
// @Description Delete a repo for the user
// @Param   user     path    string  true        "The username"
// @Param   repo     path    string  true        "The repository name, like user_name.repo_name"
// @Success 200 {object} controllers.api_controller.Response The response object
// @router /:user/repos/delete/:repo [get]
func (c *UsersController) DeleteUserRepo() {
	var resp Response
	user := c.GetString(":user")
	repo := c.GetString(":repo")
	repo = strings.Replace(repo, ".", "/", -1)

	msg := api.CheckDeleteRepo(user, repo)
	if msg != "" {
		resp = Response{Code: 0, Msg: msg, Data: ""}
	} else {
		affected := api.DeleteUserRepo(user, repo)

		if affected {
			resp = Response{Code: 200, Msg: "ok", Data: ""}
		} else {
			resp = Response{Code: 200, Msg: "not affected", Data: ""}
		}
	}

	c.Data["json"] = resp
	c.ServeJSON()
}
