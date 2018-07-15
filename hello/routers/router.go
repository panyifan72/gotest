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
	beego.Router("admin/rule_api/add",&controllers.ApiListController{},"get:Add")
	beego.Router("admin/rule_api/index",&controllers.ApiListController{})
	beego.Router("admin/rule_api/edit",&controllers.ApiListController{},"get:Edit")
	beego.Router("admin/rule_api/operation",&controllers.ApiListController{},"post:Operation")
	beego.Router("admin/get_api/index",&controllers.GetApiController{},"get:Test")
	beego.Router("admin/get_api/run",&controllers.GetApiController{},"get:RunTest")
	beego.Router("admin/get_api/testuri",&controllers.GetApiController{},"get:Show")
	//beego.Router("test",&controllers.TestController{})
}
