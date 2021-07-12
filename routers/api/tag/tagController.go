package tag

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_sugared/pkg/util"
	"go_sugared/routers/api"
	"time"
)

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 201,
		"msg":  "xxxxx",
	})
}

func AddTag(c *gin.Context) {
	fmt.Println("add tag")
	tagAddReq := &Tag{}
	if err := c.BindJSON(&tagAddReq); err != nil {
		api.ApiBaseResponse(c, 400, "request param error...")
	} else {
		tagAddReq.Id = util.GetRandomStr(4)
		tagAddReq.CreatedBy = time.Now()
		tagAddReq.State = 1
		if err := tagAddReq.Insert(); err != nil {
			fmt.Println("insert err: ", err)
			api.ApiBaseResponse(c, 500, "insert tag error...")
		} else {
			res ,_:= findOneById(tagAddReq.Id)
			api.ApiDataResponse(c, 200, "insert room success...",res)
		}
	}
}

func DeleteTag(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 201,
		"msg":  "xxxxx",
	})
}

func UpdateTag(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 201,
		"msg":  "xxxxx",
	})
}
