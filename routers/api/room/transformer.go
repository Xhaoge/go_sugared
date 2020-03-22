package room

import (
	"fmt"
	"go_sugared/pkg/util"
	"go_sugared/routers/api"
	"gopkg.in/mgo.v2"
)

// 数据库相关操作
func connectRoomMgo(db, cl string) *mgo.Collection {
	collection := api.MgoSession.DB(db).C(cl)
	return collection
}

func (doc *Room) Insert() error {
	c := connectRoomMgo("hh", "room")
	//defer api.MgoSession.Close()
	doc.PackageNumber = util.MakePackageNumber()
	return c.Insert(doc)
}

func (doc *Room) Delete() error {
	c := connectRoomMgo("hh", "room")
	defer api.MgoSession.Close()
	return c.Insert(doc)
}

// 数据处理相关
func (doc *Room) ToPrint() {
	fmt.Println("add room param: ", doc)
	//fmt.Println("roomAddReq title: ", r.RoomInfo.Title)
	//fmt.Println("roomAddReq PicIdList: ", r.RoomInfo.PicIdList)
	//fmt.Println("roomAddReq Address: ", r.RoomInfo.Address)
	//fmt.Println("roomAddReq IsElevator: ", r.RoomInfo.IsElevator)
	//fmt.Println("roomAddReq ContactPhone: ", r.Owner.ContactPhone)
	//fmt.Println("roomAddReq Supporting: ", r.RoomInfo.Supporting)
	//fmt.Println("roomAddReq Price: ", r.RoomInfo.Price)
}
