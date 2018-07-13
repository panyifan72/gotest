package main

import (
	_ "github.com/go-sql-driver/mysql" // import your used driver
	_ "hello/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "hello/initial"
)
func main() {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	orm.Debug = true
	beego.Run()
}

