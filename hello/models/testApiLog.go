package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"fmt"
)

type GoTestApiLog struct {
	Id 		int
	Loguri	string
	Log_data string
	Add_time string
}
type TestApiLogModel struct {
}
/*
添加数据
 */
func (this *TestApiLogModel)AddOne(logUri string,logData string)error{
	var err error
	o := orm.NewOrm()
	getTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("-----")
	fmt.Println(logData)
	fmt.Println(getTime)
	fmt.Println("-----")
	insertData := GoTestApiLog{0,logUri,logData,getTime}
	_, err = o.Insert(&insertData)
	return err
}
func init(){
	orm.RegisterModel(new(GoTestApiLog))
}
