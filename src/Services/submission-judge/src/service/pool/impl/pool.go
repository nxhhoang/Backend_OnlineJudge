package impl

import (
	"errors"

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

	is, err := isolateservice.NewIsolateService()
	if err != nil {
		return nil, err
	}

	newPool := &PoolServiceImpl{
		pool: &domain.Pool{
			Isolates: make(chan *domain.Isolate, cfg.Judge.Amount),
		},
		isolateService: is,
	}

	log := config.GetLogger()
	log.Info().Msgf("Offset: %d, amount: %d", cfg.Judge.IDOffset, cfg.Judge.Amount)
	// Init all isolate
	for i := cfg.Judge.IDOffset; i < (cfg.Judge.IDOffset + cfg.Judge.Amount); i++ {
		newIsolate, err := is.NewIsolate(i)
		if err != nil {
			return nil, err
		}
		is.Init(newIsolate)
		newPool.Put(newIsolate)
	}

	log.Info().Msgf("Finished initialized pool service")

	return newPool, nil
}

func (ps *PoolServiceImpl) Get() (*domain.Isolate, error) {
	i, ok := <-ps.pool.Isolates
	if !ok {
		return nil, errors.New("Channel is closed")
	}
	return i, nil
}

func (ps *PoolServiceImpl) Put(i *domain.Isolate) {
	ps.pool.Isolates <- i
}
