package extend

import "gopkg.in/mgo.v2"

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

