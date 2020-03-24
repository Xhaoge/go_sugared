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
		api.ApiResponse(c, 400, "request param error....")
	} else {
		roomAddReq.PackageNumber = util.MakePackageNumber()
		roomAddReq.ToPrint()
		if err := roomAddReq.Insert(); err != nil {
			fmt.Println("insert err: ", err)
			api.ApiResponse(c, 500, "insert room error....")
		} else {
			api.ApiResponse(c, 200, "insert room sucess....")
		}
	}
}

// 删除房源信息
func DeleteRoom(c *gin.Context) {
	fmt.Println("delete room")
	roomDeleteReq := &SingleRoomReq{}
	if err := c.BindJSON(&roomDeleteReq); err != nil {
		api.ApiResponse(c, 400, "request param error....")
	} else {
		if err := roomDeleteReq.Delete(); err != nil {
			fmt.Println("delete err: ", err)
			api.ApiResponse(c, 500, "delete room error....")
		} else {
			api.ApiResponse(c, 200, "delete room sucess....")
		}
	}
}

// 获取具体房源信息 通过packagenumber
func GetRoomDetail(c *gin.Context) {
	fmt.Println("get room detail")
	getreq := &SingleRoomReq{}
	if err := c.BindJSON(&getreq); err != nil {
		api.ApiResponse(c, 500, "request param error!!!")
	} else {
		fmt.Println("request: ", getreq)
		req := MakeSelector(*getreq)
		var res []Room
		err := FindAll("hh", "room", req, bson.M{"_id": 0}, &res)
		if err != nil {
			fmt.Println("err: ", err)
		} else {
			c.JSON(200, GetRoomList{
				Code: 200,
				Msg:  "查找所有房源成功",
				Data: res,
			})
		}

	}
	//} else {
	//	fmt.Println("request pkg: ", getreq)
	//	res, err := FindOneRoomByPackageNumber(getreq.PackageNumber)
	//	if err != nil {
	//		fmt.Println("find err: ", err)
	//		api.ApiResponse(c, 500, "get room detail error....")
	//	} else {
	//		fmt.Println("res1: ", res)
	//		api.ApiResponse(c, 200, "get room detail sucess....")
	//	}
	//}
}

// 获取具体房源信息 通过NearSubway

// 获取全部房源信息
func GetRoomIndex(c *gin.Context) {
	fmt.Println("get all room list....")
	//res, err := findAllRoomBySelector()
	var res []Room
	err := FindAll("hh", "room", nil, bson.M{"_id": 0}, &res)
	if err != nil {
		fmt.Println("find all err: ", err)
		api.ApiResponse(c, 500, "find all room error....")
	} else {
		c.JSON(200, GetRoomList{
			Code: 200,
			Msg:  "查找所有房源成功",
			Data: res,
		})
	}
}

// 编辑房源信息
func UpdateRoom(c *gin.Context) {
	c.JSON(200, "helle, world...")
}
