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
	objTarget := GetUser(target)
	qq := objTarget.QQ
	nickname := objTarget.Nickname
	hitter := objTarget.Hitter

	targetRepos := GetUserEnabledRepos(target)
	userStarringRepos := GetUserStarringRepos(user)
	starringRepos := GetIntersect(targetRepos, userStarringRepos)

	userRepos := GetUserEnabledRepos(user)
	targetStarringRepos := GetUserStarringRepos(target)
	starredRepos := GetIntersect(userRepos, targetStarringRepos)

	score := len(starredRepos) - len(starringRepos)

	canStarRepos := GetSubtract(targetRepos, userStarringRepos)
	canBeStarredRepos := GetSubtract(userRepos, targetStarringRepos)

	return UserTargetStatus{User: user, Target: target, QQ: qq, Nickname: nickname, Hitter: hitter, StarringRepos: starringRepos, StarredRepos: starredRepos, Score: score, CanStarRepos: canStarRepos, CanBeStarredRepos: canBeStarredRepos}
}

func CanPayOff(user string) int {
	canPayOff := 0
	otherUsers := GetOtherEnabledUsers(user)
	for _, otherUser := range otherUsers {
		status := GetUserTargetStatus(user, otherUser)
		if status.Score > 0 && len(status.CanStarRepos) > 0 {
			min := status.Score
			if len(status.CanStarRepos) < min {
				min = len(status.CanStarRepos)
			}
			canPayOff += min
		}
	}

	return canPayOff
}

func GetUserStatus(user string) StatusList {
	statusList := StatusList{}
	otherUsers := GetOtherEnabledUsers(user)
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
	Score  int
	ScoreR int
}

func GetUserRecommend(user string) []Entry {
	entries := []Entry{}
	statusList := GetUserStatus(user)
	for _, status := range statusList {
		for _, repo := range status.CanStarRepos {
			entries = append(entries, Entry{Target: status.Target, Repo: repo, Score: -status.Score, ScoreR: status.Score})
		}
	}

	return entries
}

type Entry2 struct {
	User         string
	QQ           string
	Nickname     string
	Target       string
	CanStarRepos []string
	Score        int
}

func GetRecommend() []Entry2 {
	entries := []Entry2{}

	objUsers := GetUserObjects()
	for _, objUser := range objUsers {
		repos := GetUserEnabledRepos(objUser.User)
		if len(repos) == 0 {
			continue
		}

		statusList := GetUserStatus(objUser.User)
		for _, status := range statusList {
			if len(status.CanStarRepos) != 0 {
				entries = append(entries, Entry2{User: objUser.User, QQ: objUser.QQ, Nickname: objUser.Nickname, Target: status.Target, CanStarRepos: status.CanStarRepos, Score: status.Score})
			}
		}
	}

	return entries
}
