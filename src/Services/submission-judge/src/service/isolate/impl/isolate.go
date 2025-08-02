package impl

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
	"github.com/bibimoni/Online-judge/submission-judge/src/pkg"
	"github.com/bibimoni/Online-judge/submission-judge/src/pkg/memory"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/problem"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/problem/impl"
)

// IsolateRoot is the root directory structure isolate is using.
var IsolateRoot = "/var/local/lib/isolate/"
var IsolateInputDirName = "in"
var IsolateWorkingDirName = "app"

// isolate -b 0 --time=2 --mem=256000 --dir=/in=/problems_dir/445985/tests/input:rw --dir=/app/=/var/local/lib/isolate/0/box:rw -i /in/01 -o /app/01 --run -- /app/main

func GetIsolateDir(i *domain.Isolate) string {
	return IsolateRoot + strconv.Itoa(i.ID) + "/box"
}

func GetIsolateInputDir(submissionId string) string {
	return submissionId + "/" + IsolateInputDirName
}

func GetIsolateWorkingDir(submissionId string) string {
	return submissionId + "/" + IsolateWorkingDirName
}

type IsolateServiceImpl struct {
	problemService problem.ProblemService
}

func NewIsolateServiceImpl() (*IsolateServiceImpl, error) {
	ps, err := problem.NewProblemService()
	if err != nil {
		return nil, err
	}
	return &IsolateServiceImpl{
		problemService: ps,
	}, nil
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
		(*res).Logger = log
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

	log := config.GetLogger()
	log.Info().Msgf("Inited isolate service with id: %d", i.ID)
	return exec.Command(cmd[0], cmd[1:]...).Run()
}

// This method added input directory to the run config, this also verify if the directory exists
func (ir *IsolateServiceImpl) addInputMappedDir(rc *domain.RunConfig, problemId string, submissionId string) error {
	inputDirAddr, err := ir.problemService.GetTestCaseDirAddr(problemId, impl.INPUT)
	if err != nil {
		return err
	}

	rc.DirectoryMaps = append(rc.DirectoryMaps, domain.DirectoryMap{
		Inside:  inputDirAddr,
		Outside: GetIsolateInputDir(submissionId),
		Options: []domain.DirectoryMapOption{domain.AllowReadWrite},
	})
	return nil
}

func buildArgs(i *domain.Isolate, rc *domain.RunConfig) ([]string, error) {
	args := []string{"isolate", "-b", strconv.Itoa(i.ID)}

	if rc.MaxProcesses > 0 {
		args = append(args, fmt.Sprintf("--processes=%d", rc.MaxProcesses))
	} else {
		args = append(args, "--processes=100")
	}

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

	if rc.TimeLimit > 0 {
		ms := rc.TimeLimit / time.Millisecond
		args = append(args, fmt.Sprintf("--time=%d.%d", ms/1000, ms%1000))
		args = append(args, fmt.Sprintf("--wall-time=%d.%d", ms/1000, ms%1000))
	}
	if rc.MemoryLimit > 0 {
		args = append(args, fmt.Sprintf("--cg-mem=%d", int(rc.MemoryLimit/memory.KiB)))
	}
	return args, nil
}

func (ir *IsolateServiceImpl) Run(i *domain.Isolate, rc *domain.RunConfig, req *pkg.SubmissionRequest) error {

	return nil
}

func (ir *IsolateServiceImpl) Judge(i *domain.Isolate, rc *domain.RunConfig, req *pkg.SubmissionRequest) error {
	return nil
}
