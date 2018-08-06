package controllers
import (
	"github.com/astaxie/beego"
	"hello/extend"
)

type LogListController struct {
	beego.Controller
}
/**
显示数据
 */
func (this *LogListController) Show(){
	page,pageErr := this.GetInt("page",1)
	pageSize,pageSizeErr := this.GetInt("pageSize",10)
	types,typesErr := this.GetInt("types",1)
	if pageErr != nil || pageSizeErr != nil || typesErr != nil{//类型错误
		this.Data["goUrl"]	=	"rule_api/index"
		this.Data["goMsg"]	=	"数据有错"
		this.TplName		=	"public/error.html"
		return
	}
	obTestExtend := extend.LogListExtend{}
	counts,list	:=	obTestExtend.GetList(page,pageSize,types)
	this.Data["count"]	=	counts
	this.Data["list"]	=	list
	this.TplName		=	"log_list/list.html"
}

