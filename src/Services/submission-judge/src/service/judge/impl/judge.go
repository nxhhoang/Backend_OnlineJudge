package impl

import (
	"bytes"
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	evalrepository "github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/evaluation"
	redisrepository "github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/redissubmission"
	screpository "github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/sourcecode"
	subrepository "github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/submission"
	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
	"github.com/bibimoni/Online-judge/submission-judge/src/pkg"
	"github.com/bibimoni/Online-judge/submission-judge/src/pkg/memory"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/checker"
	isolateservice "github.com/bibimoni/Online-judge/submission-judge/src/service/isolate"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/isolate/utils"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/judge"
	judgeutils "github.com/bibimoni/Online-judge/submission-judge/src/service/judge/utils"
	poolservice "github.com/bibimoni/Online-judge/submission-judge/src/service/pool"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/problem"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/store"
	isubmission_utils "github.com/bibimoni/Online-judge/submission-judge/src/usecase/submission/utils"
	usecase "github.com/bibimoni/Online-judge/submission-judge/src/usecase/wssubmission"
)

type JudgeServiceImpl struct {
	pService       *poolservice.PoolService
	problemService problem.ProblemService
	evalRepo       evalrepository.EvaluationRepository
	checkerService checker.CheckerService
	redisRepo      redisrepository.RedisSubmissionRepository
	submissionRepo subrepository.SubmissionRepository
	sourcecodeRepo screpository.SourcecodeRepository
}

var CompilationError = errors.New("Compile error!")
var JudgementFailedMessage = "failed to evaluate submission"
var AcceptedMessage = "Accepted"

func NewJudgeServiceImpl(
	pService *poolservice.PoolService,
	problemService problem.ProblemService,
	evalRepo evalrepository.EvaluationRepository,
	checkerS checker.CheckerService,
	redis redisrepository.RedisSubmissionRepository,
	submissionRepo subrepository.SubmissionRepository,
	sourcecodeRepo screpository.SourcecodeRepository,
) *JudgeServiceImpl {
	return &JudgeServiceImpl{
		pService:       pService,
		problemService: problemService,
		evalRepo:       evalRepo,
		checkerService: checkerS,
		redisRepo:      redis,
		submissionRepo: submissionRepo,
		sourcecodeRepo: sourcecodeRepo,
	}
}

func NewJudgeService(
	pService *poolservice.PoolService,
	problemService problem.ProblemService,
	evalRepo evalrepository.EvaluationRepository,
	checkerS checker.CheckerService,
	redis redisrepository.RedisSubmissionRepository,
	submissionRepo subrepository.SubmissionRepository,
	sourcecodeRepo screpository.SourcecodeRepository,
) judge.JudgeService {
	return NewJudgeServiceImpl(pService, problemService, evalRepo, checkerS, redis, submissionRepo, sourcecodeRepo)
}

// This will be the final wrapper to double check condition, at the end of the function
// The real judge function will be called, and it will be asynchonous
func (js *JudgeServiceImpl) Judge(ctx context.Context, req *isolateservice.SubmissionRequest, problemInfo *problem.ProblemServiceGetOutput) error {
	lang, err := store.DefaultStore.Get((*req).LanguageId)
	if err != nil {
		return err
	}
	// Create a new context since judging has nothing to do with http request
	bgCtx := context.Background()
	go js.JudgeStart(bgCtx, lang, req, problemInfo)
	return nil
}

func (js *JudgeServiceImpl) JudgeStart(ctx context.Context, lang pkg.Language, req *isolateservice.SubmissionRequest, problemInfo *problem.ProblemServiceGetOutput) error {
	// update PENDING to Websocket
	err := js.updateWS(ctx, req.EvalId)
	if err != nil {
		config.GetLogger().Panic().Msgf("Redis stopped working: %v", err)
		return err
	}

	// This should be here, inside judgeStart
	i, err := (*js.pService).Get()
	if err != nil {
		// Very unlikely, happen when channel is closed
		err = js.updateFinal(
			ctx,
			req.EvalId,
			domain.JUDGEMENT_FAILED,
			0,
			0,
			0,
			0,
			JudgementFailedMessage,
		)
		if err != nil {
			i.Logger.Panic().Msgf("Database error, can't update verdict: %v", err)
		}
		return err
	}

	i.Logger.Debug().Msgf("Took out an isolate, number of isolate remains in the pool is: %d", (*js.pService).Len())
	// Prepare all the nessessary files
	err = js.Prep(ctx, i, lang, req, problemInfo)
	if err != nil {
		i.Logger.Debug().Msgf("Error: %v", err)
		i.Logger.Debug().Msgf("Judgement failed or CompilationError, return the isolate, number of isolate in the pool is: %d", (*js.pService).Len())
		return err
	}

	i.Logger.Debug().Msgf("Preparation success!, keep using the isolate, number of isolate in the pool is: %d", (*js.pService).Len())

	switch req.SubmissionType {
	case domain.SubmissionType(domain.ICPC):
		err = js.JudgeICPC(ctx, i, lang, req, problemInfo)
	default:
		i.Logger.Error().Msgf("Other submission type is not supported")
	}

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

	lang.Compile(i, req, &errBuf)
	vert, err = judgeutils.CheckRunStatus(i, req.SubmissionId)
	msg := judgeutils.GetCompileMessage(vert, errBuf.String())

	i.Logger.Info().Msgf("Compile message: %s", errBuf.String())
	if err != nil {
		err = js.updateFinal(
			ctx,
			req.EvalId,
			domain.JUDGEMENT_FAILED,
			vert.Time,
			vert.MaxRss,
			0,
			0,
			vert.Message,
		)
		if err != nil {
			i.Logger.Panic().Msgf("Database error, can't update verdict: %v", err)
		}
		return err
	}

	switch vert.Status {
	case "RE", "SG", "TO", "XX":
		err = js.updateFinal(
			ctx,
			req.EvalId,
			domain.COMPILATION_ERROR,
			vert.Time,
			vert.MaxRss,
			0,
			0,
			msg,
		)
		if err != nil {
			i.Logger.Error().Msgf("Database error, can't update verdict: %v", err)
		}
		err = judge.CompilationError
		i.Logger.Debug().Msgf("Compile error: %v", err)
	default:
		// This is just to detect if the program failed to compile via the information given by the meta file,
		// this is basically hardcoding and i might have to find a way to make this cleaner
		if vert.Status != "" || vert.ExitCode != 0 {
			err = js.updateFinal(
				ctx,
				req.EvalId,
				domain.JUDGEMENT_FAILED,
				vert.Time,
				vert.MaxRss,
				0,
				0,
				vert.Message,
			)
			if err != nil {
				i.Logger.Error().Msgf("Database error, can't update verdict: %v", err)
			}
			err = judge.JugdgementFailed
		}

		i.Logger.Debug().Msgf("Compile error: %v", err)
		if err != nil {
			return err
		}
	}

	// Prepare the checker file
	checkerLocation, e := js.problemService.GetCheckerAddr(req.ProblemId)
	if e != nil {
		err = js.updateFinal(
			ctx,
			req.EvalId,
			domain.JUDGEMENT_FAILED,
			vert.Time,
			vert.MaxRss,
			0,
			0,
			vert.Message,
		)
		if err != nil {
			i.Logger.Error().Msgf("Database error, can't update verdict: %v", err)
		}
		return e
	}
	err = utils.CopyChecker(i, (*req).SubmissionId, checkerLocation)
	if err != nil {
		return err
	}
	i.Logger.Info().Msgf("Checker file copied successfully!")

	return nil
}

func (js *JudgeServiceImpl) JudgeICPC(ctx context.Context, i *domain.Isolate, lang pkg.Language, req *isolateservice.SubmissionRequest, problemInfo *problem.ProblemServiceGetOutput) error {
	var (
		err error
	)
	defer func(e *error) {
		judgeutils.ReturnIsolateIfFail(js.pService, i, *e)
	}(&err)

	var (
		curCpuTime     float64       = 0
		curMemoryUsage memory.Memory = 0
	)

	for tc := 1; tc <= problemInfo.TestNum; tc += 1 {
		tcInputAddr, e := js.problemService.GetTestCaseAddr(req.ProblemId, problem.TestCaseType(problem.INPUT), tc)
		err = e
		if err != nil {
			i.Logger.Debug().Msgf("Error when fetching testcase: %v", err)
			// err = js.evalRepo.UpdateVerdict(ctx, req.EvalId, domain.JUDGEMENT_FAILED)
			err = js.updateFinal(
				ctx,
				req.EvalId,
				domain.JUDGEMENT_FAILED,
				curCpuTime,
				curMemoryUsage,
				tc-1,
				0,
				JudgementFailedMessage,
			)

			if err != nil {
				i.Logger.Panic().Msgf("Database error, can't update verdict: %v", err)
			}
			return err
		}

		outaddr := utils.GetSubmissionDir(i, req.SubmissionId) + "/output_" + strconv.Itoa(tc)
		i.Logger.Debug().Msgf("Output dir: %s", outaddr)
		fout, e := os.Create(outaddr)
		err = e
		if err != nil {
			i.Logger.Panic().Msgf("Error occured when trying to create new output file: %v", err)

		}

		fin, e := os.Open(tcInputAddr)
		err = e
		if err != nil {
			i.Logger.Panic().Msgf("Error occured when trying to read input file: %v", err)
		}

		rc := domain.RunConfig{
			TimeLimit:    time.Millisecond * time.Duration(problemInfo.TimeLimit),
			MemoryLimit:  memory.Memory(problemInfo.MemoryLimit),
			Meta:         true,
			Stdout:       fout,
			Stdin:        fin,
			MaxProcesses: 1,
		}

		i.Logger.Debug().Msgf("Start to run the code, config is: %v", rc)
		lang.Run(i, &rc, req)
		fin.Close()

		vert, e := judgeutils.CheckRunStatus(i, req.SubmissionId)
		err = e

		i.Logger.Debug().Msgf("Run Status from MetaFile: %v", vert)
		if err != nil {

			err = js.updateFinal(
				ctx,
				req.EvalId,
				domain.JUDGEMENT_FAILED,
				curCpuTime,
				curMemoryUsage,
				tc-1,
				0,
				JudgementFailedMessage,
			)

			if err != nil {
				i.Logger.Panic().Msgf("Database error, can't update verdict: %v", err)
			}
			return err
		}

		curCpuTime = max(curCpuTime, vert.Time)
		curMemoryUsage = max(curMemoryUsage, vert.MaxRss)

		tcAnsAddtr, e := js.problemService.GetTestCaseAddr(req.ProblemId, problem.TestCaseType(problem.OUTPUT), tc)
		err = e
		if err != nil {
			err = js.updateFinal(
				ctx,
				req.EvalId,
				domain.JUDGEMENT_FAILED,
				curCpuTime,
				curMemoryUsage,
				tc-1,
				0,
				vert.Message,
			)
			return err
		}
		checkerLocation, e := js.problemService.GetCheckerAddr(req.ProblemId)
		err = e
		if err != nil {
			err = js.updateFinal(
				ctx,
				req.EvalId,
				domain.JUDGEMENT_FAILED,
				curCpuTime,
				curMemoryUsage,
				tc-1,
				0,
				vert.Message,
			)
			return err
		}

		cvert, msg, e := js.checkVerdict(vert, checkerLocation, tcInputAddr, outaddr, tcAnsAddtr)
		err = e
		if err != nil {
			err = js.updateFinal(
				ctx,
				req.EvalId,
				domain.JUDGEMENT_FAILED,
				curCpuTime,
				curMemoryUsage,
				tc-1,
				0,
				vert.Message,
			)
			return err
		}

		// err = js.evalRepo.UpdateCase(ctx, req.EvalId, cvert, vert.Time, vert.MaxRss, msg, 1)
		curSuccess := tc
		if cvert != domain.ACCEPTED {
			curSuccess -= 1
		}
		err = js.updateCase(
			ctx,
			req.EvalId,
			cvert,
			vert.Time,
			vert.MaxRss,
			msg,
			1,
			curCpuTime,
			curMemoryUsage,
			curSuccess,
		)

		if err != nil {
			i.Logger.Panic().Msgf("Database error, can't update verdict: %v", err)
		}
		if cvert != domain.ACCEPTED {
			err = js.updateFinal(
				ctx,
				req.EvalId,
				cvert,
				curCpuTime,
				curMemoryUsage,
				tc-1,
				0,
				msg,
			)
			if err != nil {
				i.Logger.Panic().Msgf("Database error, can't update verdict: %v", err)
			}
			(*js.pService).Put(i)
			return nil
		}
	}

	err = js.updateFinal(
		ctx,
		req.EvalId,
		domain.ACCEPTED,
		curCpuTime,
		curMemoryUsage,
		problemInfo.TestNum,
		1,
		AcceptedMessage,
	)
	if err != nil {
		i.Logger.Panic().Msgf("Database error, can't update verdict: %v", err)
	}
	(*js.pService).Put(i)
	i.Logger.Debug().Msgf("Judgement success!, return the isolate, number of isolate in the pool is: %d", (*js.pService).Len())

	return nil
}

func (js *JudgeServiceImpl) checkVerdict(vert *judge.RunVerdict, checkerAddr, inputAddr, outputAddr, answerAddr string) (domain.Verdict, string, error) {
	config.GetLogger().Debug().Msgf("Status is: %s", vert.Status)
	switch vert.Status {
	case "TO":
		return domain.TIME_LIMIT_EXCEEDED, vert.Message, nil
	case "RE", "SG":
		return domain.RUNTIME_ERROR, vert.Message, nil
	case "XX":
		return domain.JUDGEMENT_FAILED, vert.Message, nil
	}

	// this return the message and the exit code, which must be use later
	// TODO: Do something with exit code and checker message
	cvert, _, msg, err := js.checkerService.RunChecker(checkerAddr, inputAddr, outputAddr, answerAddr)
	if err != nil {
		return "", "", err
	}
	return cvert, msg, nil
}

func (js *JudgeServiceImpl) updateFinal(
	ctx context.Context,
	evalId string,
	verdict domain.Verdict,
	cpuTime float64,
	memoryUsage memory.Memory,
	nsucess int,
	points int,
	message string,
) error {
	err := js.evalRepo.UpdateFinal(
		ctx,
		evalId,
		verdict,
		cpuTime,
		memoryUsage,
		nsucess,
		points,
		message,
	)
	if err != nil {
		return err
	}
	return js.updateWS(ctx, evalId)
}

func (js *JudgeServiceImpl) updateCase(
	ctx context.Context,
	evalId string,
	verdictCase domain.Verdict,
	cpuTimeCase float64,
	memoryUsageCase memory.Memory,
	outputCase string,
	pointsCase int,
	cpuTime float64,
	memoryUsage memory.Memory,
	nsucess int,
) error {
	err := js.evalRepo.UpdateCase(
		ctx,
		evalId,
		verdictCase,
		cpuTimeCase,
		memoryUsageCase,
		outputCase,
		pointsCase,
		cpuTime,
		memoryUsage,
		nsucess,
	)
	if err != nil {
		return err
	}

	return js.updateWS(ctx, evalId)
}

func (js *JudgeServiceImpl) updateWS(ctx context.Context, evalId string) error {
	eval, err := js.evalRepo.GetEval(ctx, evalId)
	if err != nil {
		return err
	}

	sub, err := isubmission_utils.GetSubmission(
		ctx,
		js.evalRepo,
		js.sourcecodeRepo,
		js.submissionRepo,
		eval.SubmissionId.Hex(),
	)

	if err != nil {
		return err
	}

	wsUpdate := usecase.WSSubmissionResponse{
		Username:        sub.Username,
		SubmissionId:    eval.SubmissionId.Hex(),
		ProblemId:       sub.ProblemId,
		Timestamp:       sub.Timestamp,
		Language:        sub.Language,
		Verdict:         sub.Verdict,
		VerdictCase:     sub.VerdictCase,
		CpuTime:         sub.CpuTime,
		CpuTimeCase:     sub.CpuTimeCase,
		MemoryUsage:     sub.MemoryUsage,
		MemoryUsageCase: sub.MemoryUsageCase,
		NSuccess:        sub.NSuccess,
		Outputs:         sub.Outputs,
		Points:          sub.Points,
		PointsCase:      sub.PointsCase,
		Message:         sub.Message,
		EvalStatus:      sub.EvalStatus,
	}

	return js.redisRepo.PulishSubmission(ctx, wsUpdate)
}
