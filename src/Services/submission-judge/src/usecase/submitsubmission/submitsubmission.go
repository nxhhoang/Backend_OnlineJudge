package usecase

import (
	"context"

	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
)

type SubmissionUsecase interface {
	SubmitSubmission(ctx context.Context, input *SubmitSubmissionInput) (output *SubmitSubmissionResponse, err error)
}

type (
	SubmitSubmissionInput struct {
		Username       string                `json:"username,omitempty"`
		ProblemId      string                `json:"problem_id,omitempty"`
		Code           string                `json:"code,omitempty"`
		LanguageId     string                `json:"language,omitempty"`
		SubmissionType domain.SubmissionType `json:"submission_type,omitempty"`
	}

	SubmitSubmissionResponse struct {
		Message string `json:"id"`
		// ID string `json:"id"`
	}
)
