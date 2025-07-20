package impl

import (
	"context"
	"encoding/json"

	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
	"github.com/hibiken/asynq"
)

type QueueServiceImpl struct {
	queue *domain.Queue
}

type SubmissionTaskPayload struct {
	submission *domain.Submission
}

const (
	SUBMISSION_SUBMIT string = "submission:submit"
)

const WORKER_NUM int = 10

func NewQueueServiceImpl() (*QueueServiceImpl, error) {
	client, err := config.NewAsyncqClient()
	if err != nil {
		return nil, err
	}

	server, err := config.NewAsyncqServer(WORKER_NUM)
	if err != nil {
		return nil, err
	}

	log := config.GetLogger()
	return &QueueServiceImpl{
		queue: &domain.Queue{
			Client: client,
			Server: server,
			Logger: log,
		},
	}, nil
}

func (qs *QueueServiceImpl) AddSubmission(submission *domain.Submission) error {
	payload, err := json.Marshal(SubmissionTaskPayload{submission})
	if err != nil {
		return err
	}

	task := asynq.NewTask(SUBMISSION_SUBMIT, payload)

	info, err := qs.queue.Client.Enqueue(task)
	if err != nil {
		return err
	}

	qs.queue.Logger.Info().Msgf("Successfully enqueued task: %+v", info)

	return nil
}

func (qs *QueueServiceImpl) RunServer() error {
	mux := asynq.NewServeMux()

	if err := qs.queue.Server.Run(mux); err != nil {
		return err
	}
	qs.queue.Logger.Info().Msgf("Server is running")
	return nil
}

func judgeSubmission(ctx context.Context, task *asynq.Task) error {
	var submisisonPayLoad SubmissionTaskPayload
	if err := json.Unmarshal(task.Payload(), &submisisonPayLoad); err != nil {
		return err
	}

	// send submission into isolate

	return nil
}
