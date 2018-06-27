package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"hello/models"
	"time"
	"hello/extend"
	"fmt"
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
	obRuleExtend	:=	extend.RuleExtend{}
	where:=map[string]string{"rule":"ssss"}
	count,listArr	:=	obRuleExtend.ListExtend(where,1,10)
	fmt.Println(listArr)
	this.Data["list"]	=	listArr
	this.Data["count"]	=	count
	this.TplName		=	"rule/list.html"
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))  时间字符串
	//fmt.Println(time.Unix(1389058332, 0).Format("2006-01-02 15:04:05"))时间戳转字符串
}
//get  rule_info data
func (this *RuleInfoController) Get(){
	id,_	:=	strconv.Atoi(this.GetString("id"))
	obRuleExtend	:=	extend.RuleExtend{}
	info	:=	obRuleExtend.FindInfo(id)
	if info.Id>0{
		if info.Id >0{
			tm:=time.Unix(info.Ctm,0).Format("2006-01-02 03:04:05 PM")
			this.Data["time"]	=	tm
		}
	}
	this.Data["info"] = 	info
	fmt.Println(info)
	this.TplName = "rule/info.html"
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
	id,_	:=	strconv.Atoi(this.GetString("id"))
	obRuleExtend	:=	extend.RuleExtend{}
	info	:=	obRuleExtend.FindInfo(id)
	if info.Id >0{
		tm:=time.Unix(info.Ctm,0).Format("2006-01-02 03:04:05 PM")
		this.Data["time"]	=	tm
		this.Data["info"] = 	info
		this.TplName = "rule/edit.html"
	}else{
		//数据不存在进行错误页显示
		this.Data["goUrl"]	=	"/admin/user/index"
		this.Data["goMsg"]	=	"信息不存在"
		this.TplName	=	"public/error.html"
	}
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

