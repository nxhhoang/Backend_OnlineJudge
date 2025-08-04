package repository

import (
	"context"

	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
)

// This will save source code into mongoDB, this is to save user's submission
type SourcecodeRepository interface {
	CreateSourcecode(ctx context.Context, source string, languageId string) (string, error)
	GetSourcecode(ctx context.Context, id string) (*domain.SourceCode, error)
}
