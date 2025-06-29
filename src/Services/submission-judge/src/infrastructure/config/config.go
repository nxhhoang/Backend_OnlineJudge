package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Database struct {
		Port   string
		DBName string
		Host   string
	}

	Enviroment string
}

func Load() (*Config, error) {
	godotenv.Load("../../../.env")

	cfg := &Config{}

	cfg.Database.Port = getEnv("MONGODB_PORT", "37017")
	cfg.Database.DBName = getEnv("MONGODB_DBNAME", "submissionjudgedb")
	cfg.Database.Host = getEnv("MONGODB_HOST", "mongo")

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
