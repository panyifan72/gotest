package extend

import (
	"hello/models"
)

type RuleExtend struct {

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

