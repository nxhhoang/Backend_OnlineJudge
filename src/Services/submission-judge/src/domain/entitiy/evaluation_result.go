package domain

import (
	"github.com/bibimoni/Online-judge/submission-judge/src/pkg/memory"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type EvaluationResult struct {
	Id              *bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	SubmissionId    *bson.ObjectID `json:"submission_id,omitempty" bson:"submission_id"`
	Verdict         Verdict        `json:"verdict,omitempty" bson:"verdict,omitempty"`
	VerdictCase     []Verdict      `json:"verdict_case,omitempty" bson:"verdict_case,omitempty"`
	CpuTime         int            `json:"cpu_time,omitempty" bson:"cpu_time,omitempty"`
	CpuTimeCase     []int          `json:"cpu_time_case,omitempty" bson:"cpu_time_case,omitempty"`
	MemoryUsage     int            `json:"memory_usage,omitempty" bson:"memory_usage,omitempty"`
	MemoryUsageCase []int          `json:"memory_usage_case,omitempty" bson:"memory_usage_case,omitempty"`
	NSuccess        int            `json:"n_success,omitempty" bson:"n_success,omitempty"`
	Outputs         []string       `json:"outputs,omitempty" bson:"outputs,omitempty"`
	Points          int            `json:"points,omitempty" bson:"points,omitempty"`
	PointsCase      []int          `json:"points_case,omitempty" bson:"points_case,omitempty"`
	TimestampFinish int            `json:"timestamp_finish,omitempty" bson:"timestamp_finish,omitempty"`
	Message         string         `json:"message,omitempty" bson:"message,omitempty"`
	NCases          int            `json:"n_cases,omitempty" bson:"n_cases"`
	TL              int            `json:"tl,omitempty" bson:"tl"`
	ML              memory.Memory  `json:"ml,omitempty" bson:"ml"`
}

type Verdict string

const (
	ACCEPTED              Verdict = "ACCEPTED"
	COMPILATION_ERROR     Verdict = "COMPILATION_ERROR"
	REJECTED              Verdict = "REJECTED"
	RUNTIME_ERROR         Verdict = "RUNTIME_ERROR"
	TIME_LIMIT_EXCEEDED   Verdict = "TIME_LIMIT_EXCEEDED"
	MEMORY_LIMIT_EXCEEDED Verdict = "MEMORY_LIMIT_EXCEEDED"
	WRONG_ANSWER          Verdict = "WRONG_ANSWER"
	JUDGEMENT_FAILED      Verdict = "JUDGEMENT_FAILED"
)
