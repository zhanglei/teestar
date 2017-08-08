package api_controllers

import (
	"strings"

	"github.com/hsluoyz/gitstar/api"
	"github.com/hsluoyz/gitstar/util"
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
	if c.RequireLogin() {
		return
	}

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
	if c.RequireLogin() {
		return
	}

	user := c.GetString(":user")

	c.Data["json"] = api.GetUserRepos(user)
	c.ServeJSON()
}

// @Title GetUserExtendedRepos
// @Description Get all the repos the user added in GitStar with extended information
// @Param   user     path    string  true        "The username"
// @Success 200 {object} api.Repo The Repo object
// @router /:user/repos/extended [get]
func (c *UsersController) GetUserExtendedRepos() {
	if c.RequireLogin() {
		return
	}

	user := c.GetString(":user")

	c.Data["json"] = api.GetUserRepoObjects(user)
	c.ServeJSON()
}

// @Title AddUserRepo
// @Description Add a repo for the user
// @Param   user     path      string  true    "The username"
// @Param   repo     formData  string  true    "The repository name, like user_name/repo_name"
// @Success 200 {object} controllers.api_controller.Response The Response object
// @router /:user/repos/add [post]
func (c *UsersController) AddUserRepo() {
	user := c.GetString(":user")
	if c.RequireUser(user) {
		return
	}

	var resp Response
	repo := c.Input().Get("repo")

	msg := api.CheckAddRepo(user, repo)
	if msg != "" {
		resp = Response{Code: 0, Msg: msg, Data: ""}
	} else {
		api.AddUserRepo(user, repo)

		util.LogInfo(c.Ctx, "API: [%s] added repo: [%s]", user, repo)

		resp = Response{Code: 200, Msg: "添加项目成功", Data: ""}
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

// @Title DeleteUserRepo
// @Description Delete a repo for the user
// @Param   user     path      string  true    "The username"
// @Param   repo     formData  string  true    "The repository name, like user_name/repo_name"
// @Success 200 {object} controllers.api_controller.Response The Response object
// @router /:user/repos/delete [post]
func (c *UsersController) DeleteUserRepo() {
	user := c.GetString(":user")
	if c.RequireUser(user) {
		return
	}

	var resp Response
	repo := c.Input().Get("repo")

	msg := api.CheckDeleteRepo(user, repo)
	if msg != "" {
		resp = Response{Code: 0, Msg: msg, Data: ""}
	} else {
		api.DeleteUserRepo(user, repo)

		util.LogInfo(c.Ctx, "API: [%s] deleted repo: [%s]", user, repo)

		resp = Response{Code: 200, Msg: "删除项目成功", Data: ""}
	}

	c.Data["json"] = resp
	c.ServeJSON()
}
