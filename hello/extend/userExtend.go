package extend

import "hello/models"

type UserInfo struct {
	models.User
	UserPic string
	UserDec string
}
type ObUserInfo struct {

}
/**
组合详情内容
 */
func (ob *ObUserInfo) AddUserDetail(userInfo * models.User,userDetail * models.UserDetail) UserInfo{
	var UserInfoDara UserInfo
	userId:=userInfo.Id
	if userId == 0{
		return UserInfoDara
	}
	UserInfoDara = UserInfo{*userInfo,userDetail.Pic,userDetail.Dec}
	return UserInfoDara
}
