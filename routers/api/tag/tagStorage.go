package tag

import (
	"go_sugared/routers/api"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// 数据库相关操作
func connectTagMgo(db, cl string) *mgo.Collection {
	collection := api.MgoSession.DB(db).C(cl)
	return collection
}

func (doc *Tag) Insert() error {
	c := connectTagMgo("hh", "tag")
	//defer api.MgoSession.Close()
	return c.Insert(doc)
}

func findOneById(pkg string) (*Tag,error) {
	var tag Tag
	var err error
	c := connectTagMgo("hh","tag")
	err = c.Find(bson.M{"id":pkg}).Select(bson.M{"_id":0}).One(&tag)
	if err != nil {
		return nil, err
	}
	return &tag, nil
}