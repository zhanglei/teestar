package api_controller

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/hsluoyz/gitstar/api"
)

type UsersController struct {
	beego.Controller
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
