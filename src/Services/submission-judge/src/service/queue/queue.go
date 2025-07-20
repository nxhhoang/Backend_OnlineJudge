package queueservice

import (
	"fmt"

	"github.com/bibimoni/Online-judge/submission-judge/src/service/queue/impl"
)

type QueueService interface {
}

func NewQueueService() (QueueService, error) {
	taskQueueService, err := impl.NewQueueServiceImpl()
	if err != nil {
		return nil, fmt.Errorf("Error when create new Queue %v", err)
	}
	return taskQueueService, nil
}
