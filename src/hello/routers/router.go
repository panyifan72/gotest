package routers

import (
	"hello/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("login",&controllers.LoginController{})
	beego.Router("goLogin",&controllers.GoLoginController{})
	beego.Router("register",&controllers.RegisterController{})
	beego.Router("logout",&controllers.LoginOutController{})
    beego.Router("user/index",&controllers.CenterIndexController{})
}
