package transportgetsubmission

import (
	"github.com/bibimoni/Online-judge/submission-judge/src/common"
	appctx "github.com/bibimoni/Online-judge/submission-judge/src/components"
	"github.com/bibimoni/Online-judge/submission-judge/src/controller"
	ei "github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/evaluation/impl"
	sci "github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/sourcecode/impl"
	si "github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/submission/impl"
	checkerimpl "github.com/bibimoni/Online-judge/submission-judge/src/service/checker/impl"
	ji "github.com/bibimoni/Online-judge/submission-judge/src/service/judge/impl"
	pi "github.com/bibimoni/Online-judge/submission-judge/src/service/problem/impl"
	"github.com/bibimoni/Online-judge/submission-judge/src/usecase/submission/interactor"
	"github.com/gin-gonic/gin"

	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
	"github.com/bibimoni/Online-judge/submission-judge/src/usecase/submission"
)

func HandleGetSubmissionRequest(appContext appctx.AppContext) gin.HandlerFunc {
	log := config.GetLogger()

	db := appContext.GetMainDbConnection()
	submissionRepo := si.NewSubmissionRepository(db)
	sourcecodeRepo := sci.NewSourcecodeRepository(db)
	problemSvc, err := pi.NewProblemService()
	evalRepo := ei.NewEvaluationRepository(db)
	checker := checkerimpl.NewCheckerService()
	judgeSvc := ji.NewJudgeServiceImpl(appContext.GetPool(), problemSvc, evalRepo, checker)
	if err != nil {
		log.Error().Msgf("Can't initialize submit request, got error : %v", err)
		return nil
	}
	submissionInteractor := interactor.NewSubmissionInteractor(submissionRepo, sourcecodeRepo, problemSvc, judgeSvc, evalRepo)

	return common.InvokeUseCase(
		toGetSubmissionType,
		submissionInteractor.GetSubmission,
		helper.WriteCreatedOutput,
	)
}

func toGetSubmissionType(c *gin.Context) (*usecase.GetSubmissionInput, error) {
	sid := c.Param("submission_id")
	log := config.GetLogger()
	log.Debug().Msgf("get submission id from request: %s", sid)
	return &usecase.GetSubmissionInput{SubmissionId: sid}, nil
}
