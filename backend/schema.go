package backend

import "labix.org/v2/mgo/bson"

type Named struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Problem struct {
	Id     string          `json:"id"`
	Name   string          `json:"name"`
	Brief  string          `json:"brief"`
	Tagged []bson.ObjectId `json:"tagged"`
}

type Tag struct {
	ObjId  bson.ObjectId   "_id"
	Id     string          `json:"id"`
	Name   string          `json:"name"`
	Tagged []bson.ObjectId `json:"tagged"`
}
