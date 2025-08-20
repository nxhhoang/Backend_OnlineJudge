package transportsubmitsubmission

import (
	"github.com/bibimoni/Online-judge/submission-judge/src/common"
	appctx "github.com/bibimoni/Online-judge/submission-judge/src/components"
	"github.com/bibimoni/Online-judge/submission-judge/src/controller"
	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	ei "github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/evaluation/impl"
	"github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/redissubmission/impl"
	sci "github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/sourcecode/impl"
	si "github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/submission/impl"
	checkerimpl "github.com/bibimoni/Online-judge/submission-judge/src/service/checker/impl"
	ji "github.com/bibimoni/Online-judge/submission-judge/src/service/judge/impl"
	pi "github.com/bibimoni/Online-judge/submission-judge/src/service/problem/impl"
	"github.com/bibimoni/Online-judge/submission-judge/src/usecase/submission/interactor"
	"github.com/gin-gonic/gin"

	"fmt"

	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
	"github.com/bibimoni/Online-judge/submission-judge/src/usecase/submission"
)

func HandleSubmitSubmissionRequest(appContext appctx.AppContext) gin.HandlerFunc {
	log := config.GetLogger()

	db := appContext.GetMainDbConnection()
	submissionRepo := si.NewSubmissionRepository(db)
	sourcecodeRepo := sci.NewSourcecodeRepository(db)
	problemSvc, err := pi.NewProblemService()
	evalRepo := ei.NewEvaluationRepository(db)
	checker := checkerimpl.NewCheckerService()
	redis := impl.NewRedisSubmissionRepository(appContext.GetRedis())
	judgeSvc := ji.NewJudgeServiceImpl(appContext.GetPool(), problemSvc, evalRepo, checker, redis, submissionRepo, sourcecodeRepo)
	if err != nil {
		log.Error().Msgf("Can't initialize submit request, got error : %v", err)
		return nil
	}
	submissionInteractor := interactor.NewSubmissionInteractor(submissionRepo, sourcecodeRepo, problemSvc, judgeSvc, evalRepo)

	return common.InvokeUseCase(
		toSubmitSubmissionType,
		submissionInteractor.SubmitSubmission,
		helper.WriteCreatedOutput,
	)
}

func toSubmitSubmissionType(c *gin.Context) (*usecase.SubmitSubmissionInput, error) {
	log := config.GetLogger()
	var input usecase.SubmitSubmissionInput
	if err := c.BindJSON(&input); err != nil {
		log.Error().Msgf("%s", err.Error())
		return nil, fmt.Errorf("Invalid Request Body")
	}

	// Guard submission type, i think all the validation should happen here
	// as long as it doesn't require any service / repository
	if input.SubmissionType != domain.SubmissionType(domain.ICPC) {
		return nil, fmt.Errorf("Sorry, we currently don't support thi type of submission: %s", input.SubmissionType)
	}

	return &input, nil
}
