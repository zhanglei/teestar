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
	"github.com/hsluoyz/gitstar/controllers/api_controllers"
	"github.com/hsluoyz/gitstar/controllers/view_controllers"
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
	ns :=
		beego.NewNamespace("/api",
			beego.NSNamespace("/user",
				beego.NSInclude(
					&api_controllers.UserController{},
				),
			),
			beego.NSNamespace("/users",
				beego.NSInclude(
					&api_controllers.UsersController{},
				),
			),
			beego.NSNamespace("/global",
				beego.NSInclude(
					&api_controllers.GlobalController{},
				),
			),
		)
	beego.AddNamespace(ns)

	//beego.Router("/api/user/register", &api_controllers.UserController{}, "POST:Register")
	//beego.Router("/api/user/login", &api_controllers.UserController{}, "POST:Login")
	//
	//beego.Router("/api/users", &api_controllers.UsersController{}, "get:GetUsers")
	//beego.Router("/api/users/:user", &api_controllers.UsersController{},"get:GetUser")
	//
	//beego.Router("/api/users/:user/repos/all", &api_controllers.UsersController{},"get:GetUserAllRepos")
	//beego.Router("/api/users/:user/repos", &api_controllers.UsersController{}, "get:GetUserRepos")
	//beego.Router("/api/users/:user/repos/add/:repo", &api_controllers.UsersController{}, "get:AddUserRepo")
	//beego.Router("/api/users/:user/repos/delete/:repo", &api_controllers.UsersController{}, "get:DeleteUserRepo")
	//
	//beego.Router("/api/users/:user/starring-repos", &api_controllers.UsersController{}, "get:GetUserStarringRepos")
	//beego.Router("/api/users/:user/starring-repos/update", &api_controllers.UsersController{}, "get:UpdateUserStarringRepos")
	//
	//beego.Router("/api/users/:user/hitter", &api_controllers.UsersController{}, "get:GetUserHitter")
	//beego.Router("/api/users/:user/hitter/update/:hitter", &api_controllers.UsersController{}, "get:UpdateUserHitter")
	//
	//beego.Router("/api/users/:user/status/targets/:target", &api_controllers.UsersController{}, "get:GetUserTargetStatus")
	//beego.Router("/api/users/:user/status", &api_controllers.UsersController{}, "get:GetUserStatus")
	//beego.Router("/api/users/:user/status/recommend", &api_controllers.UsersController{}, "get:GetUserRecommend")
	//beego.Router("/api/users/:user/status/owe", &api_controllers.UsersController{}, "get:GetUserOwe")
	//
	//beego.Router("/api/global/starring-repos/update", &api_controllers.GlobalController{}, "get:UpdateStarringRepos")
	//beego.Router("/api/global/recommend", &api_controllers.GlobalController{}, "get:GetRecommend")
	//beego.Router("/api/global/owe", &api_controllers.GlobalController{}, "get:GetOwe")
}

func initView(){
	beego.Router("/", &view_controllers.ViewController{}, "GET:Index")
	beego.Router("/update", &view_controllers.ViewController{}, "GET:Update")
	beego.Router("/login", &view_controllers.ViewController{}, "GET:LoginPage")
	beego.Router("/login", &view_controllers.ViewController{}, "POST:Login")
	beego.Router("/register", &view_controllers.ViewController{}, "GET:RegisterPage")
	beego.Router("/register", &view_controllers.ViewController{}, "POST:Register")
	beego.Router("/logout", &view_controllers.ViewController{}, "GET:Logout")
	beego.Router("/about", &view_controllers.ViewController{}, "GET:About")

	beego.Router("/user/setting", &view_controllers.ViewController{}, "GET:SettingPage")
	beego.Router("/user/setting", &view_controllers.ViewController{}, "POST:Setting")
	beego.Router("/user/changepwd", &view_controllers.ViewController{}, "POST:ChangeUserPassword")
	beego.Router("/repo/add", &view_controllers.ViewController{}, "GET:AddRepoPage")
	beego.Router("/repo/add", &view_controllers.ViewController{}, "Post:AddRepo")
	beego.Router("/repo/delete/:repo", &view_controllers.ViewController{}, "GET:DeleteRepo")

	beego.Router("/users/:user", &view_controllers.ViewController{}, "GET:UserPage")
	beego.Router("/users", &view_controllers.ViewController{}, "GET:UsersPage")
	beego.Router("/count", &view_controllers.ViewController{}, "GET:CountPage")

	beego.Router("/owe", &view_controllers.ViewController{}, "GET:OwePage")
	beego.Router("/owes", &view_controllers.ViewController{}, "GET:OwesPage")

	beego.Router("/repo", &view_controllers.ViewController{}, "GET:RepoPage")

	beego.Router("/referrer", &view_controllers.ViewController{}, "GET:ReferrerPage")
	beego.Router("/log", &view_controllers.ViewController{}, "GET:LogPage")
}