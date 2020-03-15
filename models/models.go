package models

type WxUser struct {
	Name     string `boson:"name"`
	Password string `boson:"password"`
}

//func InitMongo() {
//	mongo, err := mgo.Dial("")
//	defer mongo.Close()
//	if err != nil {
//		fmt.Println("connect mongo error：", err)
//	}
//	client := mongo.DB("hh").C("user")
//
//	// 创建数据
//	data := WxUser{
//		Name:     "xhaoge",
//		Password: "123456",
//	}
//	// 插入数据
//	cErr := client.Insert(&data)
//	if cErr != nil {
//		fmt.Println(cErr)
//	}
//}

type UserLoginAdd struct {
	Code string `json:"code"`
}

type UserInfo struct {
	Id          string   `json:"id"`
	OpenId      string   `json:"openId"`
	Username    string   `json:"username"`
	Password    string   `json:"password"`
	Phone       string   `json:"phone"`
	IsAdmin     string   `json:"isAdmin"`
	ContainId   []string `json:"containId"`
	Description string   `json:"description"`
}

type UserUpdate struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Phone       string `json:"phone"`
	IsAdmin     string `json:"isAdmin"`
	Description string `json:"description"`
}
