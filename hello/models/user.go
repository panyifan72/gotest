package models
import (
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/astaxie/beego/orm"
)
type User struct {
	Id          int
	Phone       string
	Password    string
	Status      int
	Created     int64
	Changed     int64
}

func (this *User) TableName() string {
	return "user"
}
func Test()int{
	return 1
}
/**
查询用户信息
 */
func GetUserDayByUserIdArr(idArr []int) []User{
	var UserArr []User
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	qs.Filter("id__in",idArr).All(&UserArr)
	return UserArr
}
//根据用户名称查询用户
func QueryByName(name string) (User) {
	var user User
	o := orm.NewOrm()
	qs := o.QueryTable("user")

	err := qs.Filter("phone", name).One(&user)
	if err == nil {
		return user
	}
	return user
}
func RegisterUser(username string,password string) (int64,error){
	o := orm.NewOrm()
	var user User
	user.Phone	=	username
	user.Password	=	password
	id, err := o.Insert(&user)
	return id,err
}
func init(){
	orm.RegisterModel(new(User))
}