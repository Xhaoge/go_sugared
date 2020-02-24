package main

import (
	"fmt"
	// "go_hh_DoubleKill/api/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Engin
	//router := gin.Default()
	// 测试路由
	r := gin.Default()
	fmt.Println("test")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "congratulation to you",
		})
	})

	// router := gin.New()
	//     // 加载html文件，即template包下所有文件
	// router.LoadHTMLGlob("template/*")
	// router.GET("/hello", hello)
	// // 路由组
	// user := router.Group("/user"){
	// 	// 请求参数在请求路径上
	//     user.GET("/get/:id/:username",controller.QueryById)
	//     user.GET("/query",controller.QueryParam)
	//     user.POST("/insert",controller.InsertNewUser)
	//     user.GET("/form",controller.RenderForm)// 跳转html页面
	//     user.POST("/form/post",controller.PostForm)
	//     //可以自己添加其他，一个请求的路径对应一个函数

	//     // ...
	// }

	// file := router.Group("/file"){
	//     // 跳转上传文件页面
	//     file.GET("/view",controller.RenderView) // 跳转html页面
	//     // 根据表单上传
	//     file.POST("/insert",controller.FormUpload)
	//     file.POST("/multiUpload",controller.MultiUpload)
	//     // base64上传
	//     file.POST("/upload",controller.Base64Upload)
	// }

	// // 加载配置文件
	// var err error
	// fPath, _ := os.Getwd()
	// fPath = path.Join(fPath, "conf")
	// configPath := flag.String("c", fPath, "config file path")
	// flag.Parse()
	// err = system.LoadConfigInformation(*configPath)
	// fmt.Printf("%+v\n",common.ConfigInfo.Server)
	// if err != nil {
	//     return
	// }

	// 指定地址和端口号
	r.Run(":8090")
}
