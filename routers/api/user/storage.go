package user

import (
	"go_sugared/routers/api"
	"gopkg.in/mgo.v2"
)

func connectUserMgo(db, cl string) *mgo.Collection {
	collection := api.MgoSession.DB(db).C(cl)
	return collection
}

func newUserInfo(jscode,openid string) *UserInfo {
	return &UserInfo{
		OpenId:openid,
		JsCode:jscode,
	}
}

func (u *UserInfo)Insert() error {
	c := connectUserMgo("hh", "user")
	//defer api.MgoSession.Close()
	return c.Insert(u)
}
