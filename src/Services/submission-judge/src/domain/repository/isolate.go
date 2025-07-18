package repository

import (
	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	"github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/impl"
)

type IsolateRepository interface {
	NewIsolate(id int) (*domain.Isolate, error)
}

func NewIsolateRepository() IsolateRepository {
	return impl.NewIsolateRepositoryImpl()
}
