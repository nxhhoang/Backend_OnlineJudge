package domain

import "go.mongodb.org/mongo-driver/v2/bson"

type Batch struct {
	Id             *bson.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	JudgeDataId    *bson.ObjectID  `json:"judge_data_id,omitempty" bson:"judge_data_id"`
	BatchDependsId []bson.ObjectID `json:"batch_depends_id,omitempty" bson:"batch_depends_id"`
}
