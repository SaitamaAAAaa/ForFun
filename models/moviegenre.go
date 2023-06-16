package models

import "gopkg.in/mgo.v2/bson"

// MovieGenre information
type MovieGenre struct {
	ID          bson.ObjectId `bson:"id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
}
