package controllers

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
