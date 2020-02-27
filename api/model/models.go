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

type PicBaseReq struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type NewRoomAddReq struct {
	Title        string
	PicList      string
	Position     string
	Address      string
	RoomType     string
	IsElevator   int
	Price        int
	NearSubway   string
	Area         float64
	Floor        int
	Plot         string
	Supporting   []string
	Description  string
	ReleaseTime  string
	PayType      string
	ContactPhone string
	ContactWx    string
	CreatorId    int
}

type NewRoom struct {
	TraceId      string
	RoomInfo     *RoomInfo
	PayType      string
	ContactPhone string
	ContactWx    string
	CreatorId    int
}

type RoomInfo struct {
	Title       string
	RoomType    string
	PicList     string
	Position    string
	Address     string
	IsElevator  bool
	Price       int
	NearSubway  string
	Area        string
	Floor       int
	Plot        string
	Supporting  []string
	Description string
	ReleaseTime string
}
