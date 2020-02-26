package controller

import (
	"fmt"

	"api/model"

	"github.com/gin-gonic/gin"
)

func Test(context *gin.Context) {
	fmt.Println("file test")
	context.JSON(200, gin.H{
		"code": 201,
		"msg":  "xxxxx",
	},
	)
}

func PicAdd(context *gin.Context) {
	fmt.Println("file test add")
	var user model.PicBaseReq
	err := context.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("id: ", user.Id)
	fmt.Println("name: ", user.Username)

	context.JSON(200, gin.H{
		"code": 201,
		"msg":  "xxxxx",
		"name": user.Username,
	},
	)
}

func PicDelete(context *gin.Context) {
	fmt.Println("file test delete")
	var user model.PicBaseReq
	err := context.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("id: ", user.Id)
	fmt.Println("name: ", user.Username)

	context.JSON(200, gin.H{
		"code": 201,
		"msg":  "xxxxx",
		"name": user.Username,
	},
	)
}
