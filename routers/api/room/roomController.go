package room

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_sugared/models"
)

// 获取全部房源信息
func GetRooms(c *gin.Context) {
}

// 获取具体房源信息
func GetRoomDetail(c *gin.Context) {
	fmt.Println("file test add")
	c.JSON(200, gin.H{
		"code": 201,
		"msg":  "xxxxx",
	},
	)
}

// 新增房源信息
func AddRooms(c *gin.Context) {
	fmt.Println("add room")
	//ms,m := models.ConnectMgo("hh","room")
	var roomAddReq models.OperateDB
	roomAddReq = &models.Room{}
	mgoobj := models.NewMgoObject(roomAddReq)

	var data models.RoomBaseResponse
	if err := c.BindJSON(&roomAddReq); err != nil {
		data.Code = 500
		data.Msg = "序列化失败，行不行哦！！！！"
	} else {
		if err := mgoobj.Insert(); err != nil {
			fmt.Println("insert err: ", err)
			data.Code = 400
			data.Msg = "insert room err!!!"
		} else {
			data.Code = 200
			data.Msg = "insert room success"
		}
		//err := models.Insert("hh", "room", roomAddReq)
		//if err != nil {
		//	fmt.Println("insert err: ", err)
		//}

		c.JSON(200, data)
	}
}

// 编辑房源信息
func UpdateRooms(c *gin.Context) {
}

// 删除房源信息
func DeleteRooms(c *gin.Context) {
}
