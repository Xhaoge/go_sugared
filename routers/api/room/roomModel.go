package room

import "time"

type Room struct {
	PackageNumber 	string    	`json:"packageNumber,omitempty" bson:"packageNumber"`
	UpdateTime    	time.Time 	`json:"updateTime,omitempty" bson:"updateTime"`
	ReleaseTime   	time.Time 	`json:"releaseTime,omitempty" bson:"releaseTime"`
	Isinvalid     	bool      	`json:"isInvalid" bson:"isInvalid"`
	ReadNum       	int       	`json:"readNum" bson:"readNum"`
	RenterId		[]string	`json:"renterId,omitempty" bson:"renterId"`
	TagId   		[]string 	`json:"tagId,omitempty" bson:"tagId"`
	RoomInfo      				`json:"roomInfo" bson:"roomInfo"`
	Owner         				`json:"owner" bson:"owner"`
}

type RoomInfo struct {
	Title       string   `json:"title" bson:"title"`
	Price       int      `json:"price" bson:"price"`
	PayType     string   `json:"payType,omitempty" bson:"payType"`
	NearSubway  string   `json:"nearSubway" bson:"nearSubway"`
	Address     string   `json:"address" bson:"address"`
	Plot        string   `json:"plot" bson:"plot"`
	RoomType    []int    `json:"roomType" bson:"roomType"`
	IsElevator  bool     `json:"isElevator" bson:"isElevator"`
	Area        float64  `josn:"area" bson:"area"`
	Floor       int      `json:"floor" bson:"floor"`
	Position    string   `json:"position" bson:"position"`
	Supporting  []string `json:"supporting,omitempty" bson:"supporting"`
	PicIdList   []string `json:"picIdList" bson:"picIdList"`
	Description string   `json:"description,omitempty" bson:"description"`
}

type Owner struct {
	CreatorId    string `json:"creatorId" bson:"creatorId"`
	ContactPhone string `json:"contactPhone,omitempty" bson:"contactPhone"`
	ContactWx    string `json:"contactWx,omitempty" bson:"contactWx"`
}

type SingleRoomReq struct {
	PackageNumber string `json:"packageNumber,omitempty"`
	IsInvalid     bool   `json:"isInvalid,omitempty"`   // 是否已失效
	NearSubway    string `json:"nearSubway,omitempty"` // 临近地铁
	PriceSort     int    `json:"priceSort,omitempty"`  // 价格排序，0 是默认排序，1 是由高到低，-1是由低到高；
}

type GetRoomList struct {
	Code int
	Msg  string
	Data []Room
}
