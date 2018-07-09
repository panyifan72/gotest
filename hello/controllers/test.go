package controllers

import (
	"github.com/astaxie/beego"
	"hello/extend"
)

type TestController struct {
	beego.Controller
}
type Person struct {
	Name  string
	Phone string
}

func (this * TestController) Get(){
	obMongo	:= extend.MongoExtendClass{}
	obMongo.Data().C("people").Insert(&Person{"superWang", "13478808311"})
	extend.PublicSession.Close()//关闭mongo
	//
	//var err error
	//err = c.Insert(&Person{"superWang", "13478808311"},
	//	&Person{"David", "15040268074"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//result := Person{}
	//err = c.Find(bson.M{"name": "superWang"}).One(&result)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("Name:", result.Name)
	//fmt.Println("Phone:", result.Phone)

	this.Ctx.WriteString("hell")
}