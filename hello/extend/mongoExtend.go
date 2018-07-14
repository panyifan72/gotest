package extend
//
//import "gopkg.in/mgo.v2"
//
//type MongoExtendClass struct {
//
//}
//var PublicSession *mgo.Session
//func init(){
//	url	:=	"127.0.0.1"
//	session, err := mgo.Dial(url)
//	if err != nil {
//		panic(err)
//	}
//	PublicSession	=	session
//	//defer session.Close()
//	// Optional. Switch the session to a monotonic behavior.
//	session.SetMode(mgo.Monotonic, true)
//}
///*
//获取数据库
// */
//func (this *MongoExtendClass) Data() *mgo.Database{
//	database:="test"
//	c := PublicSession.DB(database)
//	return c
//}
///*
//插入数据
// */
//func (this *MongoExtendClass) Insert(table string,data interface{})error{
//	err := this.Data().C(table).Insert(&data)
//	defer PublicSession.Close()//关闭mongo
//	return err
//}
///**
//更新数据
// */
//func (this *MongoExtendClass) Update(table string,data interface{})error{
//	var err error
//	return err
//}
//
