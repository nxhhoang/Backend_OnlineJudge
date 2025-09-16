package router

import (
	appctx "github.com/bibimoni/Online-judge/submission-judge/src/components"
	transportgetsubmission "github.com/bibimoni/Online-judge/submission-judge/src/controller/getsubmission"
	"github.com/bibimoni/Online-judge/submission-judge/src/controller/submitsubmission"
	"github.com/bibimoni/Online-judge/submission-judge/src/controller/websocketsubmission"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(group *gin.RouterGroup, appContext appctx.AppContext) {
	submission := group.Group("/submission")
	submission.POST("/submit", transportsubmitsubmission.HandleSubmitSubmissionRequest(appContext))
	submission.GET("/view/:submission_id", transportgetsubmission.HandleGetSubmissionRequest(appContext))
	submission.GET("/ws", websocketsubmission.HandleSubmissionWSRequest(appContext))
	submission.GET("/problem/view/:problem_id", transportgetsubmission.HandleGetProblemSubmissionRequest(appContext))
}
