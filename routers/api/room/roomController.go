package room

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_sugared/pkg/util"
	"go_sugared/routers/api"
	"gopkg.in/mgo.v2/bson"
)

// 新增房源信息
func AddRoom(c *gin.Context) {
	fmt.Println("add room")
	roomAddReq := &Room{}
	if err := c.BindJSON(&roomAddReq); err != nil {
		api.ApiBaseResponse(c, 400, "request param error....")
	} else {
		roomAddReq.PackageNumber = util.MakePackageNumber()
		roomAddReq.ToPrint()
		if err := roomAddReq.Insert(); err != nil {
			fmt.Println("insert err: ", err)
			api.ApiBaseResponse(c, 500, "insert room error....")
		} else {
			api.ApiBaseResponse(c, 200, "insert room sucess....")
		}
	}
}

// 删除房源信息
func DeleteRoom(c *gin.Context) {
	fmt.Println("delete room")
	roomDeleteReq := &SingleRoomReq{}
	if err := c.BindJSON(&roomDeleteReq); err != nil {
		api.ApiBaseResponse(c, 400, "request param error....")
	} else {
		if err := roomDeleteReq.Delete(); err != nil {
			fmt.Println("delete err: ", err)
			api.ApiBaseResponse(c, 500, "delete room error....")
		} else {
			api.ApiBaseResponse(c, 200, "delete room sucess....")
		}
	}
}

// 获取具体房源信息 通过packagenumber
func GetRoomDetail(c *gin.Context) {
	fmt.Println("get room detail")
	getreq := &SingleRoomReq{}
	if err := c.BindJSON(&getreq); err != nil {
		api.ApiBaseResponse(c, 500, "request param error!!!")
	} else {
		fmt.Println("request: ", getreq)
		req := MakeSelector(*getreq)
		var res []Room
		err := api.FindAllBySelector("hh", "room", req, bson.M{"_id": 0}, &res)
		if err != nil {
			fmt.Println("err: ", err)
		} else {
			api.ApiDataResponse(c, 200, "查找所有房源成功.", res)
		}
	}
}

// 获取全部房源信息
func GetRoomIndex(c *gin.Context) {
	fmt.Println("get all room list....")
	var res []Room
	err := api.FindAllBySelector("hh", "room", nil, bson.M{"_id": 0}, &res)
	if err != nil {
		fmt.Println("find all err: ", err)
		api.ApiBaseResponse(c, 500, "find all room error....")
	} else {
		api.ApiDataResponse(c, 200, "index sucess", res)
	}
}

// 编辑房源信息
func UpdateRoom(c *gin.Context) {
	c.JSON(200, "helle, world...")
}
