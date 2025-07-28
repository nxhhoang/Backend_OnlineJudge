package impl

import (
	"os/exec"
	"path/filepath"
	"strconv"

	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
)

// IsolateRoot is the root directory structure isolate is using.
var IsolateRoot = "/var/local/lib/isolate/"

type IsolateServiceImpl struct{}

func NewIsolateServiceImpl() *IsolateServiceImpl {
	return &IsolateServiceImpl{}
}

func (ir *IsolateServiceImpl) NewIsolate(id int) (*domain.Isolate, error) {
	res := &domain.Isolate{
		ID: id,
	}
	if res.Logger == nil {
		log, err := config.NewIsolateLogger(id)
		if err != nil {
			return nil, err
		}
		res.Logger = log
	}
	return res, nil
}

func (ir *IsolateServiceImpl) Cleanup(i *domain.Isolate) error {
	cmd := []string{"isolate", "--cg", "-b", strconv.Itoa(i.ID), "--cleanup"}
	i.Logger.Info().Msgf("Cleaning up... Running: %s", cmd)
	i.Inited = false

	return exec.Command(cmd[0], cmd[1:]...).Run()
}

func (ir *IsolateServiceImpl) Init(i *domain.Isolate) error {
	if err := ir.Cleanup(i); err != nil {
		return err
	}

	cmd := []string{"isolate", "--cg", "-b", strconv.Itoa(i.ID), "--init"}
	i.Logger.Info().Msgf("Creating isolate... Running: %s", cmd)
	i.Inited = true
	i.BoxDir = filepath.Join(IsolateRoot, strconv.Itoa(i.ID))
	return exec.Command(cmd[0], cmd[1:]...).Run()
}

func (ir *IsolateServiceImpl) Judge(i *domain.Isolate, rc *domain.RunConfig) {

}

func (ir *IsolateServiceImpl) RunBinary(i *domain.Isolate, rc *domain.RunConfig) {

}
