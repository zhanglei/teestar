package api

import "time"

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

func getCurrentTime() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 03:04:05 PM")
}

func AddUser(user string) bool {
	objUser := User{User: user, Hitter: user, CreatedAt: getCurrentTime()}
	affected, err := adapter.engine.Insert(objUser)
	if err != nil {
		panic(err)
	}

	return affected != 0
}
