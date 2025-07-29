package impl

import (
	"context"
	"fmt"
	"time"

	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	"github.com/bibimoni/Online-judge/submission-judge/src/pkg/memory"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/store"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SourcecodeRepositoryImpl struct {
	collection *mongo.Collection
}

func NewSourcecodeRepositoryImpl(db *mongo.Database) *SourcecodeRepositoryImpl {
	return &SourcecodeRepositoryImpl{
		collection: db.Collection("Sourcecode"),
	}
}

func (sr *SourcecodeRepositoryImpl) CreateSourcecode(ctx context.Context, source string, languageId string) (string, error) {
	if !store.DefaultStore.Contains(languageId) {
		return "", fmt.Errorf("Currently there is no support for language %s", languageId)
	}

	newSource := domain.SourceCode{
		LanguageId: languageId,
		CreatedAt:  time.Now(),
		FileSize:   memory.Memory(len(source)),
		SourceCode: source,
	}

	got, err := sr.collection.InsertOne(ctx, newSource)

	if err != nil {
		return "", nil
	}
	return got.InsertedID.(bson.ObjectID).Hex(), nil
}

func (sr *SourcecodeRepositoryImpl) GetSourcecode(ctx context.Context, id string) (*domain.SourceCode, error) {
	var returnSourceCode *domain.SourceCode
	err := sr.collection.FindOne(ctx, bson.M{"_id": id}).Decode(returnSourceCode)
	if err != nil {
		return nil, err
	}
	return returnSourceCode, nil
}
