package appctx

import "go.mongodb.org/mongo-driver/v2/mongo"

type AppContext interface {
	GetMainDbConnection() *mongo.Database
}

type appCtx struct {
	database *mongo.Database
}

func NewAppContext(database *mongo.Database) *appCtx {
	return &appCtx{database}
}

func (ctx *appCtx) GetMainDbConnection() *mongo.Database {
	return ctx.database
}
