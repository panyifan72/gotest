package models

import (
	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id 			int64 `orm:"auto"`
	Title 		string
	Keywords 	string
	Author 		string
	Content 	string
	Status 		int8
	Ctm			string
	User_id		int
}
/**
	获取单条记录
 */
func GetAticleInfoById(id int64) (Article){
	var articles Article
	o := orm.NewOrm()
	artT := o.QueryTable("article")
	err := artT.Filter("id", id).One(&articles)
	if err== nil{
		return articles
	}
	return articles
}

func GetAticleList(where map[string]string,page int,offset int) (num int64, err error, art []Article){
	o := orm.NewOrm()
	qs := o.QueryTable("article")
	cond := orm.NewCondition()
	if where["title"] != ""{
		cond = cond.And("title",where["title"])
	}
	if where["status"] !=""{
		cond = cond.And("status",where["status"])
	}
	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset = 9
	}
	start := (page - 1) * offset
	var articles []Article
	num, err1 := qs.Limit(offset, start).All(&articles)
	return num,err1,articles
}
type ArticleClass struct {
	Val int
}
func (x *ArticleClass) Test(a int) int {
	return x.Val+a+5
}
func init(){
	orm.RegisterModel(new(Article))
}