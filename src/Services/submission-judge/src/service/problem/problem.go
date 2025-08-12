package problem

import (
	"context"
)

// This will handle the calling process, the retrieving process of getting problem
// served by the Problem service (the name might be a little confusing)
type ProblemService interface {
	Get(ctx context.Context, id string) (*ProblemServiceGetOutput, error)
	GetTestCaseAddr(problemId string, tcType TestCaseType, testNum int) (string, error)
	GetTestCaseDirAddr(problemId string, tcType TestCaseType) (string, error)
	GetCheckerAddr(problemId string) (string, error)
}

type TestCaseType string

const (
	INPUT  TestCaseType = "INPUT"
	OUTPUT TestCaseType = "OUTPUT"
)

type ProblemServiceGetOutput struct {
	ID          string   `json:"ID,omitempty"`
	ProblemId   int64    `json:"problem-id,omitempty"`
	Name        string   `json:"name,omitempty"`
	ShortName   string   `json:"short-name,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	TestNum     int      `json:"test-num,omitempty"`
	TimeLimit   int      `json:"time-limit,omitempty"`
	MemoryLimit int64    `json:"memory-limit,omitempty"`
}
