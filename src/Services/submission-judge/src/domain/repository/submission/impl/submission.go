package impl

import (
	"context"
	"time"

	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	repository "github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/submission"
	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
	"go.mongodb.org/mongo-driver/v2/bson"
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

func NewSubmissionRepository(db *mongo.Database) repository.SubmissionRepository {
	return NewSubmissionRepositoryImpl(db)
}

func (sr *SubmissionRepositoryImpl) CreateSubmission(ctx context.Context, params repository.CreateSubmissionInput) (string, error) {
	// sourcecodeId, err := bson.ObjectIDFromHex(params.SourceCodeId)
	// if err != nil {
	// 	return "", err
	// }
	// evalId, err := bson.ObjectIDFromHex(params.EvalId)
	// if err != nil {
	// 	return "", err
	// }
	newSubmission := domain.Submission{
		Username:  params.Username,
		ProblemId: params.ProblemId,
		Timestamp: time.Now(),
		Type:      params.Type,
	}
	got, err := sr.collection.InsertOne(ctx, newSubmission)
	if err != nil {
		return "", nil
	}

	log := config.GetLogger()
	log.Info().Msgf("Saved submission with id: [%s] to the database", got.InsertedID.(bson.ObjectID).Hex())
	return got.InsertedID.(bson.ObjectID).Hex(), nil
}

func (sr *SubmissionRepositoryImpl) FindSubmission(ctx context.Context, submissionId string) (*domain.Submission, error) {
	bId, err := bson.ObjectIDFromHex(submissionId)
	if err != nil {
		return nil, err
	}

	var returnSubmission domain.Submission
	err = sr.collection.FindOne(ctx, bson.M{"_id": bId}).Decode(&returnSubmission)

	if err != nil {
		return nil, err
	}
	return &returnSubmission, nil
}
