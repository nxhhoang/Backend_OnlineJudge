package impl

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/bibimoni/Online-judge/submission-judge/src/common"
	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/problem"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/problem/utils"
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

func NewProblemService() (problem.ProblemService, error) {
	return NewProblemServiceImpl()
}

func (ps *ProblemServiceImpl) Get(ctx context.Context, id string) (*problem.ProblemServiceGetOutput, error) {
	req := common.APIRequest{
		Method:  "GET",
		URL:     ps.problemServerAddr + "get/" + id + "/" + PROBLEM_INFO_FILENAME,
		Timeout: 60 * time.Second,
	}

	result, err := common.SendRequest[problem.ProblemServiceGetOutput](ctx, req)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, fmt.Errorf("There is an error occured fetch requesting from PROBLEM SERVER")
	}
	return result, nil
}

func (ps *ProblemServiceImpl) GetTestCaseDirAddr(problemId string, tcType problem.TestCaseType) (string, error) {
	cfg, err := config.Load()
	if err != nil {
		return "", err
	}

	stringAddr := cfg.ProblemsDir + "/" + problemId + "/tests"
	switch tcType {
	case problem.INPUT:
		stringAddr += "/input/"
	case problem.OUTPUT:
		stringAddr += "/output/"
	default:
		return "", fmt.Errorf("Please provide either INPUT or OUTPUT for testcase type")
	}

	stat, err := utils.FileExsits(stringAddr)
	if err != nil {
		return "", err
	}
	if !stat {
		return "", fmt.Errorf("This directory isn't available")
	}
	return stringAddr, nil
}

func (ps *ProblemServiceImpl) GetTestCaseAddr(problemId string, tcType problem.TestCaseType, testNum int) (string, error) {
	stringAddr, err := ps.GetTestCaseDirAddr(problemId, tcType)
	if err != nil {
		return "", err
	}

	conv := strconv.Itoa(testNum)
	if len(conv) == 1 {
		conv = "0" + conv
	}

	stringAddr += conv

	stat, err := utils.FileExsits(stringAddr)
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
	stat, err := utils.FileExsits(stringAddr)
	if err != nil {
		return "", err
	}
	if !stat {
		return "", fmt.Errorf("This test file isn't available")
	}
	return stringAddr, nil
}

func (ps *ProblemServiceImpl) GetAllTestcaseAddr(problemId string, problemInfo *problem.ProblemServiceGetOutput) ([]string, error) {

}
