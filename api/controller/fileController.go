package controller

import (
	"fmt"
	"io/ioutil"

	"api/model"

	"github.com/gin-gonic/gin"
)

const BASE_NAME = "E:/program/static/hh/photos"

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
	fh, err := context.FormFile("file")
	if err != nil {
		fmt.Println(err)
	}
	file, err2 := fh.Open()
	defer file.Close()
	if err2 != nil {
		fmt.Println(err2)
	}
	bytes, e := ioutil.ReadAll(file)
	e = ioutil.WriteFile(BASE_NAME+fh.Filename, bytes, 0666)
	fmt.Println(e)
	if e != nil {
		fmt.Println(e)
		context.JSON(500, gin.H{
			"code": 500,
			"msg":  "fail",
		})
	} else {
		context.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
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
