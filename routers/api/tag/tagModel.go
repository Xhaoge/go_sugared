package tag

import "time"

type Tag struct {
	Id 			string		`json:"id,omitempty" bson:"id"`
	Name       	string 		`json:"name" bson:"name"`
	CreatedBy  	time.Time 	`json:"created_by,omitempty" bson:"created_by"`
	ModifiedBy 	time.Time 	`json:"modified_by,omitempty" bson:"modified_by"`
	State      	int    		`json:"state,omitempty" bson:"state"` // 0表示无效，1表示生效
}

