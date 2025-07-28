package impl

import (
	"context"

	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SubmissionRepositoryImpl struct {
	collection *mongo.Collection
}

func NewSubmissionRepositoryImpl(db *mongo.Database) *SubmissionRepositoryImpl {
	return &SubmissionRepositoryImpl{
		collection: db.Collection("Submission"),
	}
}

func (sr *SubmissionRepositoryImpl) CreateSubmission(ctx context.Context, submission *domain.Submission) error {
	_, err := sr.collection.InsertOne(ctx, submission)
	return err
}
