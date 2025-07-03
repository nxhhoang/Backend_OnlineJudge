package main

import (
	"context"
	// "fmt"
	"github.com/bibimoni/Online-judge/submission-judge/src/components"
	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/database"
	"github.com/bibimoni/Online-judge/submission-judge/src/router"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	log := config.NewLogger(cfg.LogLevel)
	if err != nil {
		log.Fatal().Err(err).Msgf("Can not load config")
	}

	client, err := database.GetMongoDbClient(cfg.Database.Uri)
	if err != nil {
		log.Fatal().Err(err).Msgf("Can't not load mongoDB")
	}
	defer client.Disconnect(context.Background())

	appCtx := appctx.NewAppContext(client)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	v1 := r.Group("/api/v1")

	router.RegisterRouter(v1, appCtx)
}
