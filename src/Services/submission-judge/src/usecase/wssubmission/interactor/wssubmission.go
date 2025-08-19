package interactor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
	usecase "github.com/bibimoni/Online-judge/submission-judge/src/usecase/wssubmission"
	"github.com/redis/go-redis/v9"
)

type WSSubmissionInteractor struct {
	rdb *redis.Client
}

func NewWSSubmissionInteractor(rdb *redis.Client) *WSSubmissionInteractor {
	return &WSSubmissionInteractor{
		rdb,
	}
}

func (wss *WSSubmissionInteractor) SubmissionStatus(ctx context.Context, input *usecase.WSSubmissionInput, out chan<- *usecase.WSSubmissionResponse) {
	channel := fmt.Sprintf("%s:%s:%s", input.ProblemId, input.Username, input.SubmissionId)
	sub := wss.rdb.Subscribe(ctx, channel)
	defer sub.Close()
	ch := sub.Channel()
	config.GetLogger().Debug().Msgf("Test !!")
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				return
			}
			var upd usecase.WSSubmissionResponse
			err := json.Unmarshal([]byte(msg.Payload), &upd)
			if err != nil {
				config.GetLogger().Error().Msgf("Json Unmarshal failed: %v", err)
				continue
			}

			out <- &upd
		case <-ctx.Done():
			return
		}
	}
}
