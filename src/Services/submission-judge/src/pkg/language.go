package pkg

import domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"

type Language interface {
	ID() string
	DisplayName() string
	DefaultFileName() string
	Judge(i *domain.Isolate, req *SubmissionRequest) string
}

type SubmissionRequest struct {
	SubmissionId   string
	Username       string
	Sourcecode     string
	SubmissionType domain.SubmissionType
	ProblemId      string
}
