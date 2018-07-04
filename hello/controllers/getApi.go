package controllers

import (
	"github.com/astaxie/beego"
	"hello/extend"
	"fmt"
)

type GetApiController struct {
	beego.Controller
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
	fmt.Println(data)
	this.Ctx.WriteString("sss")
}
