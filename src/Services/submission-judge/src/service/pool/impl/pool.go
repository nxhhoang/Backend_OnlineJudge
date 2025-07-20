package impl

import (
	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
	isolateservice "github.com/bibimoni/Online-judge/submission-judge/src/service/isolate"
)

type PoolServiceImpl struct {
	pool           *domain.Pool
	isolateService isolateservice.IsolateService
}

func NewPoolServiceImpl() (*PoolServiceImpl, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	is := isolateservice.NewIsolateService()

	newPool := &PoolServiceImpl{
		pool: &domain.Pool{
			Isolates: make(chan *domain.Isolate, cfg.Judge.Amount),
		},
		isolateService: is,
	}

	for i := cfg.Judge.IDOffset; i < cfg.Judge.IDOffset+cfg.Judge.Amount; i++ {
		newIsolate, err := is.NewIsolate(i)
		if err != nil {
			return nil, err
		}
		newPool.Put(newIsolate)
	}
	return newPool, nil
}

func (ps *PoolServiceImpl) Get() (*domain.Isolate, error) {
	return <-ps.pool.Isolates, nil
}

func (ps *PoolServiceImpl) Put(i *domain.Isolate) {
	ps.pool.Isolates <- i
}
