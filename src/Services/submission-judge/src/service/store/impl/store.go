package impl

import (
	"github.com/bibimoni/Online-judge/submission-judge/src/pkg"
	"github.com/bibimoni/Online-judge/submission-judge/src/pkg/langs/cpp"
	"github.com/bibimoni/Online-judge/submission-judge/src/pkg/langs/python3"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/store"
)

type StoreServiceImpl struct {
	languageList []pkg.Language
}

type NotFoundError struct {
	ID string
}

func (n NotFoundError) Error() string {
	return "Language not found: " + n.ID
}

func NewStoreServiceImpl() *StoreServiceImpl {
	return &StoreServiceImpl{
		languageList: make([]pkg.Language, 0),
	}
}

func NewStoreService() store.StoreService {
	return NewStoreServiceImpl()
}

func NewStoreWithDefaultLangs() store.StoreService {
	storeService := NewStoreService()
	for _, option := range cpp.GetAllOptions() {
		storeService.Register(option)
	}
	for _, option := range python3.GetAllOptions() {
		storeService.Register(option)
	}
	return storeService
}

func (ss *StoreServiceImpl) Get(id string) (pkg.Language, error) {
	for i := range ss.languageList {
		if ss.languageList[i].ID() == id {
			return ss.languageList[i], nil
		}
	}
	return nil, NotFoundError{id}
}

func (ss *StoreServiceImpl) Register(l pkg.Language) {
	ss.languageList = append(ss.languageList, l)
}

func (ss *StoreServiceImpl) List() []pkg.Language {
	return ss.languageList
}

func (ss *StoreServiceImpl) Contains(id string) bool {
	_, err := ss.Get(id)
	return err == nil
}
