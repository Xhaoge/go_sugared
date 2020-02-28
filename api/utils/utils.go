package utils

import (
	"fmt"
	"math/rand"
)

// var (
// 	WorkSpace  string // config
// 	ServerInfo *serverModel
// 	ConfigInfo *configMoudle
// )

// 随机生成str 返回
func GetRandomStr(num int) string {
	strLetter := "qwertyuiopasdfghjklzxcvbnm"
	strNum := "01234567890"
	allStr := strNum + strLetter
	bytes := make([]byte, num)
	for i := 0; i < num; i++ {
		n := rand.Intn(36)
		bytes[i] = allStr[n]
	}
	return string(bytes)
}

// 定义最基本格式返回；
func MakeBaseResponse(code int) {
	fmt.Println("test response")
}
