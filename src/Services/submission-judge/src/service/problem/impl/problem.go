package impl

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
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

func fileExsits(path string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, fmt.Errorf("An error occured when trying to verify if a file exists")
	}
}

func (ps *ProblemServiceImpl) GetTestCaseDirAddr(problemId string, tcType TestCaseType) (string, error) {
	cfg, err := config.Load()
	if err != nil {
		return "", err
	}

	stringAddr := cfg.ProblemsDir + "/" + problemId + "/tests"
	switch tcType {
	case INPUT:
		stringAddr += "/input/"
	case OUTPUT:
		stringAddr += "/output/"
	default:
		return "", fmt.Errorf("Please provide either INPUT or OUTPUT for testcase type")
	}

	stat, err := fileExsits(stringAddr)
	if err != nil {
		return "", err
	}
	if !stat {
		return "", fmt.Errorf("This directory isn't available")
	}
	return stringAddr, nil
}

func (ps *ProblemServiceImpl) GetTestCaseAddr(problemId string, tcType TestCaseType, testNum int) (string, error) {
	stringAddr, err := ps.GetTestCaseDirAddr(problemId, tcType)
	if err != nil {
		return "", err
	}

	conv := strconv.Itoa(testNum)
	if len(conv) == 1 {
		conv = "0" + conv
	}

	stringAddr += conv

	stat, err := fileExsits(stringAddr)
	if err != nil {
		return "", err
	}
	if !stat {
		return "", fmt.Errorf("This test file isn't available")
	}
	return stringAddr, nil
}

func (ps *ProblemServiceImpl) GetCheckerAddr(problemId string) (string, error) {
	cfg, err := config.Load()
	if err != nil {
		return "", err
	}

	stringAddr := cfg.ProblemsDir + "/" + problemId + "/checker"
	stat, err := fileExsits(stringAddr)
	if err != nil {
		return "", err
	}
	if !stat {
		return "", fmt.Errorf("This test file isn't available")
	}
	return stringAddr, nil
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
