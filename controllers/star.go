package controllers

import "github.com/hsluoyz/gitstar/api"

func getUserStarringRepos(user string) []string {
	hitter := getUserHitter(user)
	if hitter != "" {
		user = hitter
	}

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
	hitter := getUserHitter(user)
	if hitter != "" {
		user = hitter
	}

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

func getUserHitter(user string) string {
	var objUser = User{User: user}
	has, err := adapter.engine.Get(&objUser)
	if err != nil {
		panic(err)
	}

	if has {
		return objUser.Hitter
	} else {
		return ""
	}
}

func (c *MainController) GetUserHitter() {
	user := c.GetString(":user")

	hitter := getUserHitter(user)

	c.Data["json"] = hitter
	c.ServeJSON()
}

func updateUserHitter(user string, hitter string) bool {
	objUser := User{User: user, Hitter: hitter}
	affected, err := adapter.engine.Id(user).Update(objUser)
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
