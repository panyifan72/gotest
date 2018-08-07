package controllers

import (
	"github.com/astaxie/beego"
	"hello/models"
)

type TokenController struct {
	beego.Controller
}
/**
获取数据
 */
func (this *TokenController) Get(){
	//获取动态token
	obToken		:=	models.TokenModel{}
	resultToken	:=	obToken.GetInfo()
	this.Data["info"]	=	resultToken
	this.TplName	=	"token/index.html"
}
/**
编辑数据
 */
func (this *TokenController) Post(){
	tokenName 	:= 	this.GetString("token_name")
	tokenVal 	:= 	this.GetString("token_val")
	obToken		:=	models.TokenModel{}
	err:=obToken.EditData(tokenName,tokenVal)
	if err !=nil{
		//数据不存在进行错误页显示
		this.Data["goUrl"]	=	"/admin/user/index"
		this.Data["goMsg"]	=	"信息不存在"
		this.TplName	=	"public/error.html"
	}else{
		this.Data["msg"]	=	"成功"
		this.Data["code"]	=	"200"
		this.TplName		=	"public/success.html"
	}
}
