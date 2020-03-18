package models

import (
	"gopkg.in/mgo.v2"
)

var RoomMgo *mgo.Collection

func ConnectRoomMgo() *mgo.Collection {
	RoomMgo = MgoSession.DB("hh").C("user")
	return RoomMgo
}

type Room struct {
	PackageNumber string `json:"packageNumber,omitempty"`
	UpdateTime    string `json:"updateTime,omitempty"`
	ReleaseTime   string `json:"releaseTime,omitempty"`
	RoomInfo      `json:"roomInfo"`
	Owner         `json:"owner"`
}

type RoomInfo struct {
	Title       string   `json:"title"`
	PicIdList   []string `json:"picIdList"`
	Position    string   `json:"position"`
	Address     string   `json:"address"`
	RoomType    []int    `json:"roomType"`
	IsElevator  int      `json:"isElevator"`
	Price       int      `json:"price"`
	NearSubway  string   `json:"nearSubway"`
	Area        float64  `josn:"area"`
	Floor       int      `json:"floor"`
	Plot        string   `json:"plot"`
	Supporting  []string `json:"supporting,omitempty"`
	Description string   `json:"description,omitempty"`
	ReleaseTime string   `json:"releaseTime,omitempty"`
	PayType     string   `json:"payType,omitempty"`
}

type Owner struct {
	CreatorId    string `json:"creatorId"`
	ContactPhone string `json:"contactPhone,omitempty"`
	ContactWx    string `json:"contactWx,omitempty"`
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
