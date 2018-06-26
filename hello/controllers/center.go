package controllers

import (
	"github.com/astaxie/beego"
	"hello/models"
	"fmt"
	"hello/extend"
	"time"
)

type CenterIndexController struct {
	beego.Controller
}
func (this *CenterIndexController) Get(){
	sessionUserNameDf	:=	this.GetSession("loginUser")
	if sessionUserNameDf == nil{
		this.Ctx.WriteString("没有登陆")
		return;
	}
	sessionUserName	:=	sessionUserNameDf.(string)
	//if this.Data["username"] ==nil{
	//	this.Redirect("/login",403)
	//	return
	//}
	getId,_ := this.GetInt64("id",2)
	//获取文章列表
	getOneArt	:=	models.GetAticleInfoById(getId)
	this.Data["art_info"]	=	getOneArt
	whereMap	:=	map[string]string{"title":"五一节","status":"1"}
	num,err,getArtList	:=	models.GetAticleList(whereMap,1,10)
	if err == nil{
		fmt.Println(err)
	}
	//根据用户名获取用户信息
	userInfo := models.QueryByName(sessionUserName)
	//获取用户的详细信息
	fmt.Println("------")
	obUserDetail :=&models.DetailClass{sessionUserName}
	userDetail := obUserDetail.GetUserDetailInfo()//用户详情信息
	obUserExtend := extend.ObUserInfo{}//实例化
	addUserDetailData := obUserExtend.AddUserDetail(&userInfo,&userDetail)
	this.Data["user_detail"]	=	addUserDetailData
	fmt.Println(addUserDetailData.Id)
	///获取类
	idArr :=extend.UserData(getArtList)
	userData:=models.GetUserDayByUserIdArr(idArr)
	reList:=extend.ArticleAddUserData(&getArtList,&userData)
	this.Data["num"]		=	num
	this.Data["art_list"]	=	reList
	this.TplName = "center.html"
}
/**
更新数据
 */
func (this *CenterIndexController) Post(){
	userDec	:=	this.GetString("dec")
	userPic := this.GetString("pic")
	sessionUserName	:=	this.GetSession("loginUser").(string)
	obUserDetail :=&models.DetailClass{sessionUserName}
	findHasData :=	obUserDetail.FindData(sessionUserName)
	var returnBool bool
	if !findHasData{
		fmt.Println("add")
		returnBool = obUserDetail.AddUserDetail(userPic,userDec)//add
	}else{
		fmt.Println("update")
		fmt.Println(obUserDetail.User_name)
		returnBool = obUserDetail.UpdateUserDetail(userPic,time.Unix(time.Now().Unix(),0))//edit
	}
	if returnBool {
		this.Redirect("/user/index",200)
	}
}