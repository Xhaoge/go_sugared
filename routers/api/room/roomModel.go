package room

var (
	ADD_ROOM_ERROR  = 400
	ADD_ROOM_SECESS = 200
)

type Room struct {
	PackageNumber string `json:"packageNumber,omitempty",bson:"packageNumber"`
	UpdateTime    string `json:"updateTime,omitempty",bson:"updateTime"`
	ReleaseTime   string `json:"releaseTime,omitempty",bson:"releaseTime"`
	Isinvalid     bool   `json:"isInvalid",bson:"isInvalid"`
	ReadNum       int    `json:"readNum",bson:"readNum"`
	RoomInfo      `json:"roomInfo",bson:"roomInfo"`
	Owner         `json:"owner",bson:"owner"`
}

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

type SingleRoomReq struct {
	PackageNumber string `json:"packageNumber"`
	IsInvalid     bool   `json:"isInvalid,omitempty"`
}

type GetRoomList struct {
	Code int
	Msg  string
	Data []Room
}
