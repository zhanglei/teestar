package api

import (
	"runtime"
	"sort"
	"sync"
)

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

	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex sync.Mutex
	var w sync.WaitGroup
	w.Add(len(users))
	for _, user := range users {
		go func(user string) {
			statusList := GetUnsortedUserOwe(user)
			mutex.Lock()
			allStatusList = append(allStatusList, statusList...)
			mutex.Unlock()
			w.Done()
		}(user)
	}
	w.Wait()

	runtime.GOMAXPROCS(1)

	sort.Sort(allStatusList)
	return allStatusList
}
