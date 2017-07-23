package routers

import (
	"github.com/astaxie/beego"
	"github.com/hsluoyz/gitstar/controllers"
)

func init() {
	beego.Router("/api/users", &controllers.MainController{}, "get:GetUsers")

    beego.Router("/api/users/:user/all-repos", &controllers.MainController{},"get:GetUserAllRepos")
	beego.Router("/api/users/:user/repos", &controllers.MainController{}, "get:GetUserRepos")
	beego.Router("/api/users/:user/repos/add/:repo", &controllers.MainController{}, "get:AddUserRepo")
	beego.Router("/api/users/:user/repos/delete/:repo", &controllers.MainController{}, "get:DeleteUserRepo")

	beego.Router("/api/users/:user/starring-repos", &controllers.MainController{}, "get:GetUserStarringRepos")
	beego.Router("/api/users/:user/starring-repos/update", &controllers.MainController{}, "get:UpdateUserStarringRepos")
	beego.Router("/api/users/:user/hitter", &controllers.MainController{}, "get:GetUserHitter")
	beego.Router("/api/users/:user/hitter/update/:hitter", &controllers.MainController{}, "get:UpdateUserHitter")

	beego.Router("/api/users/:user/targets/:target", &controllers.MainController{}, "get:GetUserTarget")
	beego.Router("/api/users/:user/targets/:target/status", &controllers.MainController{}, "get:GetUserTargetStatus")
	beego.Router("/api/users/:user/targets/:target/pool", &controllers.MainController{}, "get:GetUserTargetPool")
}
