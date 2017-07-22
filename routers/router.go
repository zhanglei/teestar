package routers

import (
	"github.com/astaxie/beego"
	"github.com/hsluoyz/gitstar/controllers"
)

func init() {
    beego.Router("/api/users/:user/all-repos", &controllers.MainController{},"get:GetUserAllRepos")
	beego.Router("/api/users/:user/repos", &controllers.MainController{}, "get:GetUserRepos")
	beego.Router("/api/users/:user/repos/add/:repo", &controllers.MainController{}, "get:AddUserRepo")
	beego.Router("/api/users/:user/repos/delete/:repo", &controllers.MainController{}, "get:DeleteUserRepo")

	beego.Router("/api/users/:user/targets/:target", &controllers.MainController{}, "get:GetUserTarget")
}
