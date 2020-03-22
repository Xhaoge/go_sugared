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
