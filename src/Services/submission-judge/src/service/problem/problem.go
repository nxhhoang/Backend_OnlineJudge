package problem

import (
	"context"

	"github.com/bibimoni/Online-judge/submission-judge/src/service/problem/impl"
)

// This will handle the calling process, the retrieving process of getting problem
// served by the Problem service (the name might be a little confusing)
type ProblemService interface {
	Get(ctx context.Context, id string) (*impl.ProblemServiceGetOutput, error)
	GetTestCaseAddr(problemId string, tcType impl.TestCaseType, testNum int) (string, error)
	GetTestCaseDirAddr(problemId string, tcType impl.TestCaseType) (string, error)
	GetCheckerAddr(problemId string) (string, error)
}

func NewProblemService() (ProblemService, error) {
	return impl.NewProblemServiceImpl()
}
