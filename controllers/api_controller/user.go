package api_controller

func getUsers() []string {
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

func getOtherUsers(user string) []string {
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

func (c *APIController) GetUsers() {
	users := getUsers()
	c.Data["json"] = users
	c.ServeJSON()
}
