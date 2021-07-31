package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go_sugared/config"
	_ "go_sugared/docs"
	"go_sugared/pkg/logging"
	"go_sugared/routers/api/pic"
	"go_sugared/routers/api/room"
	"go_sugared/routers/api/tag"
	"go_sugared/routers/api/user"
)

func InitRouter() *gin.Engine {
	logging.Info("start init routers ")
	gin.SetMode("debug")
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	return apiRoutes(engine)
}

func apiRoutes(engine *gin.Engine) *gin.Engine {
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiUser := engine.Group("hh/user")
	{
		//test
		apiUser.GET("/test", user.Test)
		//用户注册登录
		apiUser.POST("/login", user.UserLogin)
		//更新用户信息
		apiUser.POST("/update", user.UserUpdate)
	}

	apiRoom := engine.Group("hh/room")
	{
		//获取全部房源信息
		apiRoom.GET("/get", room.GetRoomIndex)
		//获取具体房源信息
		apiRoom.POST("/getDetail", room.GetRoomDetail)
		//新增房源信息
		apiRoom.POST("/add", room.AddRoom)
		//编辑房源信息
		apiRoom.POST("/update", room.UpdateRoom)
		//删除房源信息
		apiRoom.POST("/delete", room.DeleteRoom)
	}

	apiPic := engine.Group("hh/pic")
	{
		//test
		apiPic.GET("/test", pic.Test)
		//新增房源pic
		apiPic.POST("/add", pic.AddPic)
		//删除房源图片信息
		apiPic.POST("/delete", pic.DeletePic)
	}

	apiTag := engine.Group("hh/tag")
	{
		//test
		apiTag.GET("/test", tag.Test)
		//新增房源pic
		apiTag.POST("/add", tag.AddTag)
		//删除房源图片信息
		apiTag.POST("/delete", tag.DeleteTag)
		//编辑房源信息
		apiTag.POST("/update", tag.UpdateTag)
	}

	engine.GET("/test", func(c *gin.Context) {
		CONF := config.ConfigGetServicePort()
		fmt.Println("CONF port :", CONF)
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	return engine
}
