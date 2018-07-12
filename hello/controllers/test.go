package controllers

import (
	"github.com/astaxie/beego"
	"hello/extend"
	"gopkg.in/mgo.v2/bson"
	"time"
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
	obMongo	:= extend.MongoExtendClass{}
	//obMongo.Insert("people",User{ Id_: bson.NewObjectId(), Name: "Jimmy Kuu", Age: 33 })
	obMongo.Del("people",bson.M{"name":"superWang"})
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