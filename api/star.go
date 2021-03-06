package api

func GetUserStarringRepos(user string) []string {
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

func UpdateUserStarringRepos(user string) bool {
	hitter := GetUserOrHitter(user)

	affected, err := adapter.engine.Delete(&UserStarringRepo{User: user})
	if err != nil {
		panic(err)
	}

	flagged := IsGiteeUserFlagged(hitter)
	UpdateUserFlagged(user, flagged)
	if flagged {
		return affected != 0
	}

	repos := ListStarringRepos(hitter)
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

func UpdateStarringRepos() bool {
	affected := false
	users := GetActiveUsers()

	for _, user := range users {
		if UpdateUserStarringRepos(user) {
			affected = true
		}
	}

	return affected
}

func GetUserHitter(user string) string {
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

func GetUserOrHitter(user string) string {
	hitter := GetUserHitter(user)
	if hitter == "" {
		hitter = user
	}

	return hitter
}
