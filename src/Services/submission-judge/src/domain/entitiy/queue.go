package domain

import (
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog"
)

type Queue struct {
	Client *asynq.Client
	Server *asynq.Server
	Logger *zerolog.Logger
}
