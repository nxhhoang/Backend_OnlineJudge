package isolateservice

import (
	"context"

	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/isolate/impl"
)

type IsolateService interface {
	NewIsolate(id int) (*domain.Isolate, error)
	Cleanup(i *domain.Isolate, _ context.Context) error
	Init(i *domain.Isolate, ctx context.Context) error
}

func NewIsolateService() IsolateService {
	return impl.NewIsolateServiceImpl()
}
