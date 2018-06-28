package models

import (
	"github.com/astaxie/beego/orm"
	"time"
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
func (this *RuleClass)GetRuleById(id int) Test_rule{
	var test_rule Test_rule
	o := orm.NewOrm()
	artT := o.QueryTable("test_rule")
	err := artT.Filter("id", id).One(&test_rule)
	if err== nil{
		return test_rule
	}
	return test_rule
}
/**
获取全部
 */
func (this *RuleClass) GeALLRule()(int64,[]Test_rule){
	o := orm.NewOrm()
	qs := o.QueryTable("test_rule")
	var test_rule []Test_rule
	num, _ := qs.All(&test_rule)
	return num,test_rule
}
/*
get all rule
 */
func (this *RuleClass)GetRuleList(where map [string]string,page int,offset int)(int64,error,[]Test_rule){
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
func (this *RuleClass)FindOne(rule string,id int) Test_rule{
	var test_rule Test_rule
	o := orm.NewOrm()
	artT := o.QueryTable("test_rule")
	if id >0 {
		err := artT.Filter("rule", rule).Filter("id",id).One(&test_rule)
		if err!=nil{
			//日志记录错误
		}
	}else{
		err := artT.Filter("rule", rule).One(&test_rule)
		if err!=nil{
			//日志记录错误
		}
	}
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
	var reErr error
	if ruleData.Id>0{
		_, err := o.QueryTable("test_rule").Filter("id", ruleData.Id).Update(orm.Params{
			"rule_name": ruleData.Rule_name,
			"rule": ruleData.Rule,
			"ctm": ruleData.Ctm,
			"status": ruleData.Status,
		})
		reErr	=	err
	}else{
		_, err := o.Insert(&edit_rule)
		reErr	=	err
	}
	return reErr
}
func init(){
	orm.RegisterModel(new(Test_rule))
}
