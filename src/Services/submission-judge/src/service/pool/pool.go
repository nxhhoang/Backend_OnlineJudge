package poolservice

import (
	"fmt"

	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/pool/impl"
)

type PoolService interface {
	Get() (*domain.Isolate, error)
	Put(i *domain.Isolate)
}

func NewPoolSerivce() (PoolService, error) {
	poolService, err := impl.NewPoolServiceImpl()
	if err != nil {
		return nil, fmt.Errorf("Error when create new Pool %v", err)
	}
	return poolService, nil
}
