package impl

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strconv"

	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
)

// IsolateRoot is the root directory structure isolate is using.
var IsolateRoot = "/var/local/lib/isolate/"

// isolate -b 0 --time=2 --mem=256000 --dir=/in=/problems_dir/445985/tests/input:rw --dir=/app/=/var/local/lib/isolate/0/box:rw -i /in/01 -o /app/01 --run -- /app/main

type IsolateServiceImpl struct{}

func NewIsolateServiceImpl() *IsolateServiceImpl {
	return &IsolateServiceImpl{}
}

type Task struct {
	SubmissionId   string
	Username       string
	Sourcecode     string
	SubmissionType string
	ProblemId      string
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

func (ir *IsolateServiceImpl) buildArgs(i *domain.Isolate, rc *domain.RunConfig) ([]string, error) {
	args := []string{"isolate", "-b", strconv.Itoa(i.ID)}
	args = append(args, "--processes=1")
	if rc.InheritEnv {
		args = append(args, "--full-env")
	}
	for ind := range rc.Env {
		args = append(args, fmt.Sprintf("--env=%s", rc.Env[ind]))
	}
	for _, rule := range rc.DirectoryMaps {
		arg := fmt.Sprintf("--dir=%s=%s", rule.Inside, rule.Outside)
		for _, opt := range rule.Options {
			arg += ":" + string(opt)
		}
		args = append(args, arg)
	}

	return args, nil
}

func (ir *IsolateServiceImpl) Run(i *domain.Isolate, rc *domain.RunConfig, inFileLocation string, outFileLocation string) {

}
