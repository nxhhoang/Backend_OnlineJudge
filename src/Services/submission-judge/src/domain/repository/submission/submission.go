package repository

import (
	"context"

	_ "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
)

type SubmissionRepository interface {
	CreateSubmission(ctx context.Context, params CreateSubmissionInput) (string, error)
	FindSubmission(ctx context.Context, submissionId string) (*domain.Submission, error)
}

type CreateSubmissionInput struct {
	ProblemId string
	Username  string
	Type      domain.SubmissionType
}
