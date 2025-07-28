package transport

import (
	"github.com/bibimoni/Online-judge/submission-judge/src/common"
	appctx "github.com/bibimoni/Online-judge/submission-judge/src/components"
	"github.com/bibimoni/Online-judge/submission-judge/src/controller"
	scr "github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/sourcecode"
	sr "github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/submission"
	"github.com/bibimoni/Online-judge/submission-judge/src/usecase/submitsubmission/interactor"
	"github.com/gin-gonic/gin"

	"fmt"

	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
	"github.com/bibimoni/Online-judge/submission-judge/src/usecase/submitsubmission"
)

func HandleSubmitSubmissionRequest(appContext appctx.AppContext) gin.HandlerFunc {
	db := appContext.GetMainDbConnection()

	submissionRepo := sr.NewSubmissionRepository(db)
	sourcecodeRepo := scr.NewSourcecodeRepository(db)
	submissionInteractor := interactor.NewSubmissionInteractor(submissionRepo, sourcecodeRepo)

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
	return &input, nil
}
