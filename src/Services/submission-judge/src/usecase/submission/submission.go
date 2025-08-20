package usecase

import (
	"context"
	"time"

	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	"github.com/bibimoni/Online-judge/submission-judge/src/pkg/memory"
)

type SubmissionUsecase interface {
	SubmitSubmission(ctx context.Context, input *SubmitSubmissionInput) (output *SubmitSubmissionResponse, err error)
	GetSubmission(ctx context.Context, input *GetSubmissionInput) (output *GetSubmissionOutput, err error)
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
		Message string `json:"message"`
		ID      string `json:"id"`
	}

	GetSubmissionInput struct {
		SubmissionId string
	}

	GetSubmissionOutput struct {
		ProblemId       string                  `json:"problem_id,omitempty"`
		Verdict         domain.Verdict          `json:"verdict,omitempty"`
		VerdictCase     []domain.Verdict        `json:"verdict_case,omitempty"`
		CpuTime         float64                 `json:"cpu_time,omitempty"`
		CpuTimeCase     []float64               `json:"cpu_time_case,omitempty"`
		MemoryUsage     memory.Memory           `json:"memory_usage,omitempty"`
		MemoryUsageCase []memory.Memory         `json:"memory_usage_case,omitempty"`
		NSuccess        int                     `json:"n_success,omitempty"`
		Outputs         []string                `json:"outputs,omitempty"`
		Points          int                     `json:"points,omitempty"`
		PointsCase      []int                   `json:"points_case,omitempty"`
		Message         string                  `json:"message,omitempty"`
		NCases          int                     `json:"n_cases,omitempty"`
		TL              int                     `json:"tl,omitempty"`
		ML              memory.Memory           `json:"ml,omitempty"`
		Username        string                  `json:"username,omitempty"`
		Timestamp       time.Time               `json:"timestamp"`
		Type            domain.SubmissionType   `json:"type,omitempty"`
		Language        string                  `json:"language,omitempty"`
		SourceCode      string                  `json:"source_code,omitempty"`
		EvalStatus      domain.SubmissionStatus `json:"eval_status,omitempty"`
	}
)
