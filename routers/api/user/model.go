package user


type WxUser struct {
	Name     string `bson:"name"`
	Password string `bson:"password"`
}

type wxLoginReq struct {
	Code string `json:"code"`
}

type toWxLogin struct {
	appid      string
	secret     string
	js_code    string
	grant_type string
}

type UserInfo struct {
	Id          string   `json:"id"`
	OpenId      string   `json:"openId"`
	JsCode 		string	 `json:"code"`
	Username    string   `json:"username"`
	Password    string   `json:"password"`
	Phone       string   `json:"phone"`
	IsAdmin     string   `json:"isAdmin"`
	ContainId   []string `json:"containId"`
	Description string   `json:"description"`
}
