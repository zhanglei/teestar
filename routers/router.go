package routers

import (
	"github.com/astaxie/beego"
	"github.com/hsluoyz/gitstar/controllers"
)

func init() {
    beego.Router("/api/users/:user/all-repos", &controllers.MainController{}, "get:GetUserAllRepos")
}
