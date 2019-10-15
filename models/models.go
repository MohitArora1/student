package models

import "github.com/globalsign/mgo/bson"

//Student model
type Student struct {
	ID     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name"`
	Age    int           `json:"age"`
	Class  string        `json:"class"`
	RollNo int           `json:"rollno"`
}
