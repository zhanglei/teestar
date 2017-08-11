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
	users := GetUsers()

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

func UpdateUserNickname(user string, nickname string) bool {
	objUser := User{User: user, Nickname: nickname}
	affected, err := adapter.engine.Id(user).Cols("nickname").Update(objUser)
	if err != nil {
		panic(err)
	}

	return affected != 0
}

func UpdateUserEmail(user string, email string) bool {
	objUser := User{User: user, Email: email}
	affected, err := adapter.engine.Id(user).Cols("email").Update(objUser)
	if err != nil {
		panic(err)
	}

	return affected != 0
}

func UpdateUserFollowable(user string, followable bool) bool {
	objUser := User{User: user, IsFollowable: followable}
	affected, err := adapter.engine.Id(user).Cols("is_followable").Update(objUser)
	if err != nil {
		panic(err)
	}

	return affected != 0
}
