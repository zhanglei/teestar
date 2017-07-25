package api

func GetUserStarringRepos(user string) []string {
	hitter := GetUserHitter(user)
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

func UpdateUserStarringRepos(user string) bool {
	hitter := GetUserHitter(user)
	if hitter != "" {
		user = hitter
	}

	affected, err := adapter.engine.Delete(&UserStarringRepo{User: user})
	if err != nil {
		panic(err)
	}

	repos := ListStarringRepos(user)
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

func GetUser(user string) *User {
	var objUser = User{User: user}
	has, err := adapter.engine.Get(&objUser)
	if err != nil {
		panic(err)
	}

	if has {
		return &objUser
	} else {
		return nil
	}
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

func UpdateUserHitter(user string, hitter string) bool {
	objUser := User{User: user, Hitter: hitter}
	affected, err := adapter.engine.Id(user).Cols("hitter").Update(objUser)
	if err != nil {
		panic(err)
	}

	return affected != 0
}

func UpdateUserQQ(user string, qq string) bool {
	objUser := User{User: user, QQ: qq}
	affected, err := adapter.engine.Id(user).Cols("q_q").Update(objUser)
	if err != nil {
		panic(err)
	}

	return affected != 0
}
