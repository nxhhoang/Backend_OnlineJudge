package impl

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"

	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	repository "github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/evaluation"
	"github.com/bibimoni/Online-judge/submission-judge/src/pkg"
	isolateservice "github.com/bibimoni/Online-judge/submission-judge/src/service/isolate"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/isolate/utils"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/judge"
	judgeutils "github.com/bibimoni/Online-judge/submission-judge/src/service/judge/utils"
	poolservice "github.com/bibimoni/Online-judge/submission-judge/src/service/pool"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/problem"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/store"
)

type JudgeServiceImpl struct {
	pService       *poolservice.PoolService
	problemService problem.ProblemService
	evalRepo       repository.EvaluationRepository
}

var CompilationError = errors.New("Compile error!")

func NewJudgeServiceImpl(pService *poolservice.PoolService, problemService problem.ProblemService, evalRepo repository.EvaluationRepository) *JudgeServiceImpl {
	return &JudgeServiceImpl{
		pService:       pService,
		problemService: problemService,
		evalRepo:       evalRepo,
	}
}

func NewJudgeService(pService *poolservice.PoolService, problemService problem.ProblemService, evalRepo repository.EvaluationRepository) judge.JudgeService {
	return NewJudgeServiceImpl(pService, problemService, evalRepo)
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
	isolate.Logger.Debug().Msgf("Took out a isolate, number of isolate remains in the pool is: %d", (*js.pService).Len())

	// Create a new context since judging has nothing to do with http request
	bgCtx := context.Background()
	go js.JudgeStart(bgCtx, isolate, lang, req, problemInfo)
	return nil
}

func (js *JudgeServiceImpl) JudgeStart(ctx context.Context, i *domain.Isolate, lang pkg.Language, req *isolateservice.SubmissionRequest, problemInfo *problem.ProblemServiceGetOutput) error {
	// Prepare all the nessessary files
	err := js.Prep(ctx, i, lang, req, problemInfo)
	if err != nil {
		i.Logger.Debug().Msgf("Error: %v", err)
		i.Logger.Debug().Msgf("Judgement failed or CompilationError, return the isolate, number of isolate in the pool is: %d", (*js.pService).Len())
		return err
	}
	i.Logger.Debug().Msgf("Judgement success!, keep using the isolate, number of isolate in the pool is: %d", (*js.pService).Len())

	return nil
}

// This function will help copy/create the nessessary files into the isolate working directory
func (js *JudgeServiceImpl) Prep(ctx context.Context, i *domain.Isolate, lang pkg.Language, req *isolateservice.SubmissionRequest, problemInfo *problem.ProblemServiceGetOutput) error {
	i.Logger.Info().Msgf("Assigned to submission with id: %s", (*req).SubmissionId)

	var (
		errBuf bytes.Buffer
		err    error
		vert   *judge.RunVerdict
	)

	// always remember to return the isolate instance
	defer func(e *error) {
		judgeutils.ReturnIsolateIfFail(js.pService, i, *e)
	}(&err)

	_, err = utils.CreateSubmissionSourceFile(i, req.Sourcecode, req.SubmissionId, lang.DefaultFileName())
	if err != nil {
		return err
	}

	i.Logger.Info().Msgf("Created source file inside the isolate working directory")

	if lang.NeedCompile() {

		lang.Compile(i, req, &errBuf)
		vert, err = js.CheckRunStatus(i, req.SubmissionId)

		if err != nil {
			nerr := js.evalRepo.UpdateVerdict(ctx, req.EvalId, domain.JUDGEMENT_FAILED)
			if nerr != nil {
				i.Logger.Panic().Msgf("Database error, can't update verdict: %v", nerr)
			}
			return err
		}
		switch vert.Status {
		case "RE", "SG", "TO", "XX":
			nerr := js.evalRepo.UpdateVerdict(ctx, req.EvalId, domain.COMPILATION_ERROR)
			if nerr != nil {
				i.Logger.Error().Msgf("Database error, can't update verdict: %v", nerr)
			}
			err = judge.CompilationError
		default:
			if vert.Status != "" || vert.ExitCode != 0 {
				nerr := js.evalRepo.UpdateVerdict(ctx, req.EvalId, domain.JUDGEMENT_FAILED)
				if nerr != nil {
					i.Logger.Error().Msgf("Database error, can't update verdict: %v", nerr)
				}
				err = judge.JugdgementFailed
			}
		}
		if err != nil {
			return err
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

	return judgeutils.ParseMetaFile(meta)
}

func (js *JudgeServiceImpl) JudgeICPC() error {
	return nil
}
