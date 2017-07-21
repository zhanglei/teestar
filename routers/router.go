package routers

import (
	"github.com/astaxie/beego"
	"github.com/hsluoyz/gitstar/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
