package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	LogLevel   string
	Enviroment string
	Server     struct {
		Port         string
		Host         string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}
	Endpoints struct {
		Submission string
	}
}

var cfg *Config = nil

func Load() *Config {
	godotenv.Load("../../../.env")

	if cfg != nil {
		return cfg
	}
	cfg = &Config{}

	cfg.Enviroment = getEnv("ENV", "Development")

	cfg.LogLevel = getEnv("LOG_LEVEL", "debug")

	cfg.Server.Port = getEnv("PORT", "80")
	cfg.Server.Host = getEnv("HOST", "0.0.0.0")
	cfg.Server.ReadTimeout = time.Second * 15
	cfg.Server.WriteTimeout = time.Second * 15

	cfg.Endpoints.Submission = getEnv("SUBMISSION_ENDPOINT", "")
	return cfg
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
