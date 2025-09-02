package main

import (
	"os"

	database "contest/internal/infrastructure/database"
	infrastructure "contest/internal/infrastructure/redis"
	"contest/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	app := fiber.New(fiber.Config{
		ServerHeader: "HCMUT-OJ",
	})

	port := os.Getenv("CONTEST_PORT")
	if port == "" {
		log.Print("Can't get port address")
		return
	}

	host := os.Getenv("CONTEST_HOST")
	if host == "" {
		log.Print("Can't get host address")
		return
	}

	if err := database.GetMongoDbClient(); err != nil {
		panic(err)
	}

	if err := infrastructure.GetRedisClient(); err != nil {
		panic(err)
	}

	routes.ScoreboardRoutes(app)
	routes.ContestRoutes(app)

	serverAddr := host + ":" + port
	app.Listen(serverAddr)
}
