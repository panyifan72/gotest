package models

import (
	"github.com/astaxie/beego/orm"
)

type Token struct {
	Id		int
	Token_name	string
	Token_val	string
}
type TokenModel struct {

}
/**
编辑数据
 */
func (this *TokenModel) EditData(token_name string,token_val string)error{
	o := orm.NewOrm()
	_, err := o.QueryTable("token").Update(orm.Params{
		"token_name": token_name,
		"token_val": token_val,
	})
	return err
}
/**
获取当前token
 */
func (this * TokenModel) GetInfo()Token{
	var resultToken Token
	o := orm.NewOrm()
	artT := o.QueryTable("token")
	artT.One(&resultToken)
	return resultToken
}
func init(){
	orm.RegisterModel(new(Token))
}
