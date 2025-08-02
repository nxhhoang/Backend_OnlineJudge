package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/bibimoni/Online-judge/submission-judge/src/common"
	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
)

const PROBLEM_INFO_FILENAME = "problem.json"

type ProblemServiceImpl struct {
	problemServerAddr string
}

func NewProblemServiceImpl() (*ProblemServiceImpl, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("Failed to load config for PROBLEM SERVICE: %v", err)
	}
	return &ProblemServiceImpl{
		problemServerAddr: cfg.ProblemServerAddr,
	}, nil
}

func (ps *ProblemServiceImpl) Get(ctx context.Context, id string) (*ProblemServiceGetOutput, error) {
	req := common.APIRequest{
		Method:  "GET",
		URL:     ps.problemServerAddr + "get/" + id + "/" + PROBLEM_INFO_FILENAME,
		Timeout: 60 * time.Second,
	}

	result, err := common.SendRequest[ProblemServiceGetOutput](ctx, req)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, fmt.Errorf("There is an error occured fetch requesting from PROBLEM SERVER")
	}
	return result, nil
}

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
