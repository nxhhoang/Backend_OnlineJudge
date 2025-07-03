package repository

import (
	"context"

	_ "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	"github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/impl"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SubmissionRepository interface {
	CreateSubmission(ctx context.Context, submission *domain.Submission) error
}

func NewSubmissionRepository(db *mongo.Database) SubmissionRepository {
	return impl.NewSubmissionRepositoryImpl(db)
}
