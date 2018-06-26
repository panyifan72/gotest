package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type UserDetail struct {
	Id int64
	User_name	string
	Pic		string
	Ctm time.Time
	Dec string
}
type DetailClass struct {
	User_name string
}
func (this *UserDetail) TableName() string {
	return "user_detail"
}
func init(){
	orm.RegisterModel(new(UserDetail))
}
/**
添加用户信息
 */
func (this *DetailClass)AddUserDetail(pic string,dec string) bool{
	o := orm.NewOrm()
	var userDetails UserDetail
	userDetails.User_name	=	this.User_name
	userDetails.Pic	=	pic
	userDetails.Dec	=	dec
	userDetails.Ctm	=	time.Unix(time.Now().Unix(),0)
	_, err := o.Insert(&userDetails)
	returnBool:=true
	if err != nil{
		returnBool = false
	}
	return returnBool
}
/**
获取数据
 */
func (this *DetailClass)FindData(userName string)bool{
	var user UserDetail
	o := orm.NewOrm()
	qs := o.QueryTable("user_detail")
	err := qs.Filter("user_name", userName).One(&user)
	returnBool:= true
	if err != nil{
		returnBool = false
	}
	return returnBool
}
/**
更新数据
 */
func (this *DetailClass)UpdateUserDetail(pic string,ctm time.Time)bool{
	o := orm.NewOrm()
	var userDetails  UserDetail
	userDetails.Pic	=	pic
	userDetails.Ctm	=	ctm
	qs:=o.QueryTable("user_detail").Filter("user_name", this.User_name)
	_, err := qs.Update(orm.Params{"pic":pic,"ctm":ctm})
	returnBool:=true
	if err != nil{
			returnBool = false
		}
	return returnBool
}

/**
获取详情数据
 */
func (this *DetailClass) GetUserDetailInfo()UserDetail{
	var UserArr UserDetail
	o := orm.NewOrm()
	qs := o.QueryTable("user_detail")
	userId:=this.User_name
	qs.Filter("user_name",userId).One(&UserArr)
	return UserArr
}