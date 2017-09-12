package api

import (
	"runtime"
	"sort"
	"sync"
	"time"
)

type ExtendedUserList []*ExtendedUser

func (list ExtendedUserList) Len() int {
	return len(list)
}

func (list ExtendedUserList) Less(i, j int) bool {
	return list[i].CreatedAt < list[j].CreatedAt
}

func (list ExtendedUserList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func GetUser(user string) *User {
	var objUser = User{User: user}
	has, err := adapter.engine.Get(&objUser)
	if err != nil {
		panic(err)
	}

	if has {
		return &objUser
	} else {
		return nil
	}
}

func IsUserAdmin(user string) bool {
	var objUser = User{User: user}
	has, err := adapter.engine.Get(&objUser)
	if err != nil {
		panic(err)
	}

	if has {
		return objUser.IsAdmin
	} else {
		return false
	}
}

func GetExtendedUserFromUser(objUser *User) *ExtendedUser {
	starringCount := GetUserStarringCount(objUser.User)
	starredCount := GetUserStarredCount(objUser.User)

	followStatus := GetUserFollowStatus(objUser.User)
	followingCount := len(followStatus.FollowingTargets)
	followedCount := len(followStatus.FollowedTargets)

	objExtendedUser := ExtendedUser{
		User: objUser.User,
		Hitter: objUser.Hitter,
		QQ: objUser.QQ,
		CreatedAt: objUser.CreatedAt,
		Nickname: objUser.Nickname,
		Email: objUser.Email,
		IsAdmin: objUser.IsAdmin,
		IsDisabled: objUser.IsDisabled,
		IsFollowable: objUser.IsFollowable,
		IsFlagged: objUser.IsFlagged,
		RepoCount: GetUserRepoCount(objUser.User),
		StarringCount: starringCount,
		StarredCount: starredCount,
		OweCount: starredCount - starringCount,
		FollowingCount: followingCount,
		FollowedCount: followedCount,
		FollowOweCount: followedCount - followingCount}

	return &objExtendedUser
}

func GetExtendedUser(user string) *ExtendedUser {
	objUser := GetUser(user)
	if objUser == nil {
		return nil
	}

	return GetExtendedUserFromUser(objUser)
}

func GetUserObjects() []User {
	var objUsers []User
	err := adapter.engine.Asc("created_at").Find(&objUsers)
	if err != nil {
		panic(err)
	}

	return objUsers
}

func GetExtendedUserObjects() ExtendedUserList {
	objUsers := GetUserObjects()
	objExtendedUsers := ExtendedUserList{}

	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex sync.Mutex
	var w sync.WaitGroup
	w.Add(len(objUsers))
	for _, objUser := range objUsers {
		go func(objUser User) {
			objExtendedUser := GetExtendedUserFromUser(&objUser)
			mutex.Lock()
			objExtendedUsers = append(objExtendedUsers, objExtendedUser)
			mutex.Unlock()
			w.Done()
		}(objUser)
	}
	w.Wait()

	runtime.GOMAXPROCS(1)

	sort.Sort(objExtendedUsers)
	return objExtendedUsers
}

func GetUsers() []string {
	var objUsers []User
	err := adapter.engine.Find(&objUsers)
	if err != nil {
		panic(err)
	}

	users := []string{}
	for _, objUser := range objUsers {
		users = append(users, objUser.User)
	}

	return users
}

func GetOtherUsers(user string) []string {
	var objUsers []User
	err := adapter.engine.Find(&objUsers)
	if err != nil {
		panic(err)
	}

	users := []string{}
	for _, objUser := range objUsers {
		if objUser.User != user {
			users = append(users, objUser.User)
		}
	}

	return users
}

func GetActiveUsers() []string {
	var objUsers []User
	err := adapter.engine.Find(&objUsers)
	if err != nil {
		panic(err)
	}

	users := []string{}
	for _, objUser := range objUsers {
		if !objUser.IsDisabled && objUser.QQ != "" {
			users = append(users, objUser.User)
		}
	}

	return users
}

func GetOtherEnabledUsers(user string) []string {
	var objUsers []User
	err := adapter.engine.Find(&objUsers)
	if err != nil {
		panic(err)
	}

	users := []string{}
	for _, objUser := range objUsers {
		if objUser.User != user && !objUser.IsDisabled && !objUser.IsFlagged {
			users = append(users, objUser.User)
		}
	}

	return users
}

func GetFollowableUsers() []string {
	var objUsers []User
	err := adapter.engine.Find(&objUsers)
	if err != nil {
		panic(err)
	}

	users := []string{}
	for _, objUser := range objUsers {
		if !objUser.IsDisabled && objUser.QQ != "" && objUser.IsFollowable && !objUser.IsFlagged {
			users = append(users, objUser.User)
		}
	}

	return users
}

func GetOtherFollowableUsers(user string) []string {
	var objUsers []User
	err := adapter.engine.Find(&objUsers)
	if err != nil {
		panic(err)
	}

	users := []string{}
	for _, objUser := range objUsers {
		if objUser.User != user && !objUser.IsDisabled && objUser.QQ != "" && objUser.IsFollowable && !objUser.IsFlagged {
			users = append(users, objUser.User)
		}
	}

	return users
}

func HasUser(user string) bool {
	var objUsers []User
	err := adapter.engine.Find(&objUsers)
	if err != nil {
		panic(err)
	}

	for _, objUser := range objUsers {
		if objUser.User == user {
			return true
		}
	}

	return false
}

func HasHitter(user string, hitter string) bool {
	var objUsers []User
	err := adapter.engine.Find(&objUsers)
	if err != nil {
		panic(err)
	}

	for _, objUser := range objUsers {
		if objUser.User != user && objUser.Hitter == hitter {
			return true
		}
	}

	return false
}

func CheckUserPassword(user string, password string) bool {
	var objUsers []User
	err := adapter.engine.Find(&objUsers)
	if err != nil {
		panic(err)
	}

	for _, objUser := range objUsers {
		if objUser.User == user && objUser.Password == password {
			return true
		}
	}

	return false
}

func ChangeUserPassword(user string, password string) bool {
	objUser := User{User: user, Password: password}
	affected, err := adapter.engine.Id(user).Cols("password").Update(objUser)
	if err != nil {
		panic(err)
	}

	return affected != 0
}

func getCurrentTime() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04:05")
}

func AddUser(user string, password string) bool {
	objUser := User{User: user, Password: password, Hitter: "", CreatedAt: getCurrentTime(), IsAdmin: false, IsDisabled: false, IsFollowable: false, IsFlagged: false}
	affected, err := adapter.engine.Insert(objUser)
	if err != nil {
		panic(err)
	}

	return affected != 0
}
