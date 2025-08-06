package impl

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	"github.com/bibimoni/Online-judge/submission-judge/src/pkg"
	isolateservice "github.com/bibimoni/Online-judge/submission-judge/src/service/isolate"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/isolate/utils"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/judge"
	poolservice "github.com/bibimoni/Online-judge/submission-judge/src/service/pool"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/problem"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/store"
)

type JudgeServiceImpl struct {
	pService       *poolservice.PoolService
	problemService problem.ProblemService
}

var CompilationError = errors.New("Compile error!")

func NewJudgeServiceImpl(pService *poolservice.PoolService, problemService problem.ProblemService) *JudgeServiceImpl {
	return &JudgeServiceImpl{
		pService:       pService,
		problemService: problemService,
	}
}

func NewJudgeService(pService *poolservice.PoolService, problemService problem.ProblemService) judge.JudgeService {
	return NewJudgeServiceImpl(pService, problemService)
}

// This will be the final wrapper to double check condition, at the end of the function
// The real judge function will be called, and it will be asynchonous
func (js *JudgeServiceImpl) Judge(ctx context.Context, req *isolateservice.SubmissionRequest, problemInfo *problem.ProblemServiceGetOutput) error {
	lang, err := store.DefaultStore.Get((*req).LanguageId)
	if err != nil {
		return err
	}

	isolate, err := (*js.pService).Get()
	if err != nil {
		return err
	}

	go js.JudgeStart(ctx, isolate, lang, req, problemInfo)
	return nil
}

func (js *JudgeServiceImpl) JudgeStart(ctx context.Context, i *domain.Isolate, lang pkg.Language, req *isolateservice.SubmissionRequest, problemInfo *problem.ProblemServiceGetOutput) error {
	// Prepare all the nessessary files
	err := js.Prep(ctx, i, lang, req, problemInfo)
	return err
}

// This function will help copy/create the nessessary files into the isolate working directory
func (js *JudgeServiceImpl) Prep(ctx context.Context, i *domain.Isolate, lang pkg.Language, req *isolateservice.SubmissionRequest, problemInfo *problem.ProblemServiceGetOutput) error {
	i.Logger.Info().Msgf("Assigned to submission with id: %s", (*req).SubmissionId)

	var errBuf bytes.Buffer

	i.Logger.Debug().Msgf("Output: %s", errBuf.String())

	_, err := utils.CreateSubmissionSourceFile(i, req.Sourcecode, req.SubmissionId, lang.DefaultFileName())
	if err != nil {
		return err
	}

	i.Logger.Info().Msgf("Created source file inside the isolate working directory")
	err = lang.Compile(i, req, &errBuf)
	if err != nil {
		return err
	}

	if lang.NeedCompile() {
		vert, err := js.CheckRunStatus(i, req.SubmissionId)
		// TODO: update evaluation result when first compile the program
		if err != nil {

		}
	}

	// Prepare the checker file
	checkerLocation, err := js.problemService.GetCheckerAddr(req.ProblemId)
	err = utils.CopyChecker(i, (*req).SubmissionId, checkerLocation)
	if err != nil {
		return err
	}
	return nil
}

func (js *JudgeServiceImpl) CheckRunStatus(i *domain.Isolate, submissionId string) (*judge.RunVerdict, error) {
	metaAddr, err := utils.GetMetaFilePath(i, submissionId)
	if err != nil {
		return nil, err
	}

	meta, err := os.ReadFile(metaAddr)
	if err != nil {
		return nil, fmt.Errorf("Can't open meta file: %v", err)
	}

	var m judge.RunVerdict
	if err := json.Unmarshal(meta, &m); err != nil {
		return nil, fmt.Errorf("Can't parse meta file to json")
	}

	return &m, nil
}

func (js *JudgeServiceImpl) JudgeICPC() error {
	return nil
}
