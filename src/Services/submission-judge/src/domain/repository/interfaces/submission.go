package interfaces

import (
	"context"
	_ "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
)

type SubmissionRepository interface {
	Create(ctx context.Context) error
}
