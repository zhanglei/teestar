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

	hitter := GetUserHitter(user)
	if hitter == "" {
		hitter = user
	}

	count := 0
	for _, repo := range userRepos {
		total, err := adapter.engine.Where("repo = ?", repo).And("user != ?", hitter).Count(userStarringRepo)
		if err != nil {
			panic(err)
		}

		count += int(total)
	}

	return count
}
