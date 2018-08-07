package extend

import (
	"hello/models"
	"time"
	"strings"
	"strconv"
	"github.com/mikemintang/go-curl"
	"fmt"
)

type GetApiExtend struct {
	Id	*int
}
/**
返回数据类型
 */
type ApiInfoData struct {
	Test_api 		models.Test_api
	Ctm_display		string
	TestRuleList	[]models.Test_rule
}
/**
测试
 */
func(this *GetApiExtend) GoTest(v string,info models.Test_api,id string)string{
	req:=curl.NewRequest()
	var addResult string
	if info.Api_method == 1{
		result,_ := req.SetUrl(v).Get()
		if result.IsOk(){
			addResult	=	result.Body
			//记录测试过程日志
			obTestLogModel := models.TestApiLogModel{}
			obTestLogModel.AddOne(v,result.Body,id)//添加测试数据
		}else {
			fmt.Println(result.Raw)
		}
	}else if info.Api_method == 2 {
		result,_ := req.SetUrl(v).Post()
		if result.IsOk(){
			addResult	=	result.Body
		}else {
			fmt.Println(result.Raw)
		}
	}else{
		addResult	=	"解析失败"
	}
	return addResult
}
/**
 返回链接进行测试
 */
func (this *GetApiExtend) ReturnUrl(rule_id []string,id string)(models.Test_api,[]string){
	var rsstring []string
	var getApiInfo	models.Test_api
	int_id,_ := strconv.Atoi(id)
	if int_id<=0 {
		return getApiInfo,rsstring
	}
	obTestApi	:=	models.ApiRuleClass{}
	getApiInfo =	obTestApi.FindById(int_id)
	if getApiInfo.Id <=0 {
		return getApiInfo,rsstring
	}
	ruleList := this.GetRuleListByIdIn(&rule_id)
	apiParams := strings.Split(getApiInfo.Api_param," ")
	//获取动态token
	obToken		:=	models.TokenModel{}
	resultToken	:=	obToken.GetInfo()
	var pingjie string
	var newApiUrl string
	for _,vals := range ruleList{
		for pa_key,pa_val := range apiParams{//顺序返回
			pingjie += pa_val+"="+vals.Rule+"&"
			if len(pingjie) <=0 {
				continue
			}
			checkStrings := strings.Contains(getApiInfo.Api_url,"?")
			newApiUrl	=	getApiInfo.Api_url
			if checkStrings{
				newApiUrl	=	newApiUrl+"&"
			}else{
				newApiUrl	=	newApiUrl+"?"
			}
			if getApiInfo.Token_status >0 {
				newApiUrl	=	newApiUrl+resultToken.Token_name+"="+resultToken.Token_val+"&"
			}
			onlyString := newApiUrl+pa_val+"="+vals.Rule//单独
			rsstring = append(rsstring,onlyString)
			if pa_key>0{
				thisUrl:=strings.TrimRight(newApiUrl+pingjie,"&")//组合
				rsstring = append(rsstring,thisUrl)
			}
		}
		pingjie = ""
	}
	if getApiInfo.Success_data != ""{//正确数据添加
		apiSuccesData	:=	strings.Split(getApiInfo.Success_data," ")
		var newStr string
		for i:=0;i<len(apiSuccesData);i++ {
			newStr += apiParams[i]+"="+apiSuccesData[i]+"&"
		}
		newStr	=	strings.TrimRight(newStr,"&")
		checkOnly := strings.Contains(getApiInfo.Api_url,"?")
		newApiUrl	=	getApiInfo.Api_url
		if checkOnly{
			newApiUrl	=	newApiUrl+"&"
		}else{
			newApiUrl	=	newApiUrl+"?"
		}
		if getApiInfo.Token_status >0 {//添加自动参数token
			newApiUrl	=	newApiUrl+resultToken.Token_name+"="+resultToken.Token_val+"&"
		}
		successUrl		:=	newApiUrl+newStr
		rsstring = append(rsstring,successUrl)
	}
	return getApiInfo,rsstring
}
/*
abc
a   b   c
ab   ac  bc
abc
/*
获取相关列表
 */
 func paramsArr(str []string){
	strLen := len(str)
	for i:=0;i<strLen;i++{
		if i!=(strLen-1){
			//paramsArr()
		}
	}
 }
func (this *GetApiExtend) GetRuleListByIdIn(id *[]string)[]models.Test_rule{
	var result []models.Test_rule
	var idList []int
	for _,val := range *id{
		id,_ := strconv.Atoi(val)
		idList = append(idList,id)
	}
	obTestRule := models.RuleClass{}
	_,result	=	obTestRule.GetRuleByIdIn(&idList)
	return result
}
/**
获取记录信息
 */
func (this * GetApiExtend) GetApiInfo()(ApiInfoData,ReturnErr){
	var ReturnApiInfoData ApiInfoData
	obApiRule	:= 	models.ApiRuleClass{}
	getInfo 	:=	obApiRule.FindById(*this.Id)
	if getInfo.Id == 0{
		return ReturnApiInfoData,GoReturn(404,"")
	}
	ReturnApiInfoData.Test_api	=	getInfo
	ReturnApiInfoData.Ctm_display	=	time.Unix(getInfo.Ctm, 0).Format("2006-01-02 15:04:05")
	if len(getInfo.Test_rule_id) == 0{
		return ReturnApiInfoData,GoReturn(200,"")
	}
	idIntSlice		:=	getIntSlice(&getInfo.Test_rule_id)
	//获取规则列表数据
	obTestRule 		:= 	models.RuleClass{}
	nums,RuleList	:=	obTestRule.GetRuleByIdIn(&idIntSlice)
	if nums <=0{
		return ReturnApiInfoData,GoReturn(404,"")
	}
	ReturnApiInfoData.TestRuleList = RuleList
	return ReturnApiInfoData,GoReturn(200,"")
}
/**
修改id字符串为id数组
 */
func getIntSlice(id *string) []int{
	idStrings := strings.Split(*id,",")
	idArr	:=	make([]int,len(idStrings))
	for i:=0 ;i<len(idStrings);i++ {
		intId,_	:=	strconv.Atoi(idStrings[i])
		idArr 	= 	append(idArr,intId)
	}
	return idArr
}
