package api_controller

import (
	"strings"

	"github.com/astaxie/beego"
)

type APIController struct {
	beego.Controller
}

var adapter *Adapter

func init() {
	adapter = NewAdapter("mysql", "root:@tcp(127.0.0.1:3306)/")
}

func (c *APIController) GetUserAllRepos() {
	user := c.GetString(":user")

	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"

	repos := getAllUserAndOrganRepos(user)
	c.Data["json"] = repos
	c.ServeJSON()
}

func getUserRepos(user string) []string {
	var userRepos []UserRepo
	err := adapter.engine.Find(&userRepos, &UserRepo{User: user})
	if err != nil {
		panic(err)
	}

	repos := []string{}
	for _, userRepo := range userRepos {
		repos = append(repos, userRepo.Repo)
	}

	return repos
}

func (c *APIController) GetUserRepos() {
	user := c.GetString(":user")

	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"

	c.Data["json"] = getUserRepos(user)
	c.ServeJSON()
}

func addUserRepo(user string, repo string) bool {
	userRepo := UserRepo{User: user, Repo: repo}
	has, err := adapter.engine.Get(&userRepo)
	if err != nil {
		panic(err)
	}

	if has {
		return false
	}

	affected, err := adapter.engine.Insert(userRepo)
	if err != nil {
		panic(err)
	}

	return affected != 0
}

func (c *APIController) AddUserRepo() {
	user := c.GetString(":user")
	repo := c.GetString(":repo")
	repo = strings.Replace(repo, ".", "/", -1)

	affected := addUserRepo(user, repo)

	if affected {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = "not affected"
	}
	c.ServeJSON()
}

func deleteUserRepo(user string, repo string) bool {
	userRepo := UserRepo{User: user, Repo: repo}
	has, err := adapter.engine.Get(&userRepo)
	if err != nil {
		panic(err)
	}

	if !has {
		return false
	}

	affected, err := adapter.engine.Delete(&userRepo)
	if err != nil {
		panic(err)
	}

	return affected != 0
}

func (c *APIController) DeleteUserRepo() {
	user := c.GetString(":user")
	repo := c.GetString(":repo")
	repo = strings.Replace(repo, ".", "/", -1)

	affected := deleteUserRepo(user, repo)

	if affected {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = "not affected"
	}
	c.ServeJSON()
}
