package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_sugared/routers/api"
)

var (
	wxLoginUrl = "https://api.weixin.qq.com/sns/jscode2session"
	wxappid    = "wxd20112c596493f06"
	wxSecret   = "c156da5ff391cd48fb76c90f99150c24"
)

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 201,
		"msg":  "xxxxx",
	})
}

func UserLogin(c *gin.Context) {
	fmt.Println("UserLogin test")
	wxReq := &wxLoginReq{}
	if err := c.BindJSON(&wxReq); err != nil {
		api.ApiBaseResponse(c, 400, "request param error..")
	} else {

	}
	c.JSON(200, gin.H{
		"code": 201,
		"msg":  "xxxxx",
	},
	)
}

func UserUpdate(c *gin.Context) {
	fmt.Println("file test add")
	c.JSON(200, gin.H{
		"code": 201,
		"msg":  "xxxxx",
	},
	)
}
