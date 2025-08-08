package repository

import (
	domain "contest/domain/entity"
	"context"
)

type ContestRepository interface {
	GetById(ctx context.Context, c *domain.Contest) (*domain.Contest, error)
	Create(ctx context.Context, author uint64) (uint64, error)

	AddAuthor(contestId string, authorId uint64) error
	RemoveAuthor(ctx context.Context, contestId string, authorId uint64) error
}
