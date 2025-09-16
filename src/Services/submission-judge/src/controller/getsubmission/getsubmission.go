package transportgetsubmission

import (
	"github.com/bibimoni/Online-judge/submission-judge/src/common"
	appctx "github.com/bibimoni/Online-judge/submission-judge/src/components"
	helper "github.com/bibimoni/Online-judge/submission-judge/src/controller"
	controller_utils "github.com/bibimoni/Online-judge/submission-judge/src/controller/utils"
	"github.com/gin-gonic/gin"

	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
	usecase "github.com/bibimoni/Online-judge/submission-judge/src/usecase/submission"
)

func HandleGetSubmissionRequest(appContext appctx.AppContext) gin.HandlerFunc {
	submissionInteractor, err := controller_utils.InitInteractor(appContext)
	if err != nil {
		return nil
	}

	return common.InvokeUseCase(
		toGetSubmissionType,
		submissionInteractor.GetSubmission,
		helper.WriteSuccessOutput,
	)
}

func HandleGetProblemSubmissionRequest(appContext appctx.AppContext) gin.HandlerFunc {
	submissionInteractor, err := controller_utils.InitInteractor(appContext)
	if err != nil {
		return nil
	}
	return common.InvokeUseCase(
		toGetProblemSubmissionType,
		submissionInteractor.GetProblemSubmission,
		helper.WriteSuccessOutput,
	)
}

func toGetProblemSubmissionType(c *gin.Context) (*usecase.GetProblemSubmissionInput, error) {
	pid := c.Param("problem_id")
	log := config.GetLogger()
	log.Debug().Msgf("get problem id from request: %s", pid)
	return &usecase.GetProblemSubmissionInput{
		ProblemId: pid,
	}, nil
}

func toGetSubmissionType(c *gin.Context) (*usecase.GetSubmissionInput, error) {
	sid := c.Param("submission_id")
	log := config.GetLogger()
	log.Debug().Msgf("get submission id from request: %s", sid)
	return &usecase.GetSubmissionInput{SubmissionId: sid}, nil
}
