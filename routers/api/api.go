package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_sugared/config"
	"gopkg.in/mgo.v2"
	"os"
)

var MgoSession *mgo.Session

func InitDatabase() error {
	var err error
	cfg := config.ConfigGetMongo()
	MgoSession, err = mgo.Dial(cfg.Host + ":" + cfg.Port)
	//defer mongo.Close()
	if err != nil {
		return err
	}
	MgoSession.SetMode(mgo.Eventual, true)
	myDB := MgoSession.DB("admin") //这里的关键是连接mongodb后，选择admin数据库，然后登录，确保账号密码无误之后，该连接就一直能用了
	//出现server returned error on SASL authentication step: Authentication failed. 这个错也是因为没有在admin数据库下登录
	err = myDB.Login(cfg.User, cfg.Pwd)
	if err != nil {
		fmt.Println("Login-error:", err)
		os.Exit(0)
	}
	//myDB = session.DB(mDBName) //如果要在这里就选择数据库，这个myDB可以定义为全局变量
	MgoSession.SetPoolLimit(10)
	return nil
}

func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	ms := MgoSession.Copy()
	c := ms.DB(db).C(collection)
	ms.SetMode(mgo.Monotonic, true)
	return ms, c
}

type BaseResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type DataResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ApiBaseResponse(c *gin.Context, code int, msg string) {
	data := BaseResponse{
		Code: code,
		Msg:  msg}
	c.JSON(200, data)
}

func ApiDataResponse(c *gin.Context, code int, msg string, data interface{}) {
	res := DataResponse{
		Code: code,
		Msg:  msg,
		Data: data}
	c.JSON(200, res)

}

// 数据库操作封装
func FindOneBySelector(db, collectin string, query, selector, res interface{}) error {
	_, c := connect(db, collectin)
	//defer ms.Close()
	return c.Find(query).Select(selector).One(&res)
}

func FindAllBySelector(db, collectin string, query, selector, result interface{}) error {
	_, c := connect(db, collectin)
	//defer ms.Close()
	return c.Find(query).Select(selector).All(result)
}

func UpdateBySelector(db, collectin string, query, update interface{}) error {
	_, c := connect(db, collectin)
	//defer ms.Close()
	return c.Update(query, update)
}
