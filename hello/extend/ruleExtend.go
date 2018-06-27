package extend

import (
	"hello/models"
	"fmt"
)

type RuleExtend struct {

}
func (this *RuleExtend)FindInfo(id int)models.Test_rule{
	ruleInfo := models.Test_rule{}
	if id<=0{
		return ruleInfo
	}
	obRuleClass :=models.RuleClass{}
	ruleInfo	=	obRuleClass.GetRuleById(id)
	return ruleInfo
}
/**
获取列表信息
 */
func (this *RuleExtend)ListExtend(where map[string]string,page int,offset int)(count int64,arr []models.Test_rule){
	count = 0
	arr = []models.Test_rule{}
	if page <1 || offset<1{
		return count,arr
	}
	obTestRule	:=	models.RuleClass{}
	count,err,arr		:=	obTestRule.GetRuleList(where,page,offset)
	fmt.Println(err)
	return count,arr
}
/**
operationData
 */
func (this *RuleExtend)OperationExtend(ruleData models.Test_rule) ReturnErr{
	if len(ruleData.Rule) == 0 || len(ruleData.Rule_name) == 0{
		return ReturnErr{403,""}
	}
	obRuleClass :=models.RuleClass{}
	//find one first
	findOne:=obRuleClass.FindOne(ruleData.Rule)
	if findOne.Id>0{
		return ReturnErr{601,"has this data"}
	}
	err:=obRuleClass.EditRule(ruleData)
	if err!=nil {
		//error errlog
		return ReturnErr{0,"system err"}
	}
	return ReturnErr{200,"success"}
}

