package appctx

import (
	poolservice "github.com/bibimoni/Online-judge/submission-judge/src/service/pool"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AppContext interface {
	GetMainDbConnection() *mongo.Database
	GetPool() *poolservice.PoolService
}

type appCtx struct {
	database *mongo.Database
	pool     *poolservice.PoolService
}

func NewAppContext(database *mongo.Database, pool *poolservice.PoolService) *appCtx {
	return &appCtx{
		database,
		pool,
	}
}

func (ctx *appCtx) GetMainDbConnection() *mongo.Database {
	return ctx.database
}

func (ctx *appCtx) GetPool() *poolservice.PoolService {
	return ctx.pool
}
