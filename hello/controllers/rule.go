package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"hello/models"
	"time"
	"hello/extend"
)
//rule list
type RuleListController struct {
	beego.Controller
}
//rule detail
type RuleInfoController struct {
	beego.Controller
}
//add rule
type RuleAddController struct {
	beego.Controller
}
//edit rule
type RuleEditController struct {
	beego.Controller
}
type RuleOperationController struct {
	beego.Controller
}
//get rule list
func (this *RuleListController) Get(){
	this.Ctx.WriteString("xueru ok")
}
//get  rule_info data
func (this *RuleInfoController) Get(){
}
//add and edit rule_info data
func (this *RuleInfoController) Post(){
}
func (this *RuleAddController) Get(){
	this.Data["title"] 	= 	"add rule"
	this.Data["key"]	=	"add_rule_key"
	this.TplName		=	"rule/add.html"
}
func (this *RuleEditController) Get(){
}
// post data edit
func (this *RuleOperationController) Post(){
	statusData,_:=strconv.Atoi(this.GetString("status"))
	id,_:=strconv.Atoi(this.GetString("id"))
	var postData = models.Test_rule{
		Id:id,
		Rule_name:this.GetString("rule_name"),
		Rule:this.GetString("rule"),
		Ctm:time.Now().Unix(),
		Status:statusData,
	}
	obRuleExtend:=extend.RuleExtend{}
	returnArr := obRuleExtend.OperationExtend(postData)
	this.Data["msg"]	=	returnArr.Msg
	this.Data["code"]	=	returnArr.Code
	this.TplName		=	"public/success.html"
}
