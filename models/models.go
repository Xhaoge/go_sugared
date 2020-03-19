package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

var (
	MgoSession *mgo.Session
	UserMgo    *mgo.Collection
)

type OperateDB interface {
	Insert() error
	//Update()
	//Delete()
	//GetAll()
	//GetOneDetails()
	//IsEmpty()
}

type MgoObject struct {
	Control OperateDB
	//MgoS *mgo.Session
	//MgoC  *mgo.Collection
}

func NewMgoObject(op OperateDB) *MgoObject {
	return &MgoObject{
		Control: op,
	}
}

func (this *MgoObject) Insert() error {
	return this.Control.Insert()
}

func ConnectMgo(db, collection string) (*mgo.Session, *mgo.Collection) {
	ms := MgoSession.Copy()
	conn := ms.DB(db).C(collection)
	ms.SetMode(mgo.Monotonic, true)
	return ms, conn
}

func (rdb *Room) Insert() error {
	ms, c := connectMgo("hh", "room")
	defer ms.Close()
	return c.Insert(rdb)
}

func ConnectRoomMgo() *mgo.Collection {
	roomMgo := MgoSession.DB("hh").C("room")
	return roomMgo
}

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
func connectMgo(db, collection string) (*mgo.Session, *mgo.Collection) {
	ms := MgoSession.Copy()
	conn := ms.DB(db).C(collection)
	ms.SetMode(mgo.Monotonic, true)
	return ms, conn
}

func Insert(db, collection string, doc interface{}) error {
	ms, c := connectMgo(db, collection)
	defer ms.Close()
	return c.Insert(doc)
}

func ConnectUserMgo() *mgo.Collection {
	UserMgo = MgoSession.DB("hh").C("user")
	return UserMgo
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
