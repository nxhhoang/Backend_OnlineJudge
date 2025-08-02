package router

import (
	appctx "github.com/bibimoni/Online-judge/submission-judge/src/components"
	"github.com/bibimoni/Online-judge/submission-judge/src/controller/submitsubmission"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(group *gin.RouterGroup, appContext appctx.AppContext) {
	submission := group.Group("/submission")
	submission.POST("/submit", transport.HandleSubmitSubmissionRequest(appContext))
}
