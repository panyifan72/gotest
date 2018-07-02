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
