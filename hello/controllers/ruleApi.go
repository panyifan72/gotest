package controllers

import (
	"github.com/astaxie/beego"
	"hello/extend"
	"strconv"
	"fmt"
)

type ApiListController struct {
	beego.Controller
}
type ApiAddController struct {
	beego.Controller
}
type ApiEditController struct {
	beego.Controller
}
type ApiOperationController struct {
	beego.Controller
}
/*
列表详情
 */
func (this * ApiListController) Get(){
	this.Data["title"]	=	"apiListTitle"
	this.Data["key"]	=	"apiListKey"
	this.TplName		=	"rule_api/list.html"
}
/*
添加数据
 */
func (this *ApiAddController) Get() {
	obRuleExtend	:=	extend.RuleExtend{}
	_,ruleList 	:=	obRuleExtend.GetAllRule()
	this.Data["title"]	=	"apiAddTitle"
	this.Data["key"]	=	"apiAddKey"
	this.Data["rule_list"]	=	ruleList
	this.TplName		=	"rule_api/add.html"
}
/*
编辑数据
 */
func (this *ApiEditController) Get(){
	id,_:= strconv.Atoi(this.GetString("id"))
	if id<=0 {
		this.Data["goUrl"]	=	"rule_api/index"
		this.Data["goMsg"]	=	"数据不存在"
		this.TplName		=	"public/error.html"
		return
	}
	obRuleExtend	:=	extend.RuleExtend{}
	_,ruleList 	:=	obRuleExtend.GetAllRule()
	this.Data["title"]	=	"apiEditTitle"
	this.Data["key"]	=	"apiEditKey"
	this.Data["rule_list"]	=	ruleList
	this.TplName		=	"rule_api/edit.html"
}
/*
修改数据
 */
func (this *ApiOperationController) Post(){
	api_name	:=	this.GetString("api_name")
	api_param	:=	this.GetString("api_param")
	api_url		:=	this.GetString("api_url")
	api_method	:=	this.GetString("api_method")
	api_test_rule_id	:=	this.GetStrings("test_rule_id")
}

