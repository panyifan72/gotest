package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"strconv"
)

type GoTestApiLog struct {
	Id 		int
	Loguri	string
	Log_data string
	Add_time string
	Test_api_id	int
}
type TestApiLogModel struct {
}
/*
添加数据
 */
func (this *TestApiLogModel)AddOne(logUri string,logData string,test_api string)error{
	var err error
	o := orm.NewOrm()
	getTime := time.Now().Format("2006-01-02 15:04:05")
	test_api_int,_	:=	strconv.Atoi(test_api)
	insertData := GoTestApiLog{0,logUri,logData,getTime,test_api_int}
	_, err = o.Insert(&insertData)
	return err
}
/**
数据库获取数据
 */
func (this *TestApiLogModel)GetList(page int,offset int,types int)(int64,[]GoTestApiLog){
	var resultData	[]GoTestApiLog
	o := orm.NewOrm()
	qs := o.QueryTable("go_test_api_log")
	cond := orm.NewCondition()
	if types>0 {
		cond = cond.And("test_api_id",types)
	}
	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset = 9
	}
	start := (page - 1) * offset
	qs.Limit(offset, start).All(&resultData)
	coutData,_	:=	qs.Count()//统计数据
	return coutData,resultData
}

func init(){
	orm.RegisterModel(new(GoTestApiLog))
}
