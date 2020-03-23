package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

var (
	MgoSession *mgo.Session
)

func InitMongo() {
	var err error
	MgoSession, err = mgo.Dial("")
	//defer mongo.Close()
	if err != nil {
		fmt.Println("connect mongo error：", err)
	}

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

type RoomBaseResponse struct {
	BaseResponse
	Data map[string]interface{} `json:"data,omitempty"`
}

func ApiResponse(c *gin.Context, code int, msg string) {
	data := BaseResponse{
		Code: code,
		Msg:  msg}
	c.JSON(200, data)
}

// 数据库操作封装
func FindOneBySelector(db, collectin string, query, selector, res interface{}) (interface{}, error) {
	ms, c := connect(db, collectin)
	defer ms.Close()
	if err := c.Find(query).Select(selector).One(&res); err != nil {
		return nil, err
	}
	return res, nil
}

func FindAllBySelector(db, collectin string, query, selector, result interface{}) error {
	ms, c := connect(db, collectin)
	defer ms.Close()
	return c.Find(query).Select(selector).All(result)
}
