package domain

import "go.mongodb.org/mongo-driver/v2/bson"

type SourceCode struct {
	Id        bson.ObjectID `json:"id,omitempty" bson:"_id"`
	Language  string        `json:"language,omitempty" bson:"language"`
	CreatedAt string        `json:"create_at,omitempty" bson:"create_at"`
	Name      string        `json:"name,omitempty" bson:"name"`
	FileSize  int           `json:"file_size,omitempty" bson:"file_size"`
}
