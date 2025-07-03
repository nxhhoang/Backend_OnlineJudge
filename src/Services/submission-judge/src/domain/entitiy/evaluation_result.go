package domain

import "go.mongodb.org/mongo-driver/v2/bson"

type EvaluationResult struct {
	Id              bson.ObjectID `json:"id,omitempty" bson:"_id"`
	SubmissionId    bson.ObjectID `json:"submission_id,omitempty" bson:"submission_id"`
	Verdict         string        `json:"verdict,omitempty" bson:"verdict"`
	VerdictCase     []string      `json:"verdict_case,omitempty" bson:"verdict_case"`
	CpuTime         int           `json:"cpu_time,omitempty" bson:"cpu_time"`
	CpuTimeCase     []int         `json:"cpu_time_case,omitempty" bson:"cpu_time_case"`
	MemoryUsage     int           `json:"memory_usage,omitempty" bson:"memory_usage"`
	MemoryUsageCase []int         `json:"memory_usage_case,omitempty" bson:"memory_usage_case"`
	Message         string        `json:"message,omitempty" bson:"message"`
	NSuccess        int           `json:"n_success,omitempty" bson:"n_success"`
	NCases          int           `json:"n_cases,omitempty" bson:"n_cases"`
	TL              int           `json:"tl,omitempty" bson:"tl"`
	ML              int           `json:"ml,omitempty" bson:"ml"`
	Outputs         []string      `json:"outputs,omitempty" bson:"outputs"`
	TimestampFinish int           `json:"timestamp_finish,omitempty" bson:"timestamp_finish"`
	Points          int           `json:"points,omitempty" bson:"points"`
}
