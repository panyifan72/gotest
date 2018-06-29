package extend

import (
	"strings"
	"hello/models"
	"fmt"
	"strconv"
)

type RuleApiClass struct {
}
func (this *RuleApiClass) GetList(){

}
func (this *RuleApiClass) GetInfo(){

}
type OperData struct {
	Id               string
	Api_name         string
	Api_param        string
	Api_url          string
	Api_method       string
	Api_test_rule_id []string
}
/**
获取详情信息
 */
func (this *RuleApiClass) GetInfoById(id int) OperData{
	var newDataInfo OperData
	if id <=0{
		return newDataInfo
	}
	obApiRuleModel	:=	models.ApiRuleClass{}
	dataInfo		:=	obApiRuleModel.FindById(id)
	if dataInfo.Id <=0{
		return newDataInfo
	}
	newDataInfo.Id			=	dataInfo.Id
	newDataInfo.Api_name	=	dataInfo.Api_name
	newDataInfo.Api_param	=	dataInfo.Api_param
	newDataInfo.Api_url		=	dataInfo.Api_url
	newDataInfo.Api_method	=	dataInfo.Api_method
	newDataInfo.Api_test_rule_id	=	dataInfo.test_rule_id//格式为1,2,3,4,5需要转换[1,2,3,4,5]

	return newDataInfo
}
func (this *RuleApiClass) Operation(data OperData) ReturnErr{
	//id处理
	apiId,_	:=	strconv.Atoi(data.Id)
	//检测是否存在记录
	obApiRuleModel	:=	models.ApiRuleClass{}
	checkResult		:=	obApiRuleModel.FindOne(data.Api_name,apiId)
	if checkResult.Id >0 { //说明记录存在，则返回失败
		return GoReturn(207,"")
	}
	var idStr string
	if len(data.Api_test_rule_id) > 0 {
		idStr = strings.Join(data.Api_test_rule_id,",")
	}
	var test_rule models.Test_api
	test_rule.Id			=	apiId
	test_rule.Api_name		=	data.Api_name
	test_rule.Api_param		=	data.Api_param
	test_rule.Api_url		=	data.Api_url
	test_rule.Api_method	=	data.Api_method
	test_rule.Test_rule_id	=	idStr
	fmt.Println(test_rule)

	err := obApiRuleModel.Oper(test_rule)
	if err != nil{
		return GoReturn(0,"")
	}
	return GoReturn(200,"")
}
