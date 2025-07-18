package domain

import "go.mongodb.org/mongo-driver/v2/bson"

type Problem struct {
	Id     bson.ObjectID `bson:"_id" json:"id,omitempty"`
	Index  string        `bson:"index" json:"index,omitempty"`
	Name   string        `bson:"name" json:"name,omitempty"`
	Points int           `bson:"points" json:"points,omitempty"`
	Rating int           `bson:"rating" json:"rating,omitempty"`
	Tags   []string      `bson:"tags" json:"tags,omitempty"`
}
