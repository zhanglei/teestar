package api

import "time"

func GetUserObjects() []User {
	var objUsers []User
	err := adapter.engine.Find(&objUsers)
	if err != nil {
		panic(err)
	}

	return objUsers
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
