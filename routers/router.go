// @APIVersion 1.0.0
// @Title GitStar RESTful API
// @Description This is the RESTful API for GitStar.cn.
// @Contact hsluoyz@qq.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/hsluoyz/gitstar/controllers/api_controller"
	"github.com/hsluoyz/gitstar/controllers/view_controller"
)

func init() {
	// beego.InsertFilter("/*", beego.BeforeRouter, FilterIP)

	initCrossSite()

	initAPI()
	initView()

	beego.SetStaticPath("/swagger", "swagger")
}

func initCrossSite() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
}

func initAPI() {
	beego.Router("/api/register", &api_controller.RegisterController{}, "POST:Register")
	beego.Router("/api/login", &api_controller.RegisterController{}, "POST:Login")

	beego.Router("/api/users", &api_controller.UsersController{}, "get:GetUsers")
	beego.Router("/api/users/:user", &api_controller.UsersController{},"get:GetUser")
    beego.Router("/api/users/:user/all-repos", &api_controller.UsersController{},"get:GetUserAllRepos")
	beego.Router("/api/users/:user/repos", &api_controller.UsersController{}, "get:GetUserRepos")
	beego.Router("/api/users/:user/repos/add/:repo", &api_controller.UsersController{}, "get:AddUserRepo")
	beego.Router("/api/users/:user/repos/delete/:repo", &api_controller.UsersController{}, "get:DeleteUserRepo")

	beego.Router("/api/users/:user/starring-repos", &api_controller.UsersController{}, "get:GetUserStarringRepos")
	beego.Router("/api/users/:user/starring-repos/update", &api_controller.UsersController{}, "get:UpdateUserStarringRepos")

	beego.Router("/api/users/:user/hitter", &api_controller.UsersController{}, "get:GetUserHitter")
	beego.Router("/api/users/:user/hitter/update/:hitter", &api_controller.UsersController{}, "get:UpdateUserHitter")

	beego.Router("/api/users/:user/targets/:target", &api_controller.UsersController{}, "get:GetUserTarget")
	beego.Router("/api/users/:user/targets/:target/status", &api_controller.UsersController{}, "get:GetUserTargetStatus")
	beego.Router("/api/users/:user/targets/:target/pool", &api_controller.UsersController{}, "get:GetUserTargetPool")
	beego.Router("/api/users/:user/status", &api_controller.UsersController{}, "get:GetUserStatus")

	beego.Router("/api/users/:user/recommend", &api_controller.UsersController{}, "get:GetUserRecommend")

	beego.Router("/api/users/:user/owe", &api_controller.UsersController{}, "get:GetUserOwe")

	beego.Router("/api/global/starring-repos/update", &api_controller.GlobalController{}, "get:UpdateStarringRepos")
	beego.Router("/api/global/recommend", &api_controller.GlobalController{}, "get:GetRecommend")
	beego.Router("/api/global/owe", &api_controller.GlobalController{}, "get:GetOwe")
}

func initView(){
	beego.Router("/", &view_controller.ViewController{}, "GET:Index")
	beego.Router("/update", &view_controller.ViewController{}, "GET:Update")
	beego.Router("/login", &view_controller.ViewController{}, "GET:LoginPage")
	beego.Router("/login", &view_controller.ViewController{}, "POST:Login")
	beego.Router("/register", &view_controller.ViewController{}, "GET:RegisterPage")
	beego.Router("/register", &view_controller.ViewController{}, "POST:Register")
	beego.Router("/logout", &view_controller.ViewController{}, "GET:Logout")
	beego.Router("/about", &view_controller.ViewController{}, "GET:About")

	beego.Router("/user/setting", &view_controller.ViewController{}, "GET:SettingPage")
	beego.Router("/user/setting", &view_controller.ViewController{}, "POST:Setting")
	beego.Router("/repo/add", &view_controller.ViewController{}, "GET:AddRepoPage")
	beego.Router("/repo/add", &view_controller.ViewController{}, "Post:AddRepo")
	beego.Router("/repo/delete/:repo", &view_controller.ViewController{}, "GET:DeleteRepo")

	beego.Router("/users/:user", &view_controller.ViewController{}, "GET:UserPage")
	beego.Router("/users", &view_controller.ViewController{}, "GET:UsersPage")
	beego.Router("/count", &view_controller.ViewController{}, "GET:CountPage")

	beego.Router("/owe", &view_controller.ViewController{}, "GET:OwePage")
	beego.Router("/owes", &view_controller.ViewController{}, "GET:OwesPage")

	beego.Router("/repo", &view_controller.ViewController{}, "GET:RepoPage")

	beego.Router("/referrer", &view_controller.ViewController{}, "GET:ReferrerPage")
	beego.Router("/log", &view_controller.ViewController{}, "GET:LogPage")
}