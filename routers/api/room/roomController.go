package room

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_sugared/pkg/util"
	"go_sugared/routers/api"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// 新增房源信息
func AddRoom(c *gin.Context) {
	fmt.Println("add room")
	roomAddReq := &Room{}
	if err := c.BindJSON(&roomAddReq); err != nil {
		api.ApiBaseResponse(c, 400, "request param error...")
	} else {
		roomAddReq.PackageNumber = util.MakePackageNumber()
		roomAddReq.ToPrint()
		roomAddReq.UpdateTime = time.Now()
		roomAddReq.ReleaseTime = time.Now()
		if err := roomAddReq.Insert(); err != nil {
			fmt.Println("insert err: ", err)
			api.ApiBaseResponse(c, 500, "insert room error...")
		} else {
			res, _ := FindOneRoomByPackageNumber(roomAddReq.PackageNumber)
			api.ApiDataResponse(c, 200, "insert room seccess...", res)
		}
	}
}

// 删除房源信息
func DeleteRoom(c *gin.Context) {
	fmt.Println("delete room")
	roomDeleteReq := &SingleRoomReq{}
	if err := c.BindJSON(&roomDeleteReq); err != nil {
		api.ApiBaseResponse(c, 400, "request param error...")
	} else {
		if err := roomDeleteReq.Delete(); err != nil {
			fmt.Println("delete err: ", err)
			api.ApiBaseResponse(c, 500, "delete room error...")
		} else {
			api.ApiBaseResponse(c, 200, "delete room seccess...")
		}
	}
}

// 编辑房源信息
func UpdateRoom(c *gin.Context) {
	roomUpdateReq := &Room{}
	if err := c.BindJSON(&roomUpdateReq); err != nil {
		api.ApiBaseResponse(c, 400, "request param error...")
	} else {
		update := MakeRoomUpdate(*roomUpdateReq)
		res, err := UpdateRoomBySelector("hh", "room", update, roomUpdateReq.PackageNumber)
		if err != nil {
			fmt.Println(err)
			api.ApiBaseResponse(c, 500, "update error...")
		} else {
			api.ApiDataResponse(c, 200, "update room seccess...", res)
		}
	}
}

// 获取具体房源信息 通过筛选条件
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
		if err != nil || len(res)==0{
			api.ApiDataResponse(c, 404, "find one room error or not find...", nil)
		} else {
			api.ApiDataResponse(c, 200, "find one room seccess...", res)
		}
	}
}

// 获取全部房源信息
func GetRoomIndex(c *gin.Context) {
	fmt.Println("get all room list....")
	var res []Room
	err := api.FindAllBySelector("hh", "room", nil, bson.M{"_id": 0}, &res)
	if err != nil {
		api.ApiBaseResponse(c, 500, "find all room error...")
	} else {
		api.ApiDataResponse(c, 200, "index seccess...", res)
	}
}
