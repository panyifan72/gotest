package controllers

import (
	"github.com/astaxie/beego"
<<<<<<< HEAD
=======
	"hello/extend"
	"gopkg.in/mgo.v2/bson"
	"time"
	"fmt"
>>>>>>> e41798e1bf8c4352e1806e0dc657acb23a1ed485
)

type TestController struct {
	beego.Controller
}
type User struct
{
	Id_ bson.ObjectId `bson:"_id"`
	Name string `bson:"name"`
	Age int `bson:"age"`
	JoinedAt time.Time `bson:"joned_at"`
}

func (this * TestController) Get(){
<<<<<<< HEAD
	//obMongo	:= extend.MongoExtendClass{}
	//obMongo.Insert("people",&Person{"superWang", "13478808311"})
	////obMongo.Data().C("people").Insert(&Person{"superWang", "13478808311"})
	//extend.PublicSession.Close()//关闭mongo
	//
	//var err error
	//err = c.Insert(&Person{"superWang", "13478808311"},
	//	&Person{"David", "15040268074"})
	//if err != nil {
	//	log.Fatal(err)
	//}
=======
	obMongo	:= extend.MongoExtendClass{}
	//obMongo.Insert("people",User{ Id_: bson.NewObjectId(), Name: "Jimmy Kuu", Age: 33 })
	//obMongo.Del("people",bson.M{"name":"superWang"})
	findWhere	:=	bson.M{"ip":1901531091}
	d := obMongo.FindAll("tender_custody_log",findWhere)
	fmt.Println(d)
>>>>>>> e41798e1bf8c4352e1806e0dc657acb23a1ed485
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