package pic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_sugared/config"
	"go_sugared/pkg/file"
	"go_sugared/pkg/util"
	"io/ioutil"
)

func Test(c *gin.Context) {
	fmt.Println("file test")
	c.JSON(200, gin.H{
		"code": 201,
		"msg":  "xxxxx",
	},
	)
}

func AddPic(c *gin.Context) {
	fmt.Println("file test add")
	fh, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
	}
	fPic, _ := fh.Open()
	defer fPic.Close()
	bytes, err := ioutil.ReadAll(fPic)
	if err != nil {
		fmt.Println(err)
	}
	fileName := util.GetRandomStr(4) + "_" + util.GetRandomStr(4) + ".jpg"
	picPath := config.ConfigGetStaticPath()
	fmt.Println("file path:", picPath+"/"+fileName)
	if !file.PathExist(picPath){
		file.MkDir(picPath)
	}
	err = ioutil.WriteFile(picPath+"/"+fileName, bytes, 0666)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "fail",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}

func DeletePic(c *gin.Context) {
	fmt.Println("file test delete")
	// var user model.PicBaseReq
	// err := context.ShouldBindJSON(&user)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("id: ", user.Id)
	// fmt.Println("name: ", user.Username)

	// context.JSON(200, gin.H{
	// 	"code": 201,
	// 	"msg":  "xxxxx",
	// 	"name": user.Username,
	// },
	// )
}
