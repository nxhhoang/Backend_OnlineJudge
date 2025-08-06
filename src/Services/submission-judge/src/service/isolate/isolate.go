package isolateservice

import (
	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
)

type IsolateService interface {
	NewIsolate(id int) (*domain.Isolate, error)
	Cleanup(i *domain.Isolate) error
	Init(i *domain.Isolate) error
	Run(i *domain.Isolate, rc domain.RunConfig, req *SubmissionRequest, toRun string, toRunArgs ...string) error
}

type SubmissionRequest struct {
	SubmissionId   string
	Username       string
	Sourcecode     string
	SubmissionType domain.SubmissionType
	ProblemId      string
	IService       IsolateService
	LanguageId     string
}
