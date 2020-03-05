package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func UserLogin(context *gin.Context) {
	fmt.Println("file test")
	context.JSON(200,gin.H{
			"code":201,
			"msg":"xxxxx",
		},
	)
}


func UserUpdate(context *gin.Context) {
	fmt.Println("file test add")
	context.JSON(200,gin.H{
			"code":201,
			"msg":"xxxxx",
		},
	)
}