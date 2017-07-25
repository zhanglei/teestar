package api

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
