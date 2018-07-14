package extend

import (
	"hello/models"
	"time"
	"strings"
	"strconv"
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
func (this *GetApiExtend) ReturnUrl(rule_id []string,id string)[]string{
	var rsstring []string
	int_id,_ := strconv.Atoi(id)
	if int_id<=0 {
		return rsstring
	}
	obTestApi	:=	models.ApiRuleClass{}
	getApiInfo :=	obTestApi.FindById(int_id)
	if getApiInfo.Id <=0 {
		return rsstring
	}
	ruleList := this.GetRuleListByIdIn(&rule_id)
	apiParams := strings.Split(getApiInfo.Api_param," ")
	var pingjie string
	for _,vals := range ruleList{
		for _,pa_val := range apiParams{
			pingjie += pa_val+"="+vals.Rule+"&"
			if len(pingjie) <=0 {
				continue
			}
		}
		thisUrl:=strings.TrimRight(getApiInfo.Api_url+"?"+pingjie,"&")
		pingjie = ""
		rsstring = append(rsstring,thisUrl)
	}
	return rsstring
}
/*
获取相关列表
 */
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
