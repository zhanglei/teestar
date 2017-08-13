package api

import "sort"

func GetUnsortedUserOwe(user string) StatusList {
	statusList := StatusList{}
	otherUsers := GetOtherEnabledUsers(user)
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

	return statusList
}

func GetUserOwe(user string) StatusList {
	statusList := GetUnsortedUserOwe(user)
	sort.Sort(statusList)
	return statusList
}

func GetOwe() StatusList {
	users := GetUsers()
	allStatusList := StatusList{}

	for _, user := range users {
		statusList := GetUnsortedUserOwe(user)
		allStatusList = append(allStatusList, statusList...)
	}

	sort.Sort(allStatusList)
	return allStatusList
}
