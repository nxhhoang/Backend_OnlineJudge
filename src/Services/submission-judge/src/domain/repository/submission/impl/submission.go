package impl

import (
	"context"
	"time"

	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SubmissionRepositoryImpl struct {
	collection *mongo.Collection
}

type CreateSubmissionInput struct {
	ProblemId    string
	Username     string
	SourceCodeId string
	Type         domain.SubmissionType
}

func NewSubmissionRepositoryImpl(db *mongo.Database) *SubmissionRepositoryImpl {
	return &SubmissionRepositoryImpl{
		collection: db.Collection("Submission"),
	}
}

func (sr *SubmissionRepositoryImpl) CreateSubmission(ctx context.Context, params CreateSubmissionInput) error {
	sourcecodeId, err := bson.ObjectIDFromHex(params.SourceCodeId)
	if err != nil {
		return err
	}
	newSubmission := domain.Submission{
		Username:     params.Username,
		ProblemId:    params.SourceCodeId,
		SourceCodeId: &sourcecodeId,
		Timestamp:    time.Now(),
		Type:         params.Type,
	}
	got, err := sr.collection.InsertOne(ctx, newSubmission)
	if err != nil {
		return nil
	}

	log := config.GetLogger()
	log.Info().Msgf("Saved submission with id: [%s] to the database", got.InsertedID.(bson.ObjectID).Hex())
	return err
}
