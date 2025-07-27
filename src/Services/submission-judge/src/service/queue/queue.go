package queueservice

import (
	"fmt"

	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/queue/impl"
)

type QueueService interface {
	RunServer() error
	AddSubmission(s *domain.Submission) error
}

func NewQueueService() (QueueService, error) {
	taskQueueService, err := impl.NewQueueServiceImpl()
	if err != nil {
		return nil, fmt.Errorf("Error when create new Queue %v", err)
	}
	return taskQueueService, nil
}
