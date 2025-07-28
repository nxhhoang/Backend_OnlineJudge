package domain

import "go.mongodb.org/mongo-driver/v2/bson"

type Submission struct {
	Id           *bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	JudgeDataId  *bson.ObjectID `json:"judge_data_id,omitempty" bson:"judge_data_id"`
	MetricId     *bson.ObjectID `json:"metric_id,omitempty" bson:"metric_id"`
	Username     string         `json:"username,omitempty" bson:"username"`
	ProblemId    *bson.ObjectID `json:"problem_id,omitempty" bson:"problem_id"`
	SourceCodeId *bson.ObjectID `json:"source_code_id,omitempty" bson:"source_code_id"`
	Timestamp    int            `json:"timestamp,omitempty" bson:"timestamp"`
	Type         SubmissionType `json:"type,omitempty" bson:"type"`
}

type SubmissionType string

const (
	CUSTOM SubmissionType = "CUSTOM"
	ACTUAL SubmissionType = "ACTUAL"
)
