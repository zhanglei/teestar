package api

func GetUserStarringCount(user string) int {
	githubStarringRepos := GetUserStarringRepos(user)
	userRepos := GetUserRepos(user)
	allRepos := GetAllRepos()
	otherRepos := GetSubtract(allRepos, userRepos)
	starringRepos := GetIntersect(githubStarringRepos, otherRepos)
	return len(starringRepos)
}

func GetUserStarredCount(user string) int {
	userRepos := GetUserRepos(user)
	userStarringRepo := new(UserStarringRepo)

	count := 0
	for _, repo := range userRepos {
		total, err := adapter.engine.Where("repo = ?", repo).And("user != ?", user).Count(userStarringRepo)
		if err != nil {
			panic(err)
		}

		count += int(total)
	}

	return count
}

func GetRepoStargazers(user string, repo string) []string {
	var userRepos []UserStarringRepo
	adapter.engine.Where("repo = ?", repo).And("user != ?", user).Find(&userRepos)

	stargazers := []string{}
	for _, userRepo := range userRepos {
		stargazers = append(stargazers, userRepo.User)
	}

	return stargazers
}

func GetExtendedUserRepoObjects(user string) []ExtendedUserRepo {
	objExtendedUserRepos := []ExtendedUserRepo{}
	objUserRepos := GetUserRepoObjects(user)

	for _, objUserRepo := range objUserRepos {
		stargazers := GetRepoStargazers(user, objUserRepo.Repo)
		objRepo := ExtendedUserRepo{
			User: objUserRepo.User,
			Repo: objUserRepo.Repo,
			IsDisabled: objUserRepo.IsDisabled,
			Stargazers: stargazers}
		objExtendedUserRepos = append(objExtendedUserRepos, objRepo)
	}

	return objExtendedUserRepos
}
