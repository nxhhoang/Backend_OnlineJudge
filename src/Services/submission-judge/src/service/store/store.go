package store

import (
	"github.com/bibimoni/Online-judge/submission-judge/src/pkg"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/store/impl"
)

type StoreService interface {
	Get(id string) (pkg.Language, error)
	Register(l pkg.Language)
	List() []pkg.Language
}

func NewStoreService() StoreService {
	return impl.NewStoreServiceImpl()
}
