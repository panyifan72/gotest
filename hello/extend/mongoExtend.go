package extend

import (
	"gopkg.in/mgo.v2"
	"fmt"
)

type MongoExtendClass struct {

}
var PublicSession *mgo.Session
func init(){
	url	:=	"127.0.0.1"
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	PublicSession	=	session
	//defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
}
/*
获取数据库
 */
func (this *MongoExtendClass) Data() *mgo.Database{
	database:="test"
	c := PublicSession.DB(database)
	return c
}

/*
插入数据
使用方法
type User struct
{
	Id_ bson.ObjectId `bson:"_id"`
	Name string `bson:"name"`
	Age int `bson:"age"`
	JoinedAt time.Time `bson:"joned_at"`
}
obMongo	:= extend.MongoExtendClass{}
obMongo.Insert("people",User{ Id_: bson.NewObjectId(), Name: "Jimmy Kuu", Age: 33 })

 */
func (this *MongoExtendClass) Insert(table string,data interface{})error{
	fmt.Println(data)
	err := this.Data().C(table).Insert(&data)
	//defer PublicSession.Close()//关闭mongo
	return err
}
/**
更新数据
update
使用方法
	upData := bson.M{"name":"superWang"}
	changeData := bson.M{"name":"superWang777","phone":"13855556666"}
	err:=obMongo.Update("people",upData,changeData,1)
----

updateAll
使用方法
upData := bson.M{"age":33}
	changeData := bson.M{"$set":bson.M{"name":"superWang777","phone":"13855556666"}}
	err:=obMongo.Update("people",upData,changeData,2)
 */
func (this *MongoExtendClass) Update(table string,where interface{},data interface{},status int)error{
	var err error
	if status == 1{
		deerr := 	this.Data().C(table).Update(where,&data)
		err = deerr
	}else{
		_,deErr :=	this.Data().C(table).UpdateAll(where,&data)
		err = deErr
	}
	//defer PublicSession.Close()//关闭mongo
	return err
}
/*
删除数据
使用方法
obMongo.Del("people",bson.M{"name":"superWang"})
 */
func (this *MongoExtendClass) Del(table string,data interface{}) error{
	var err error
	err = this.Data().C(table).Remove(&data)
	return err
}
/*
查询一条记录
 */
func (this *MongoExtendClass) FindOne(table string,id string){
	info,err := this.Data().C(table).FindId().One()
}
/*
查询多条记录
 */
func (this *MongoExtendClass) FindAll(){

}

