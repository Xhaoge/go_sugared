package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

var RoomMgo *mgo.Collection

func connectRoomMgo() *mgo.Collection {
	RoomMgo = MgoSession.DB("hh").C("room")
	return RoomMgo
}

//func InsertRoom(data *Room) {
//	rMo := ConnectRoomMgo()
//	rerr := rMo.Insert(&data)
//	if rerr != nil {
//		fmt.Println("insert room failed :", rerr)
//	} else {
//		fmt.Println("insert room success; ")
//	}
//}

func GetRoomDetail() {

}

type Room struct {
	PackageNumber string `json:"packageNumber,omitempty",bson:"packageNumber"`
	UpdateTime    string `json:"updateTime,omitempty",bson:"updateTime"`
	ReleaseTime   string `json:"releaseTime,omitempty",bson:"releaseTime"`
	RoomInfo      `json:"roomInfo",bson:"roomInfo"`
	Owner         `json:"owner",bson:"owner"`
}

//func (doc *Room)Insert() error {
//	ms := connectRoomMgo()
//	defer MgoSession.Close()
//	return ms.Insert(doc)
//}

type RoomInfo struct {
	Title       string   `json:"title",bson:"title"`
	PicIdList   []string `json:"picIdList",bson:"picIdList"`
	Position    string   `json:"position",bson:"position"`
	Address     string   `json:"address",bson:"address"`
	RoomType    []int    `json:"roomType",bson:"roomType"`
	IsElevator  int      `json:"isElevator",bson:"isElevator"`
	Price       int      `json:"price",bson:"price"`
	NearSubway  string   `json:"nearSubway",bson:"nearSubway"`
	Area        float64  `josn:"area",bson:"area"`
	Floor       int      `json:"floor",bson:"floor"`
	Plot        string   `json:"plot",bson:"plot"`
	Supporting  []string `json:"supporting,omitempty",bson:"supporting"`
	Description string   `json:"description,omitempty",bson:"description"`
	ReleaseTime string   `json:"releaseTime,omitempty",bson:"releaseTime"`
	PayType     string   `json:"payType,omitempty",bson:"payType"`
}

type Owner struct {
	CreatorId    string `json:"creatorId",bson:"creatorId"`
	ContactPhone string `json:"contactPhone,omitempty",bson:"contactPhone"`
	ContactWx    string `json:"contactWx,omitempty",bson:"contactWx"`
}

type RoomBaseResponse struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data,omitempty"`
}

type NewRoom struct {
	TraceId      string
	RoomInfo     *RoomInfo
	PayType      string
	ContactPhone string
	ContactWx    string
	CreatorId    int
}

type NewRoomAddReq struct {
	Title        string   `json:"title"`
	PicIdList    []string `json:"picIdList"`
	Position     string   `json:"position"`
	Address      string   `json:"address"`
	RoomType     []int    `json:"roomType"`
	IsElevator   int      `json:"isElevator"`
	Price        int      `json:"price"`
	NearSubway   string   `json:"nearSubway"`
	Area         float64  `josn:"area"`
	Floor        int      `json:"floor"`
	Plot         string   `json:"plot"`
	Supporting   []string `json:"supporting"`
	Description  string   `json:"description"`
	ReleaseTime  string   `json:"releaseTime"`
	PayType      string   `json:"payType"`
	ContactPhone string   `json:"contactPhone"`
	ContactWx    string   `json:"contactWx"`
	CreatorId    int      `json:"creatorId"`
}

func (r *Room) ToPrint() {
	fmt.Println("roomAddReq title: ", r.RoomInfo.Title)
	fmt.Println("roomAddReq PicIdList: ", r.RoomInfo.PicIdList)
	fmt.Println("roomAddReq Address: ", r.RoomInfo.Address)
	fmt.Println("roomAddReq IsElevator: ", r.RoomInfo.IsElevator)
	fmt.Println("roomAddReq ContactPhone: ", r.Owner.ContactPhone)
	fmt.Println("roomAddReq Supporting: ", r.RoomInfo.Supporting)
	fmt.Println("roomAddReq Price: ", r.RoomInfo.Price)
}
