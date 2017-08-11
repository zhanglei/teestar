package api

import (
	"time"
)

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

func GetExtendedUser(user string) *ExtendedUser {
	objUser := GetUser(user)
	if objUser == nil {
		return nil
	}

	starringCount := GetUserStarringCount(objUser.User)
	starredCount := GetUserStarredCount(objUser.User)

	objExtendedUser := ExtendedUser{
		User: objUser.User,
		Hitter: objUser.Hitter,
		QQ: objUser.QQ,
		CreatedAt: objUser.CreatedAt,
		Nickname: objUser.Nickname,
		Email: objUser.Email,
		IsAdmin: objUser.IsAdmin,
		IsDisabled: objUser.IsDisabled,
		RepoCount: GetUserRepoCount(objUser.User),
		StarringCount: starringCount,
		StarredCount: starredCount,
		OweCount: starredCount - starringCount}

	return &objExtendedUser
}

func GetUserObjects() []User {
	var objUsers []User
	err := adapter.engine.Asc("created_at").Find(&objUsers)
	if err != nil {
		panic(err)
	}

	return objUsers
}

func GetExtendedUserObjects() []ExtendedUser {
	objUsers := GetUserObjects()

	objExtendedUsers := []ExtendedUser{}
	for _, objUser := range objUsers {
		objExtendedUsers = append(objExtendedUsers, *GetExtendedUser(objUser.User))
	}

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

func GetOtherEnabledUsers(user string) []string {
	var objUsers []User
	err := adapter.engine.Find(&objUsers)
	if err != nil {
		panic(err)
	}

	users := []string{}
	for _, objUser := range objUsers {
		if objUser.User != user && !objUser.IsDisabled {
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
		if objUser.User != user && !objUser.IsDisabled && objUser.IsFollowable {
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
	objUser := User{User: user, Password: password, Hitter: "", CreatedAt: getCurrentTime(), IsAdmin: false}
	affected, err := adapter.engine.Insert(objUser)
	if err != nil {
		panic(err)
	}

	return affected != 0
}
