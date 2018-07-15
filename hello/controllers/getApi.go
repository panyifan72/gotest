package controllers

import (
	"github.com/astaxie/beego"
	"hello/extend"
	"fmt"
	"github.com/mikemintang/go-curl"
)

type GetApiController struct {
	beego.Controller
}
func (this *GetApiController) RunTest(){
	ruleId := this.GetStrings("rule_id")
	obGetApiExtend	:=	extend.GetApiExtend{}
	urlList:=obGetApiExtend.ReturnUrl(ruleId,this.Input().Get("id"))
	if len(urlList) <=0 {
		this.Data["goUrl"]	=	"rule_api/index"
		this.Data["goMsg"]	=	"编号不能为空"
		this.TplName		=	"public/error.html"
		return
	}
	for _,v := range urlList{
		this
		req:=curl.NewRequest()
		result,_ := req.SetUrl(v).Post()
		this.Ctx.WriteString(result.Body)
		this.Ctx.WriteString("<br />")
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

