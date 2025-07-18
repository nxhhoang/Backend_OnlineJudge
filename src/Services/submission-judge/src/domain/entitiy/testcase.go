package domain

import "go.mongodb.org/mongo-driver/v2/bson"

type TestCase struct {
	Id          bson.ObjectID `json:"id,omitempty" bson:"_id"`
	BatchId     bson.ObjectID `json:"batch_id,omitempty" bson:"batch_id"`
	ProblemId   bson.ObjectID `json:"problem_id,omitempty" bson:"problem_id"`
	JudgeDataId bson.ObjectID `json:"judge_data_id,omitempty" bson:"judge_data_id"`
	Input       string        `json:"input,omitempty" bson:"input"`
	Output      string        `json:"output,omitempty" bson:"output"`
	Points      int           `json:"points,omitempty" bson:"points"`
}
