package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Problem struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ProblemId uint64             `json:"problem-id",bson"problem-id"`
	Name      string             `json:"name",bson:"name"`
	ShortName string             `json:"short-name",bson:"short-name"`
	Tags      []string           `json:"tags",bson:"tags"`
}
