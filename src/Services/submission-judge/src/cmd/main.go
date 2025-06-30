package main

import (
	"context"
	// "fmt"

	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/database"
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
}
