package api

import "sort"

func GetIntersect(a []string, b []string) []string {
	res := []string{}
	for _, ia := range a {
		found := false
		for _, ib := range b {
			if ia == ib {
				found = true
				break
			}
		}

		if found {
			res = append(res, ia)
		}
	}

	return res
}

func GetSubtract(a []string, b []string) []string {
	res := []string{}
	for _, ia := range a {
		found := false
		for _, ib := range b {
			if ia == ib {
				found = true
				break
			}
		}

		if !found {
			res = append(res, ia)
		}
	}

	return res
}

func GetUserTargetStatus(user string, target string) UserTargetStatus {
	targetRepos := GetUserRepos(target)
	userStarringRepos := GetUserStarringRepos(user)
	starringRepos := GetIntersect(targetRepos, userStarringRepos)

	userRepos := GetUserRepos(user)
	targetStarringRepos := GetUserStarringRepos(target)
	starredRepos := GetIntersect(userRepos, targetStarringRepos)

	score := len(starredRepos) - len(starringRepos)

	canStarRepos := GetSubtract(targetRepos, userStarringRepos)

	return UserTargetStatus{Target: target, StarringRepos: starringRepos, StarredRepos: starredRepos, Score: score, CanStarRepos: canStarRepos}
}

func GetUserTargetPool(user string, target string) []string {
	targetRepos := GetUserRepos(target)
	userStarringRepos := GetUserStarringRepos(user)
	poolRepos := GetSubtract(targetRepos, userStarringRepos)
	return poolRepos
}

func GetUserStatus(user string) StatusList {
	statusList := StatusList{}
	otherUsers := GetOtherUsers(user)
	for _, otherUser := range otherUsers {
		status := GetUserTargetStatus(user, otherUser)
		statusList = append(statusList, &status)
	}

	sort.Sort(statusList)
	return statusList
}

type Entry struct {
	Target string
	Repo   string
}

func GetUserRecommend(user string) []Entry {
	entries := []Entry{}
	statusList := GetUserStatus(user)
	for _, status := range statusList {
		for _, repo := range status.CanStarRepos {
			entries = append(entries, Entry{Target: status.Target, Repo: repo})
		}
	}

	return entries
}
