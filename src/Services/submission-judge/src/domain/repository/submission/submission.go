package repository

import (
	"context"

	_ "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	"github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/submission/impl"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SubmissionRepository interface {
	CreateSubmission(ctx context.Context, params impl.CreateSubmissionInput) (string, error)
}

func NewSubmissionRepository(db *mongo.Database) SubmissionRepository {
	return impl.NewSubmissionRepositoryImpl(db)
}
