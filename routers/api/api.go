package api

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

var MgoSession *mgo.Session

func InitDatabase() error {
	var err error
	MgoSession, err = mgo.Dial("")
	//defer mongo.Close()
	if err != nil {
		return err
	}
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
