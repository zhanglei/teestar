package api

func GetAllRepos() []string {
	var userRepos []UserRepo
	err := adapter.engine.Find(&userRepos)
	if err != nil {
		panic(err)
	}

	repos := []string{}
	for _, userRepo := range userRepos {
		repos = append(repos, userRepo.Repo)
	}

	return repos
}

func GetUserRepos(user string) []string {
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

func GetUserRepoCount(user string) int {
	userRepo := new(UserRepo)
	total, err := adapter.engine.Where("user = ?", user).Count(userRepo)
	if err != nil {
		panic(err)
	}

	return int(total)
}

func HasUserRepo(user string, repo string) bool {
	userRepo := UserRepo{User: user, Repo: repo}
	has, err := adapter.engine.Get(&userRepo)
	if err != nil {
		panic(err)
	}

	return has
}

func AddUserRepo(user string, repo string) bool {
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

func DeleteUserRepo(user string, repo string) bool {
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

func EnableUserRepo(user string, repo string, enable bool) bool {
	objUserRepo := UserRepo{User: user, Repo: repo, IsDisabled: !enable}
	affected, err := adapter.engine.Where("user=? and repo=?", user, repo).Cols("is_disabled").Update(objUserRepo)
	if err != nil {
		panic(err)
	}

	return affected != 0
}
