package repository

import (
	"context"

	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	"github.com/bibimoni/Online-judge/submission-judge/src/pkg/memory"
)

type EvaluationRepository interface {
	CreateEval(ctx context.Context, submissionId string, TL int, ML memory.Memory, nCase int) (string, error)
	UpdateVerdict(ctx context.Context, evalId string, vert domain.Verdict) error
}
