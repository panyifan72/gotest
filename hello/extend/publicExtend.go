package extend
type PublicExtend struct {

}
var ReturnData =map[int]string{
	200:"success",
	403:"aaaa",
	0:"system error",
}
type ReturnErr struct {
	Code int
	Msg string
}
func TimeChange(){

}
/*
go public errReturn
 */
func GoReturn(code int,msg string) ReturnErr{
	if len(msg) == 0 {
		msg = ReturnData[code]
	}
	return ReturnErr{code,msg}
}
