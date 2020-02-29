package controller

import (
	"api/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RoomAdd(context *gin.Context) {
	fmt.Println("file test")
	roomAddReq := &model.NewRoomAddReq{}
	if err := context.BindJSON(&roomAddReq); err != nil {
		data := &model.RoomBaseResponse{
			Code: 500,
			Msg:  "xingbuxingo "}
		context.JSON(500, data)
		// context.JSON(200, gin.H{
		// 	"code": 500,
		// 	"msg":  "roomadd err",
		// },
		// )
	} else {
		fmt.Println("roomAddReq title: ", roomAddReq.Title)
		fmt.Println("roomAddReq PicIdList: ", roomAddReq.PicIdList)
		fmt.Println("roomAddReq Address: ", roomAddReq.Address)
		fmt.Println("roomAddReq IsElevator: ", roomAddReq.IsElevator)
		fmt.Println("roomAddReq ContactPhone: ", roomAddReq.ContactPhone)
		fmt.Println("roomAddReq Supporting: ", roomAddReq.Supporting)
		fmt.Println("roomAddReq Price: ", roomAddReq.Price)
		data2 := &model.RoomBaseResponse{
			Code: 200,
			Msg:  "xingbuxingo "}
		context.JSON(200, data2)

		// return &model.RoomBaseResponse{
		// 	Code: 200,
		// 	Msg:  "xingbuxingo ",
		// }
	}
}

func RoomIndex(context *gin.Context) {
	fmt.Println("file test add")
	context.JSON(200, gin.H{
		"code": 201,
		"msg":  "xxxxx",
	},
	)
}

func RoomGet(context *gin.Context) {
	fmt.Println("file test")
	context.JSON(200, gin.H{
		"code": 201,
		"msg":  "xxxxx",
	},
	)
}

func RoomDelete(context *gin.Context) {
	fmt.Println("file test add")
	context.JSON(200, gin.H{
		"code": 201,
		"msg":  "xxxxx",
	},
	)
}

func RoomUpdate(context *gin.Context) {
	fmt.Println("file test")
	context.JSON(200, gin.H{
		"code": 201,
		"msg":  "xxxxx",
	},
	)
}
