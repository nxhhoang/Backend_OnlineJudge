package router

import (
	"github.com/bibimoni/Online-judge/submission-judge/src/components"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(gin *gin.RouterGroup, appCtx appctx.AppContext) {
	submission := gin.Group("/submission")

}
