package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"fmt"
)

type Test_rule struct {
	Id			int
	Rule_name	string
	Rule		string
	Status 		int
	Ctm			int64
}

type RuleClass struct {
	Id 			int64
}
/*
get by id
 */
func GetRuleById(id int) Test_rule{
	var test_rule Test_rule
	o := orm.NewOrm()
	artT := o.QueryTable("test_rule")
	err := artT.Filter("id", id).One(&test_rule)
	if err== nil{
		return test_rule
	}
	return test_rule
}
/*
get all rule
 */
func GetRuleList(where map [string]string,page int,offset int)(int64,error,[]Test_rule){
	o := orm.NewOrm()
	qs := o.QueryTable("test_rule")
	cond := orm.NewCondition()
	if where["rule_name"] != ""{
		cond = cond.And("title",where["rule_name"])
	}
	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset = 9
	}
	start := (page - 1) * offset
	var test_rule []Test_rule
	num, err1 := qs.Limit(offset, start).All(&test_rule)
	return num,err1,test_rule
}
//find one by rule_name
func (this RuleClass)FindOne(rule string) Test_rule{
	var test_rule Test_rule
	o := orm.NewOrm()
	artT := o.QueryTable("test_rule")
	err := artT.Filter("rule", rule).One(&test_rule)
	if err== nil{
		return test_rule
	}
	fmt.Println(test_rule)
	return test_rule
}
//editData
func (this *RuleClass)EditRule( ruleData Test_rule) error{
	o := orm.NewOrm()
	var edit_rule Test_rule
	edit_rule.Id		=	ruleData.Id
	edit_rule.Rule_name	=	ruleData.Rule_name
	edit_rule.Rule	=	ruleData.Rule
	edit_rule.Ctm	=	time.Now().Unix()
	_, err := o.Insert(&edit_rule)
	return err
}
func init(){
	orm.RegisterModel(new(Test_rule))
}