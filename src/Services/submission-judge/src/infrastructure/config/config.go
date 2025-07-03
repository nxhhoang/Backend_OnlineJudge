package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Database struct {
		Uri  string
		Name string
	}
	LogLevel   string
	Enviroment string
	Server     struct {
		Port         string
		Host         string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}
}

func Load() (*Config, error) {
	godotenv.Load("../../../.env")

	cfg := &Config{}

	cfg.Database.Uri = getEnv("MONGODB_URI", "mongodb://mongo:37017/submissionjudgedb")
	cfg.Database.Name = getEnv("MONGODB_DATABASE_NAME", "submissionjudgedb")
	cfg.Enviroment = getEnv("ENV", "Development")

	cfg.LogLevel = getEnv("LOG_LEVEL", "debug")

	cfg.Server.Port = getEnv("PORT", "8080")
	cfg.Server.Host = getEnv("HOST", "http://localhost")
	cfg.Server.ReadTimeout = time.Second * 15
	cfg.Server.WriteTimeout = time.Second * 15

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
