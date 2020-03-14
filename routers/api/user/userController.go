package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {

}

func UserLogin(c *gin.Context) {
	fmt.Println("file test")
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
