package transport

import (
	"github.com/bibimoni/Online-judge/submission-judge/src/common"
	appctx "github.com/bibimoni/Online-judge/submission-judge/src/components"
	"github.com/bibimoni/Online-judge/submission-judge/src/controller"
	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	scr "github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/sourcecode"
	sr "github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/submission"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/problem"
	"github.com/bibimoni/Online-judge/submission-judge/src/usecase/submitsubmission/interactor"
	"github.com/gin-gonic/gin"

	"fmt"

	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
	"github.com/bibimoni/Online-judge/submission-judge/src/usecase/submitsubmission"
)

func HandleSubmitSubmissionRequest(appContext appctx.AppContext) gin.HandlerFunc {
	log := config.GetLogger()

	db := appContext.GetMainDbConnection()
	submissionRepo := sr.NewSubmissionRepository(db)
	sourcecodeRepo := scr.NewSourcecodeRepository(db)
	problemSvc, err := problem.NewProblemService()
	if err != nil {
		log.Error().Msgf("Can't initialize submit request, got error : %v", err)
		return nil
	}
	submissionInteractor := interactor.NewSubmissionInteractor(submissionRepo, sourcecodeRepo, problemSvc, appContext.GetPool())

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
