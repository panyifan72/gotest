package controllers
import (
	"github.com/astaxie/beego"
	"hello/models"
)
type LoginController struct {
	beego.Controller
}
func (this *LoginController)Get(){
	this.Data["title"]	=	"login"
	this.Data["key"]	=	"loginKey"
	//a :=[]map[string]string {0:{"a":"sdfsdf","b":"wqewqeqe","c":"wertre","d":"werewrw"},1:{"a":"qqqqqq","b":"wwwwww","c":"fffff","d":"sdfsdf"},2:{"a":"qqqqqq","b":"wwwwww","c":"fffff","d":"fgdfgdfg"}}
	//x:=extend.Public{}
	//returnData := x.GetFieldData(&a,"d")
	//fmt.Println(returnData)

	this.TplName	=	"login.html";
}
type GoLoginController struct {
	beego.Controller
}
type postData struct {
	username string
	password string
}
func (this *GoLoginController)Post(){
	postData := postData{this.GetString("username"),this.GetString("password")}
	if len(postData.username)<=0 && len(postData.password)<=0{
		this.Ctx.WriteString("错误")
		return
	}else{
		this.Ctx.WriteString("正确")
	}
	//userId,err := models.AddData(postData.username,postData.password)
	//if(user){
	//
	//}
}
type LoginOutController struct {
	beego.Controller
}
func (this *LoginOutController)Get(){
	this.Data["loginOut"]	=	"loginout"
	this.Data["key"]		=	"loginOutKey"
	this.TplName			=	"logout.html"
	this.DelSession("loginUser")
	this.DelSession("loginId")
	this.DestroySession()
}
type RegisterController struct {
	beego.Controller
}
func (this *RegisterController) Get(){
	this.Data["loginOut"]	=	"register"
	this.Data["key"]		=	"registerKey"
	this.TplName			=	"register.html"
}
func (this *RegisterController)Post(){
	////获取session用户是否登陆
	//if this.GetSession("loginUser") != nil {
	//	this.Ctx.WriteString("用户 已登陆")
	//	return
	//}
	postData := postData{this.GetString("username"),this.GetString("password")}
	//进行session记录
	this.SetSession("loginUser",postData.username)
	userM :=models.QueryByName(postData.username)
	if len(userM.Phone)>0{
		this.Ctx.WriteString("是滴是滴")
		return
	}
	if len(postData.username)<=0 && len(postData.password)<=0{
		this.Ctx.WriteString("错误")
		return
	}else{
		this.Ctx.WriteString("正确")
	}
	userId,err:=models.RegisterUser(postData.username,postData.password)
	if err !=nil{
		this.Ctx.WriteString("有错误的")
	}else{
		//注册成功走登陆流程
		if userId < 0{
			this.Ctx.WriteString("注册失败但无错误")
		}else{
			//进行session记录
			this.SetSession("loginUser",postData.username)
			this.SetSession("loginId",userId)
		}
	}
}