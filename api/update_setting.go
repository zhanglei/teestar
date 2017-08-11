package api

func UpdateUserHitter(user string, hitter string) bool {
	objUser := User{User: user, Hitter: hitter}
	affected, err := adapter.engine.Id(user).Cols("hitter").Update(objUser)
	if err != nil {
		panic(err)
	}

	return affected != 0
}

func UpdateUserQQ(user string, qq string) bool {
	objUser := User{User: user, QQ: qq}
	affected, err := adapter.engine.Id(user).Cols("q_q").Update(objUser)
	if err != nil {
		panic(err)
	}

	return affected != 0
}

func UpdateUserNickname(user string, nickname string) bool {
	objUser := User{User: user, Nickname: nickname}
	affected, err := adapter.engine.Id(user).Cols("nickname").Update(objUser)
	if err != nil {
		panic(err)
	}

	return affected != 0
}

func UpdateUserEmail(user string, email string) bool {
	objUser := User{User: user, Email: email}
	affected, err := adapter.engine.Id(user).Cols("email").Update(objUser)
	if err != nil {
		panic(err)
	}

	return affected != 0
}

func UpdateUserFollowable(user string, followable bool) bool {
	objUser := User{User: user, IsFollowable: followable}
	affected, err := adapter.engine.Id(user).Cols("is_followable").Update(objUser)
	if err != nil {
		panic(err)
	}

	return affected != 0
}
