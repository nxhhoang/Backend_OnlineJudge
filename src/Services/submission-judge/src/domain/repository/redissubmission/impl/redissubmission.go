package impl

import (
	"context"
	"encoding/json"
	"fmt"

	repository "github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/redissubmission"
	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
	usecase "github.com/bibimoni/Online-judge/submission-judge/src/usecase/wssubmission"
	"github.com/redis/go-redis/v9"
)

type RedisSubmissionRepositoryImpl struct {
	rdb *redis.Client
}

func NewRedisSubmissionRepositoryImpl(rdb *redis.Client) *RedisSubmissionRepositoryImpl {
	return &RedisSubmissionRepositoryImpl{
		rdb,
	}
}

func (rs *RedisSubmissionRepositoryImpl) GetChannelString(problemId, username, submissionId string) string {
	return fmt.Sprintf("%s:%s:%s", problemId, username, submissionId)
}

func NewRedisSubmissionRepository(rdb *redis.Client) repository.RedisSubmissionRepository {
	return NewRedisSubmissionRepositoryImpl(rdb)
}

func (rs *RedisSubmissionRepositoryImpl) PulishSubmission(ctx context.Context, res usecase.WSSubmissionResponse) error {
	channel := rs.GetChannelString(res.ProblemId, res.Username, res.SubmissionId)

	byte, err := json.Marshal(res)
	if err != nil {
		return err
	}
	config.GetLogger().Debug().Msgf("Channel: %s --- Receive event %v", channel, res)
	return rs.rdb.Publish(ctx, channel, byte).Err()
}

func (rs *RedisSubmissionRepositoryImpl) Subscribe(ctx context.Context, channelId string) (<-chan *usecase.WSSubmissionResponse, error) {
	sub := rs.rdb.PSubscribe(ctx, channelId)
	raw := sub.Channel()

	out := make(chan *usecase.WSSubmissionResponse)

	go func() {
		defer sub.Close()
		defer close(out)

		for {
			select {
			case msg, ok := <-raw:
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
	}()
	return out, nil
}
