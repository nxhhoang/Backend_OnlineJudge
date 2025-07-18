package impl

import (
	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
)

type IsolateRepositoryImpl struct {
}

func NewIsolateRepositoryImpl() *IsolateRepositoryImpl {
	return &IsolateRepositoryImpl{}
}

func (ir *IsolateRepositoryImpl) NewIsolate(id int) (*domain.Isolate, error) {
	res := &domain.Isolate{
		ID: id,
	}
	if res.Logger == nil {
		res.Logger = config.GetLogger()
	}
	return res, nil
}
