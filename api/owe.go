package api

import "sort"

func GetUserOwe(user string) StatusList {
	statusList := StatusList{}
	otherUsers := GetOtherUsers(user)
	for _, otherUser := range otherUsers {
		status := GetUserTargetStatus(user, otherUser)
		status.Score = -status.Score

		if status.Score <= 0 {
			continue
		}

		if len(status.CanBeStarredRepos) == 0 {
			continue
		}

		statusList = append(statusList, &status)
	}

	sort.Sort(statusList)
	return statusList
}
