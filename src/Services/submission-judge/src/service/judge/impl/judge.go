package impl

import (
	"context"

	isolateservice "github.com/bibimoni/Online-judge/submission-judge/src/service/isolate"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/judge"
	poolservice "github.com/bibimoni/Online-judge/submission-judge/src/service/pool"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/problem"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/store"
)

type JudgeServiceImpl struct {
	pService *poolservice.PoolService
}

func NewJudgeServiceImpl(pService *poolservice.PoolService) *JudgeServiceImpl {
	return &JudgeServiceImpl{
		pService: pService,
	}
}

func NewJudgeService(pService *poolservice.PoolService) judge.JudgeService {
	return NewJudgeServiceImpl(pService)
}

func (js *JudgeServiceImpl) Judge(ctx context.Context, req *isolateservice.SubmissionRequest, problemInfo *problem.ProblemServiceGetOutput) error {
	lang, err := store.DefaultStore.Get((*req).LanguageId)
	if err != nil {
		return err
	}

	isolate, err := (*js.pService).Get()
	if err != nil {
		return err
	}

	isolate.Logger.Info().Msgf("Assigned to submission with id: %s", (*req).SubmissionId)
	err = lang.Judge(isolate, req)
	if err != nil {
		return err
	}
	return nil
}
