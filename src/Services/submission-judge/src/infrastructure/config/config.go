package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database struct {
		Uri string
	}
	LogLevel   string
	Enviroment string
}

func Load() (*Config, error) {
	godotenv.Load("../../../.env")

	cfg := &Config{}
	cfg.Database.Uri = getEnv("MONGODB_URI", "mongodb://mongo:37017/submissionjudgedb")
	cfg.Enviroment = getEnv("ENV", "Development")
	cfg.LogLevel = getEnv("LOG_LEVEL", "debug")
	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
