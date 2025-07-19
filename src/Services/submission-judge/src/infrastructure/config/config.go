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
	LogLevel        string
	SandboxLogLevel string
	Enviroment      string
	Server          struct {
		Port         string
		Host         string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}
	ProblemsDir string
}

func Load() (*Config, error) {
	godotenv.Load()
	cfg := &Config{}

	cfg.Database.Uri = getEnv("SUBMISSION_MONGODB_URI", "mongodb://mongosubmissiondb:37017/submissionjudgedb")
	cfg.Database.Name = getEnv("SUBMISSION_MONGODB_DATABASE_NAME", "submissionjudgedb")

	cfg.Enviroment = getEnv("SUBMISSION_ENV", "Development")

	cfg.LogLevel = getEnv("SUBMISSION_LOG_LEVEL", "debug")
	cfg.SandboxLogLevel = getEnv("SUBMISSION_SANDBOX_LOG_LEVEL", "debug")

	cfg.ProblemsDir = getEnv("SUBMISSION_JUDGE_PROBLEM_DIR", "problems_dir")

	cfg.Server.Port = getEnv("SUBMISSION_PORT", "8000")
	cfg.Server.Host = getEnv("SUBMISSION_HOST", "0.0.0.0")
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
