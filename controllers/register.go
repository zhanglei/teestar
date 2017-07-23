package controllers

import (
	"errors"
	"strings"

	"github.com/astaxie/beego"
	"github.com/hsluoyz/gitstar/api"
)

type MainController struct {
	beego.Controller
}

var adapter *Adapter

func init() {
	adapter = NewAdapter("mysql", "root:@tcp(127.0.0.1:3306)/")
}

func (c *MainController) GetUserAllRepos() {
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

func (c *MainController) GetUserRepos() {
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

func (c *MainController) AddUserRepo() {
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

func (c *MainController) DeleteUserRepo() {
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
