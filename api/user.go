package api

import (
	"strings"
	"time"
)

func GetUserObjects() []User {
	var objUsers []User
	err := adapter.engine.Asc("created_at").Find(&objUsers)
	if err != nil {
		panic(err)
	}

	return objUsers
}

func GetExtendedUserObjects() []ExtenedUser {
	objUsers := GetUserObjects()

	objExtendedUsers := []ExtenedUser{}
	for _, objUser := range objUsers {
		starringCount := GetUserStarringCount(objUser.User)
		starredCount := GetUserStarredCount(objUser.User)

		objExtendedUser := ExtenedUser{
			User: objUser.User,
			Hitter: objUser.Hitter,
			QQ: objUser.QQ,
			CreatedAt: objUser.CreatedAt,
			Nickname: objUser.Nickname,
			RepoCount: GetUserRepoCount(objUser.User),
			StarringCount: starringCount,
			StarredCount: starredCount,
			OweCount: starredCount - starringCount}
		objExtendedUsers = append(objExtendedUsers, objExtendedUser)
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
		if objUser.User != user && objUser.IsDisabled == false {
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

func CheckUserRegister(user string, password string) string {
	if len(user) == 0 || len(password) == 0 {
		return "用户名或密码不能为空"
	} else if HasUser(user) {
		return "用户名已被注册"
	} else if HasHitter("", user) {
		return "用户名已被其他用户注册为点赞小号"
	} else if strings.Contains(user, "@") {
		return "请不要使用邮箱，GitHub profile（如https://github.com/abc）中，abc是用户名"
	} else if !HasGitHubUser(user) {
		return "用户名不是合法的、已存在的GitHub用户名"
	} else {
		return ""
	}
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
