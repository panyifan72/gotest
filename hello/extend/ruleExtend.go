package extend

import (
	"hello/models"
	"fmt"
	//"time"
	"time"
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
type ListTestRule struct {
	Test_rule_in models.Test_rule
	Time string
}
/**
获取列表信息
 */
func (this *RuleExtend)ListExtend(where map[string]string,page int,offset int)(count int64,arr []ListTestRule){
	count = 0
	arr = []ListTestRule{}
	if page <1 || offset<1{
		return count,arr
	}

	obTestRule	:=	models.RuleClass{}
	count,err,oldArr		:=	obTestRule.GetRuleList(where,page,offset)
	fmt.Println(err)
	if count>0{
		for _,val:=range oldArr{
			arr = append(arr,ListTestRule{
				Test_rule_in:models.Test_rule{Id:val.Id,Rule_name:val.Rule_name,Rule:val.Rule,Status:val.Status,Ctm:val.Ctm},
				Time:time.Unix(val.Ctm, 0).Format("2006-01-02 15:04:05"),
			})
		}
	}
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
	findOne:=obRuleClass.FindOne(ruleData.Rule,ruleData.Id)
	fmt.Println(findOne)
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

