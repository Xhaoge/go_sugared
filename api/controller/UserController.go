package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAddResponse(c *gin.Context) {
	reply1 := &model.Reply1{}
	if err := c.Bind(reply1); err != nil {
		fmt.Println(err.Error())
		return
	}
	reply1.UName = c.Request.Header["Authorization"][0]
	reply1.CreateTime = utils.GetTimeStr()
	reply1.ID = bson.NewObjectId()
	if reply1.Save(reply1.TID) {
		c.JSON(http.StatusOK, gin.H{
			"reply": *reply1,
		})
	} else {
		c.String(http.StatusOK, "内部错误")
	}
}
