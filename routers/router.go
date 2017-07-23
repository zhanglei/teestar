package routers

import (
	"github.com/astaxie/beego"
	"github.com/hsluoyz/gitstar/controllers/api_controller"
)

func init() {
	initAPI()
}

func initAPI() {
	beego.Router("/api/users", &api_controller.APIController{}, "get:GetUsers")

    beego.Router("/api/users/:user/all-repos", &api_controller.APIController{},"get:GetUserAllRepos")
	beego.Router("/api/users/:user/repos", &api_controller.APIController{}, "get:GetUserRepos")
	beego.Router("/api/users/:user/repos/add/:repo", &api_controller.APIController{}, "get:AddUserRepo")
	beego.Router("/api/users/:user/repos/delete/:repo", &api_controller.APIController{}, "get:DeleteUserRepo")

	beego.Router("/api/users/:user/starring-repos", &api_controller.APIController{}, "get:GetUserStarringRepos")
	beego.Router("/api/users/:user/starring-repos/update", &api_controller.APIController{}, "get:UpdateUserStarringRepos")
	beego.Router("/api/users/:user/hitter", &api_controller.APIController{}, "get:GetUserHitter")
	beego.Router("/api/users/:user/hitter/update/:hitter", &api_controller.APIController{}, "get:UpdateUserHitter")

	beego.Router("/api/users/:user/targets/:target", &api_controller.APIController{}, "get:GetUserTarget")
	beego.Router("/api/users/:user/targets/:target/status", &api_controller.APIController{}, "get:GetUserTargetStatus")
	beego.Router("/api/users/:user/targets/:target/pool", &api_controller.APIController{}, "get:GetUserTargetPool")
	beego.Router("/api/users/:user/status", &api_controller.APIController{}, "get:GetUserStatus")
	beego.Router("/api/users/:user/recommend", &api_controller.APIController{}, "get:GetUserRecommend")
}
