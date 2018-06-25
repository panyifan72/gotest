package extend

type Public struct {

}

/*
返回指定字段数组
 */
func (this * Public) GetFieldData(arr *[]map[string]string,field string) map[int]string{
	lenArr	:=	len(*arr)
	var retrunArr map[int]string
	 retrunArr =  make(map[int]string)
	if lenArr == 0{
		return retrunArr
	}
	keyStep:=0
	for _,val:= range *arr{
		if val[field] != ""{
			retrunArr[keyStep] = val[field]
			keyStep++
		}
	}
	return retrunArr
}
/**
修改数组顺序模式为指定字段
*/
func (this *Public) ChangeArr(arr *[]map[string]string,field string) {
	lenArr := len(*arr)
	if lenArr ==0{
		return
	}

}
/**
多维素组转字数组
*/
func (this *Public)getArrFieldStr(arr *[]map[string]string,fieldOne string,fieldTwo string){

}
/*
数组转字符串
*/
func (this *Public)getFieldStr(){

}

