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
