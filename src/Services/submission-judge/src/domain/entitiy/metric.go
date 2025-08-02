package domain

import (
	"github.com/bibimoni/Online-judge/submission-judge/src/pkg/memory"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Metric struct {
	Id            *bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ProblemId     string         `json:"problem_id,omitempty" bson:"problem_id"`
	CheckerId     *bson.ObjectID `json:"checker_id,omitempty" bson:"checker_id"`
	InteractorId  *bson.ObjectID `json:"interactor_id,omitempty" bson:"interactor_id"`
	TL            int            `json:"tl,omitempty" bson:"tl"`
	ML            memory.Memory  `json:"ml,omitempty" bson:"ml"`
	IsInteractive bool           `json:"is_interactive,omitempty" bson:"is_interactive"`
}
