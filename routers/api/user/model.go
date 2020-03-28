package user

import (
	"go_sugared/routers/api"
	"gopkg.in/mgo.v2"
)

func ConnectMgo(db, collection string) (*mgo.Session, *mgo.Collection) {
	ms := api.MgoSession.Copy()
	conn := ms.DB(db).C(collection)
	ms.SetMode(mgo.Monotonic, true)
	return ms, conn
}

func ConnectRoomMgo() *mgo.Collection {
	roomMgo := api.MgoSession.DB("hh").C("room")
	return roomMgo
}

type WxUser struct {
	Name     string `bson:"name"`
	Password string `bson:"password"`
}

func connectMgo(db, collection string) (*mgo.Session, *mgo.Collection) {
	ms := api.MgoSession.Copy()
	conn := ms.DB(db).C(collection)
	ms.SetMode(mgo.Monotonic, true)
	return ms, conn
}

func Insert(db, collection string, doc interface{}) error {
	ms, c := connectMgo(db, collection)
	defer ms.Close()
	return c.Insert(doc)
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
	Username    string   `json:"username"`
	Password    string   `json:"password"`
	Phone       string   `json:"phone"`
	IsAdmin     string   `json:"isAdmin"`
	ContainId   []string `json:"containId"`
	Description string   `json:"description"`
}
