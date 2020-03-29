package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_sugared/pkg/logging"
	"go_sugared/pkg/util"
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
		h := util.NewHttpSend(wxLoginUrl)
		data := map[string]string{
			"appid": wxappid,
			"secret":wxSecret,
			"js_code":wxReq.Code,
			"grant_type":"authorization_code",
		}
		h.Url = util.GetUrlBuild(h.Url,data)
		logging.Info("开始向小程序发起login请求，url： ",h.Url)
		fmt.Println("h.url: ",h.Url)
		resp, _ := h.Get()
		logging.Info("小程序login响应数据data： ",string(resp))
	}
	c.JSON(200, gin.H{
		"code": 201,
		"msg":  "xxxxx",
	},
	)
}

func UserUpdate(c *gin.Context) {
	fmt.Println("user test update")
	c.JSON(200, gin.H{
		"code": 201,
		"msg":  "xxxxx",
	},
	)
}
