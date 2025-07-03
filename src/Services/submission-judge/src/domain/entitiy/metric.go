package domain

import "go.mongodb.org/mongo-driver/v2/bson"

type Metric struct {
	Id            bson.ObjectID `json:"id,omitempty" bson:"_id"`
	ProblemId     bson.ObjectID `json:"problem_id,omitempty" bson:"problem_id"`
	CheckerId     bson.ObjectID `json:"checker_id,omitempty" bson:"checker_id"`
	InteractorId  bson.ObjectID `json:"interactor_id,omitempty" bson:"interactor_id"`
	TL            int           `json:"tl,omitempty" bson:"tl"`
	ML            int           `json:"ml,omitempty" bson:"ml"`
	IsInteractive bool          `json:"is_interactive,omitempty" bson:"is_interactive"`
}
