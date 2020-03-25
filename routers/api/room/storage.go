package room

import (
	"fmt"
	"go_sugared/routers/api"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// 数据库相关操作
func connectRoomMgo(db, cl string) *mgo.Collection {
	collection := api.MgoSession.DB(db).C(cl)
	return collection
}

func (doc *Room) Insert() error {
	c := connectRoomMgo("hh", "room")
	//defer api.MgoSession.Close()
	return c.Insert(doc)
}

func (singleDoc *SingleRoomReq) Delete() error {
	c := connectRoomMgo("hh", "room")
	//defer api.MgoSession.Close()
	return c.Remove(bson.M{"packagenumber": singleDoc.PackageNumber})
}

func FindOneRoomByPackageNumber(pkg string) (interface{}, error) {
	var result Room
	var err error
	c := connectRoomMgo("hh", "room")
	err = c.Find(bson.M{"packagenumber": pkg}).Select(bson.M{"_id": 0}).One(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func MakeSelector(s SingleRoomReq) interface{} {
	selector := bson.M{"isinvalid": false}
	if s.PackageNumber != "" {
		selector["packagenumber"] = s.PackageNumber
	} else {
		if s.NearSubway != "" {
			selector["roominfo.nearsubway"] = s.NearSubway
		}
		if s.IsInvalid {
			selector["isinvalid"] = true
		}
	}
	fmt.Println("selector: ", selector)
	return selector
}
