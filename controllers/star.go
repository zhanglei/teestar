package controllers

import "github.com/hsluoyz/gitstar/api"

func getUserStarringRepos(user string) []string {
	var userRepos []UserStarringRepo
	err := adapter.engine.Find(&userRepos, &UserStarringRepo{User: user})
	if err != nil {
		panic(err)
	}

	repos := []string{}
	for _, userRepo := range userRepos {
		repos = append(repos, userRepo.Repo)
	}

	return repos
}

func (c *MainController) GetUserStarringRepos() {
	user := c.GetString(":user")

	c.Data["json"] = getUserStarringRepos(user)
	c.ServeJSON()
}

func updateUserStarringRepos(user string) bool {
	affected, err := adapter.engine.Delete(&UserStarringRepo{User: user})
	if err != nil {
		panic(err)
	}

	repos := api.ListStarringRepos(user)
	userRepos := []UserStarringRepo{}
	for _, repo := range repos {
		userRepos = append(userRepos, UserStarringRepo{User: user, Repo: repo})
	}

	affected, err = adapter.engine.Insert(&userRepos)
	if err != nil {
		panic(err)
	}

	return affected != 0
}

func (c *MainController) UpdateUserStarringRepos() {
	user := c.GetString(":user")

	affected := updateUserStarringRepos(user)

	if affected {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = "not affected"
	}
	c.ServeJSON()
}

func updateUserHitter(user string, hitter string) bool {
	affected, err := adapter.engine.Delete(&UserHitter{User: user})
	if err != nil {
		panic(err)
	}

	userHitter := UserHitter{User: user, Hitter: hitter}
	affected, err = adapter.engine.Insert(userHitter)
	if err != nil {
		panic(err)
	}

	return affected != 0
}

func (c *MainController) UpdateUserHitter() {
	user := c.GetString(":user")
	hitter := c.GetString(":hitter")

	affected := updateUserHitter(user, hitter)

	if affected {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = "not affected"
	}
	c.ServeJSON()
}
