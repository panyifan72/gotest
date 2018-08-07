package extend

import (
	"strings"
	"hello/models"
	"fmt"
	"strconv"
	"time"
)

type RuleApiClass struct {
}
func (this *RuleApiClass) GetInfo(){

}
type OperData struct {
	Id               string
	Api_name         string
	Api_param        string
	Api_url          string
	Success_data	 string
	Api_method       int
	Api_test_rule_id []string
	Api_param_arr	[]string
	Token_status	int
	Api_test_rule_id_int []int
}
type TestApiList struct {
	Test_api models.Test_api
	Api_test_rule_str string
	Show_time	string
	Token_show string
}
/**
获取列表
 */
func (this *RuleApiClass) GetList(where map[string]string,page int,offset int)(int64,error,[]TestApiList){
	var returnTestApiList []TestApiList
	obApiTestClass:= models.ApiRuleClass{}
	count,err,apiTestList	:=	obApiTestClass.GetList(where,page,offset)
	if err != nil{
		return count,err,returnTestApiList
	}
	//获取全部规则信息
	obRuleClass := models.RuleClass{}
	_,RuleList 	:=	obRuleClass.GeALLRule()
	RuleNewList :=	ruleChange(&RuleList)
	splitStr	:=	""
	if count >0{
		for _,val:= range apiTestList{
			var mapString []string
			if len(val.Test_rule_id)>0 {//存在数据
				newString:=strings.Split(val.Test_rule_id,",")
				for _,stringVal := range newString{
					stringToInt,_:=strconv.Atoi(stringVal)
					if  stringToInt>0 {
						if _,ok := RuleNewList[stringToInt];ok{//判断key是否在map中
							mapString = append(mapString,RuleNewList[stringToInt].Rule_name)
						}
					}
				}
				if len(mapString)>0 {
					splitStr = strings.Join(mapString,",")
				}
			}else{
				splitStr	=	"-"
			}
			var token_shows	string
			if val.Token_status == 0{
				token_shows		=	"关闭"
			}else{
				token_shows		=	"开启"
			}
			returnTestApiList	=	append(returnTestApiList,TestApiList{
				Api_test_rule_str:splitStr,
				Test_api:val,
				Show_time:time.Unix(val.Ctm, 0).Format("2006-01-02 15:04:05"),
				Token_show:token_shows,
			})
		}
	}
	return count,err,returnTestApiList
}
/**
修改map类型
 */
func ruleChange(ruleList *[]models.Test_rule) map[int]models.Test_rule{
	var ruleChangeList map[int]models.Test_rule
	ruleChangeList	=	make(map[int]models.Test_rule)
	if len(*ruleList)  == 0{
		return ruleChangeList
	}
	for _,val:=range *ruleList{
		ruleChangeList[val.Id]	=	val
	}
	return ruleChangeList
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
	newDataInfo.Id			=	strconv.Itoa(dataInfo.Id)
	newDataInfo.Api_name	=	dataInfo.Api_name
	newDataInfo.Api_param	=	dataInfo.Api_param
	newDataInfo.Api_url		=	dataInfo.Api_url
	newDataInfo.Api_method	=	dataInfo.Api_method
	newDataInfo.Success_data	=	dataInfo.Success_data
	newDataInfo.Token_status	=	dataInfo.Token_status
	newDataInfo.Api_test_rule_id	=	strings.Split(dataInfo.Test_rule_id,",")//格式为1,2,3,4,5需要转换[1,2,3,4,5]
	if len(newDataInfo.Api_test_rule_id)>0 {
		for i:=0;i<len(newDataInfo.Api_test_rule_id);i++{
			strInt,_:=strconv.Atoi(newDataInfo.Api_test_rule_id[i])
			newDataInfo.Api_test_rule_id_int = append(newDataInfo.Api_test_rule_id_int,strInt)
		}
	}
	newDataInfo.Api_param_arr		=	strings.Split(dataInfo.Api_param,",")
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
	test_rule.Success_data	=	data.Success_data
	test_rule.Token_status	=	data.Token_status
	test_rule.Test_rule_id	=	idStr
	fmt.Println(test_rule)
	err := obApiRuleModel.Oper(test_rule)
	if err != nil{
		return GoReturn(0,"")
	}
	return GoReturn(200,"")
}
