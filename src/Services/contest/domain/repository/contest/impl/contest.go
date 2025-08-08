package repository

import (
	domain "contest/domain/entity"
	repository "contest/domain/repository/contest"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ContestRepositoryImpl struct {
	collection *mongo.Collection
}

func NewContestRepositoryImpl(db *mongo.Database) *ContestRepositoryImpl {
	return &ContestRepositoryImpl{
		collection: db.Collection("Contest"),
	}
}

func NewContestRepository(db *mongo.Database) repository.ContestRepository {
	return NewContestRepositoryImpl(db)
}

func (cr *ContestRepositoryImpl) Create(ctx context.Context, author uint64) (uint64, error) {
	newContest := domain.Contest{
		Id:          uuid.NewString(),
		Name:        "",
		Description: "",

		Authors:     []string{},
		Curators:    []string{},
		Testers:     []string{},
		Contestants: []domain.Contestant{},

		ProblemLabels: []string{},
		Problems:      []uint64{},

		ScoreboardVisibility: domain.SCOREBOARD_HIDDEN,

		StartTime: time.Now(),
		EndTime:   time.Now(),
	}

	_, err := cr.collection.InsertOne(ctx, newContest)
	if err != nil {
		return 0, err
	}

	return 0, nil
}

func (cr *ContestRepositoryImpl) GetById(ctx context.Context, c *domain.Contest) (*domain.Contest, error) {
	return nil, nil
}

func (cr *ContestRepositoryImpl) AddAuthor(contestId string, authorId uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := cr.collection.UpdateOne(
		ctx,
		bson.M{"id": contestId},
		bson.M{"$push": bson.M{"authors": strconv.Itoa(int(authorId))}},
	)
	if err != nil {
		return err
	}

	log.Info().Msg(fmt.Sprintf("Add author %d to the contest %s", authorId, contestId))

	return nil
}
func (cr *ContestRepositoryImpl) RemoveAuthor(ctx context.Context, contestId string, authorId uint64) error {
	return nil
}
