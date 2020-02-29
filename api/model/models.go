package model

// import "golang.org/x/tools/go/ssa/interp"

type RoomBaseResponse struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg`
	Data map[string]string `json:data,omitempty`
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
	Title        string   `json:title`
	PicIdList    []string `json:picIdList`
	Position     string   `json:position`
	Address      string   `json:address`
	RoomType     []int    `json:roomType`
	IsElevator   int      `json:isElevator`
	Price        int      `json:price`
	NearSubway   string   `json:nearSubway`
	Area         float64  `josn:area`
	Floor        int      `json:floor`
	Plot         string   `json:plot`
	Supporting   []string `json:supporting`
	Description  string   `json:description`
	ReleaseTime  string   `json:releaseTime`
	PayType      string   `json:payType`
	ContactPhone string   `json:contactPhone`
	ContactWx    string   `json:contactWx`
	CreatorId    int      `json:creatorId`
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
