package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

var (
	MgoSession *mgo.Session
	UserMgo    *mgo.Collection
)

type WxUser struct {
	Name     string `bson:"name"`
	Password string `bson:"password"`
}

func InitMongo() {
	var err error
	MgoSession, err = mgo.Dial("")
	//defer mongo.Close()
	if err != nil {
		fmt.Println("connect mongo error：", err)
	}

}

func ConnectUserMgo() *mgo.Collection {
	UserMgo = MgoSession.DB("hh").C("user")
	return UserMgo
}

func InsertUser(data *WxUser) {
	uMgo := ConnectUserMgo()
	err := uMgo.Insert(data)
	if err != nil {
		fmt.Println("insert user failed :", err)
	} else {
		fmt.Println("insert user success; ")
	}

}

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
