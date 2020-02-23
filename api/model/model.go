package model

type BaseResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg`
}

type UserLoginAdd struct {
	Code string `json:"code`
}

type UserUpdate struct {
	UserName string `json:"userName"`
}
