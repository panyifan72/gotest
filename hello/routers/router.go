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
	beego.Router("admin/rule/index",&controllers.RuleListController{})//ruleList data
	beego.Router("admin/rule/info",&controllers.RuleInfoController{})//rule_info data
	beego.Router("admin/rule/add",&controllers.RuleAddController{})//add
	beego.Router("admin/rule/edit",&controllers.RuleEditController{})//edit
	beego.Router("admin/rule/operation",&controllers.RuleOperationController{})//operation
}