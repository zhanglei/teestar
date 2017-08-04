package api

import "strings"

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

func CheckUserLogin(user string, password string) string {
	if !HasUser(user) {
		return "用户名不存在，请先注册"
	} else if !CheckUserPassword(user, password) {
		return "密码错误"
	} else {
		return ""
	}
}
