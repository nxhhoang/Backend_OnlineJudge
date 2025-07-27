package main

import (
	"context"
	// "fmt"
	"net/http"

	appctx "github.com/bibimoni/Online-judge/submission-judge/src/components"
	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/database"
	"github.com/bibimoni/Online-judge/submission-judge/src/router"
	queueservice "github.com/bibimoni/Online-judge/submission-judge/src/service/queue"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic("Can't load config")
	}
	log := config.NewLogger(cfg.LogLevel)

	client, err := database.GetMongoDbClient(cfg.Database.Uri)
	if err != nil {
		log.Fatal().Err(err).Msgf("Can't not load mongoDB")
	}
	defer client.Disconnect(context.Background())

	queueService, err := queueservice.NewQueueService()
	if err != nil {
		panic("Can't initialize new queue service: " + err.Error())
	}

	appCtx := appctx.NewAppContext(client.Database(cfg.Database.Name), &queueService)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	v1 := r.Group("/api/v1")

	router.RegisterRouter(v1, appCtx)

	serverAddr := cfg.Server.Host + ":" + cfg.Server.Port

	log.Info().Msgf("Submission-Judge server is listening on: %s", serverAddr)

	srv := &http.Server{
		Addr:         serverAddr,
		Handler:      r,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msgf("Failed to start Submission-Judge server: %s", err)
	}
}
