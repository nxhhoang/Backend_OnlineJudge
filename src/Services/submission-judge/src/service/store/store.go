package store

import (
	"github.com/bibimoni/Online-judge/submission-judge/src/pkg"
	"github.com/bibimoni/Online-judge/submission-judge/src/pkg/langs/cpp"
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

func NewStoreWithDefaultLangs() StoreService {
	storeService := NewStoreService()
	for _, option := range cpp.GetAllOptions() {
		storeService.Register(option)
	}
	return storeService
}

var DefaultStore StoreService = nil

func init() {
	DefaultStore = NewStoreWithDefaultLangs()
}
