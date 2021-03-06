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
		return "请不要使用邮箱，需要使用GitHub用户名。在GitHub profile（如https://github.com/abc）中，abc是用户名"
	} else if !HasGiteeUser(user) {
		return "用户名不是合法的、已存在的GitHub用户名"
	} else if !IsGiteeUserOldEnough(user) {
		return "GitHub账号注册时间需要至少满30天"
	} else if !IsGiteeUserStarringRepo(user, "Sable/abc") {
		return "为了验证你对所填GitHub账号的所有权，请Star这个仓库：https://github.com/Sable/abc，然后再点击注册。注册成功后，可以取消该点赞"
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

func CheckUserUpdateHitter(user string, hitter string) string {
	if hitter == "" {
		return ""
	} else if hitter == user {
		return "不需要把点赞账号（小号）设置为与用户名（大号）一致，留空即表示用大号点赞"
	} else if HasUser(hitter) {
		return "点赞账号与其他用户的用户名（大号）重复，无法使用"
	} else if HasHitter(user, hitter) {
		return "点赞账号与其他用户的点赞账号（小号）重复，无法使用"
	} else if !HasGiteeUser(hitter) {
		return "点赞账号不是合法的、已存在的GitHub用户名"
	} else if !IsGiteeUserActive(hitter) {
		return "点赞账号至少要有3个仓库，并且其中至少有1个仓库是非fork的，请完善该小号后再设置"
	} else {
		return ""
	}
}

func CheckUserUpdateQQ(qq string) string {
	if qq == "" {
		return "请填写QQ号"
	} else if !HasQQUser(qq) {
		return "QQ号不是合法的、已存在的QQ号码"
	} else {
		return ""
	}
}

func CheckUserChangePassword(user string, oldPassword string, newPassword string) string {
	if !HasUser(user) {
		return "用户不存在"
	} else if !CheckUserPassword(user, oldPassword) {
		return "旧密码错误"
	} else if oldPassword == newPassword {
		return "新密码不能与旧密码一致"
	} else if newPassword == "" {
		return "新密码不能为空"
	} else {
		return ""
	}
}

func CheckAddRepo(user string, repo string) string {
	if !HasUser(user) {
		return "用户不存在"
	} else if GetUser(user).QQ == "" {
		return "填写QQ号后才能添加项目"
	} else if len(repo) == 0 {
		return "项目不能为空"
	} else if HasUserRepo(user, repo) {
		return "该项目已经存在"
	} else if HasRepo(repo) {
		return "该项目已经被其他用户添加，如有疑问请联系管理员"
	} else if !HasGiteeRepo(repo) {
		return "项目地址不是合法的、已存在的GitHub项目地址"
	} else {
		return ""
	}
}

func CheckEnableRepo(user string, repo string) string {
	if !HasUser(user) {
		return "用户不存在"
	} else if len(repo) == 0 {
		return "项目不能为空"
	} else if !HasUserRepo(user, repo) {
		return "该项目不存在"
	} else {
		return ""
	}
}

func CheckDeleteRepo(user string, repo string) string {
	if !HasUser(user) {
		return "用户不存在"
	} else if len(repo) == 0 {
		return "项目不能为空"
	} else if !HasUserRepo(user, repo) {
		return "该项目不存在"
	} else if CanPayOff(user) >= 10 {
		return "未还欠赞小于10次时，才可以隐藏、删除项目，请及时还清欠赞"
	} else {
		return ""
	}
}
