package initial

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com/astaxie/beego/session"
)

func InitSql() {
	dbhost 		:= 	beego.AppConfig.String("mysqlurls")
	dbport 		:= 	beego.AppConfig.String("mysqlport")
	dbuser 		:= 	beego.AppConfig.String("mysqluser")
	dbpassword 	:= 	beego.AppConfig.String("mysqlpass")
	db 			:= 	beego.AppConfig.String("mysqldb")
	//注册mysql Driver
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//构造conn连接
	//用户名:密码@tcp(url地址)/数据库
	conn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + db + "?charset=utf8"
	//注册数据库连接
	orm.RegisterDataBase("default", "mysql", conn)
	fmt.Printf("数据库连接成功！%s\n", conn)
}
func InitSession(){
	sessionConfig := &session.ManagerConfig{
		CookieName:"gosessionid",
		EnableSetCookie: true,
		Gclifetime:3600,
		Maxlifetime: 3600,
		Secure: false,
		CookieLifeTime: 3600,
		ProviderConfig: "./tmp",
	}
	globalSessions, _ := session.NewManager("memory",sessionConfig)
	go globalSessions.GC()
}