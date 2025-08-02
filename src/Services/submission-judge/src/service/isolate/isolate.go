package isolateservice

import (
	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/isolate/impl"
)

type IsolateService interface {
	NewIsolate(id int) (*domain.Isolate, error)
	Cleanup(i *domain.Isolate) error
	Init(i *domain.Isolate) error
	Judge(i *domain.Isolate, rc *domain.RunConfig)
	RunBinary(i *domain.Isolate, rc *domain.RunConfig)
}

func NewIsolateService() IsolateService {
	return impl.NewIsolateServiceImpl()
}
