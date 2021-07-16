package room

import (
	"fmt"
	"go_sugared/routers/api"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
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
	return c.Remove(bson.M{"packageNumber": singleDoc.PackageNumber})
}

func FindOneRoomByPackageNumber(pkg string) (*Room, error) {
	var result Room
	var err error
	c := connectRoomMgo("hh", "room")
	err = c.Find(bson.M{"packageNumber": pkg}).Select(bson.M{"_id": 0}).One(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func findAllRoomBySelector() ([]Room, error) {
	var result []Room
	var err error
	c := connectRoomMgo("hh", "room")
	err = c.Find(nil).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateRoomBySelector(db, collectin string, update interface{}, pkg string) (interface{}, error) {
	query := bson.M{"packageNumber": pkg}
	err := api.UpdateBySelector(db, collectin, query, update)
	if err != nil {
		return nil, err
	}
	res, err := FindOneRoomByPackageNumber(pkg)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func MakeSelector(s SingleRoomReq) interface{} {
	selector := bson.M{"isInvalid": false}
	if s.PackageNumber != "" {
		selector["packageNumber"] = s.PackageNumber
	}
	if s.NearSubway != "" {
		selector["roomInfo.nearSubway"] = s.NearSubway
	}
	if s.IsInvalid {
		selector["isInvalid"] = true
	}

	fmt.Println("selector: ", selector)
	return selector
}

func MakeRoomUpdate(m Room) interface{} {
	newTime := time.Now()
	update := bson.M{"$set": bson.M{"updateTime": newTime, "roomInfo": m.RoomInfo, "owner": m.Owner}}
	return update
}
