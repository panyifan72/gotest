package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Test_api struct {
	Id 				int
	Api_name 		string
	Api_param		string
	Success_data	string
	Test_rule_id 	string
	Api_url			string
	Api_method		int
	Ctm				int64
}
type ApiRuleClass struct {

}
/**
获取列表数据
 */
func(this *ApiRuleClass) GetList(where map [string]string,page int,offset int)(int64,error,[]Test_api){
	var TestApi []Test_api
	o := orm.NewOrm()
	qs := o.QueryTable("test_api")
	cond := orm.NewCondition()
	if where["api_name"] != ""{
		cond = cond.And("api_name",where["api_name"])
	}
	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset = 9
	}
	start := (page - 1) * offset
	num, err1 := qs.Limit(offset, start).All(&TestApi)
	return num,err1,TestApi
}
/**
获取单条记录
 */
func (this *ApiRuleClass) FindById(id int) Test_api{
	var reTestApi Test_api
	o := orm.NewOrm()
	artT := o.QueryTable("test_api")
	err := artT.Filter("id",id).One(&reTestApi)
	if err != nil{
		//日志记录
	}
	return reTestApi
}
/**
查询是否重复
 */
func (this *ApiRuleClass) FindOne(api_name string,id int) Test_api{
	var reTestApi Test_api
	o := orm.NewOrm()
	artT := o.QueryTable("test_api")
	if id >0 {
		err := artT.Filter("api_name", api_name).Filter("id__gt",id).Filter("id__lt",id).One(&reTestApi)

		if err!=nil{
			//日志记录错误
		}
	}else{
		err := artT.Filter("api_name", api_name).One(&reTestApi)
		if err!=nil{
			//日志记录错误
		}
	}
	return reTestApi
}
/*
增加修改数据
 */
func (this *ApiRuleClass)Oper(data Test_api)error{
	o := orm.NewOrm()
	var reErr	error
	getTime := time.Now().Unix()
	data.Ctm	=	getTime
	if data.Id>0{ //更新
		_, err := o.QueryTable("test_api").Filter("id", data.Id).Update(orm.Params{
			"api_name": data.Api_name,
			"api_param": data.Api_param,
			"test_rule_id": data.Test_rule_id,
			"api_url": data.Api_url,
			"api_method": data.Api_method,
			"success_data"	:	data.Success_data,
			"ctm": getTime,
		})
		reErr	=	err
	}else {//新增
		_, err := o.Insert(&data)
		reErr	=	err
	}
	return reErr
}
func init(){
	orm.RegisterModel(new(Test_api))
}
