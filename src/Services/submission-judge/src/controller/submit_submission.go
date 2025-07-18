package transport

import (
	"github.com/bibimoni/Online-judge/submission-judge/src/common"
	appctx "github.com/bibimoni/Online-judge/submission-judge/src/components"
	"github.com/bibimoni/Online-judge/submission-judge/src/controller/helper"
	"github.com/bibimoni/Online-judge/submission-judge/src/domain/repository"
	"github.com/bibimoni/Online-judge/submission-judge/src/usecase/interactor"
	"github.com/gin-gonic/gin"
)

func HandleSubmitSubmissionRequest(appContext appctx.AppContext) gin.HandlerFunc {
	db := appContext.GetMainDbConnection()

	submissionRepo := repository.NewSubmissionRepository(db)
	submissionInteractor := interactor.NewSubmissionInteractor(submissionRepo)

	return common.InvokeUseCase(
		helper.ToSubmitSubmissionType,
		submissionInteractor.SubmitSubmission,
		helper.WriteCreatedOutput,
	)
}
