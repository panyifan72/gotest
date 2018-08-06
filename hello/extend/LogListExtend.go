package extend

import (
	"hello/models"
)

type LogListExtend struct {

}
type GoTestApiLogN struct {
	models.GoTestApiLog
	Test_api_name 	string
}
/**
获取列表数据
 */
func (this *LogListExtend) GetList(page int,pageSize int,types int)(int64,[]GoTestApiLogN){
	//获取规则名称
	var cnums int64
	var getList []GoTestApiLogN
	obTestApiModel	:=	models.ApiRuleClass{}
	getApiRuleInfo	:=	obTestApiModel.FindById(types)
	if getApiRuleInfo.Id ==0{
		return cnums,getList
	}
	obLogModel		:=	models.TestApiLogModel{}
	AllNums,getListDe 	:= 	obLogModel.GetList(page,pageSize,types)
	if AllNums<=0{
		return cnums,getList
	}
	cnums	=	AllNums
	var cData	GoTestApiLogN
	for _,val:=range getListDe{
			cData	=	GoTestApiLogN{val,getApiRuleInfo.Api_name}
			getList	=	append(getList,cData)
	}
	return cnums,getList
}
