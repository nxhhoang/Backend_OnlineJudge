package config

import "github.com/hibiken/asynq"

// This will handle the distribution of submission to our isolate instances
func NewAsyncqServer(concurrency int) (*asynq.Server, error) {
	cfg, err := Load()
	if err != nil {
		return nil, err
	}
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     cfg.Redis.Uri,
			Password: cfg.Redis.Password,
		},
		asynq.Config{
			Concurrency: concurrency,
		},
	)
	return srv, nil
}

func NewAsyncqClient() (*asynq.Client, error) {
	cfg, err := Load()
	if err != nil {
		return nil, err
	}
	client := asynq.NewClient(
		asynq.RedisClientOpt{
			Addr:     cfg.Redis.Uri,
			Password: cfg.Redis.Password,
		},
	)

	return client, nil
}
