package impl

import (
	"context"

	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	repository "github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/evaluation"
	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
	"github.com/bibimoni/Online-judge/submission-judge/src/pkg/memory"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type EvaluationRepositoryImpl struct {
	collection *mongo.Collection
}

func NewEvaluationRepositoryImpl(db *mongo.Database) *EvaluationRepositoryImpl {
	return &EvaluationRepositoryImpl{
		collection: db.Collection("Evaluation"),
	}
}

func NewEvaluationRepository(db *mongo.Database) repository.EvaluationRepository {
	return NewEvaluationRepositoryImpl(db)
}

func (er *EvaluationRepositoryImpl) CreateEval(ctx context.Context, submissionId string, TL int, ML memory.Memory, nCase int) (string, error) {
	sid, err := bson.ObjectIDFromHex(submissionId)
	if err != nil {
		return "", err
	}

	newEvalRes := domain.EvaluationResult{
		SubmissionId: &sid,
		TL:           TL,
		ML:           ML,
		NCases:       nCase,
	}

	got, err := er.collection.InsertOne(ctx, newEvalRes)
	if err != nil {
		return "", nil
	}

	log := config.GetLogger()
	log.Info().Msgf("Created new eval result with id: [%s] to the database", got.InsertedID.(bson.ObjectID).Hex())
	return got.InsertedID.(bson.ObjectID).Hex(), nil
}

func (er *EvaluationRepositoryImpl) UpdateVerdict(ctx context.Context, evalId string, vert domain.Verdict) error {
	bid, err := bson.ObjectIDFromHex(evalId)
	if err != nil {
		return err
	}

	_, err = er.collection.UpdateOne(ctx, bson.M{"_id": bid}, bson.M{"$set": bson.M{"verdict": vert}})
	if err != nil {
		return err
	}
	return nil
}
