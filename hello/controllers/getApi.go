package controllers

import (
	"github.com/astaxie/beego"
	"hello/extend"
	"fmt"
)

type GetApiController struct {
	beego.Controller
}
/*
运行测试
 */
func (this *GetApiController) RunTest(){
	ruleId := this.GetStrings("rule_id")
	obGetApiExtend	:=	extend.GetApiExtend{}
	id	:=	this.Input().Get("id")
	info,urlList:=obGetApiExtend.ReturnUrl(ruleId,id)
	this.Ctx.WriteString("<br />")
	if len(urlList) <=0 {
		this.Data["goUrl"]	=	"rule_api/index"
		this.Data["goMsg"]	=	"编号不能为空"
		this.TplName		=	"public/error.html"
		return
	}
	var result string
	for _,v := range urlList{
		result =obGetApiExtend.GoTest(v,info,id)
		this.Ctx.WriteString("<hr />")
		this.Ctx.WriteString(result)
		}
}
func (this *GetApiController) Show(){
	fmt.Println(this.Input())
	this.Ctx.WriteString(this.Input().Encode())
}
func (this *GetApiController) Test(){
	id,err := this.GetInt("id")
	if err != nil{
		this.Data["goUrl"]	=	"rule_api/index"
		this.Data["goMsg"]	=	"编号不能为空"
		this.TplName		=	"public/error.html"
		return
	}
	obGetApiExtend	:=	extend.GetApiExtend{&id}
	data,result 	:=	obGetApiExtend.GetApiInfo()
	if result.Code!=200{
		this.Data["goUrl"]	=	"rule_api/index"
		this.Data["goMsg"]	=	result.Msg
		this.TplName		=	"public/error.html"
		return
	}
	//get rulelist
	obRuleList	:=	extend.RuleExtend{}
	_,rule_list	:=	obRuleList.GetAllRule()
	this.Data["title"]	=	"apitesttitle"
	this.Data["key"]	=	"apiEditKey"
	this.Data["info"]	=	data
	this.Data["rule_list"]	=	rule_list
	this.TplName	=	"get_api/info.html"
}

