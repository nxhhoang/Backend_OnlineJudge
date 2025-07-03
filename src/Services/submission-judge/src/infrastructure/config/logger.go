package config

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

var log *zerolog.Logger

func GetLogger() *zerolog.Logger {
	return log
}

func NewLogger(level string) *zerolog.Logger {
	var logLevel zerolog.Level
	if level == "info" {
		logLevel = zerolog.InfoLevel
	} else {
		logLevel = zerolog.DebugLevel
	}
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		NoColor:    false,
		TimeFormat: time.RFC3339,
		FormatLevel: func(i any) string {
			return fmt.Sprintf("[%s]", i)
		},
		FormatCaller: func(i any) string {
			return fmt.Sprintf("[%s]", i)
		},
		FormatMessage: func(i any) string {
			return fmt.Sprintf("{ %-20s }", i)
		},
	}
	logger := zerolog.New(consoleWriter).
		Level(logLevel).
		With().
		Caller().
		Timestamp().
		Logger()

	log = &logger
	return &logger
}
