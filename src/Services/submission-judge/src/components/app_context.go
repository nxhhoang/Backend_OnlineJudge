package appctx

import (
	queueservice "github.com/bibimoni/Online-judge/submission-judge/src/service/queue"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AppContext interface {
	GetMainDbConnection() *mongo.Database
	GetQueueService() *queueservice.QueueService
}

type appCtx struct {
	database     *mongo.Database
	queueService *queueservice.QueueService
}

func NewAppContext(database *mongo.Database, queueService *queueservice.QueueService) *appCtx {
	return &appCtx{
		database,
		queueService,
	}
}

func (ctx *appCtx) GetQueueService() *queueservice.QueueService {
	return ctx.queueService
}

func (ctx *appCtx) GetMainDbConnection() *mongo.Database {
	return ctx.database
}
