package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RoomAdd(context *gin.Context) {
	fmt.Println("file test")
	context.JSON(200,gin.H{
			"code":201,
			"msg":"roomadd",
		},
	)
}


func RoomIndex(context *gin.Context) {
	fmt.Println("file test add")
	context.JSON(200,gin.H{
			"code":201,
			"msg":"xxxxx",
		},
	)
}



func RoomGet(context *gin.Context) {
	fmt.Println("file test")
	context.JSON(200,gin.H{
			"code":201,
			"msg":"xxxxx",
		},
	)
}


func RoomDelete(context *gin.Context) {
	fmt.Println("file test add")
	context.JSON(200,gin.H{
			"code":201,
			"msg":"xxxxx",
		},
	)
}


func RoomUpdate(context *gin.Context) {
	fmt.Println("file test")
	context.JSON(200,gin.H{
			"code":201,
			"msg":"xxxxx",
		},
	)
}
