package impl_test

import (
	"context"
	"testing"

	"github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/submission/impl"
	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/database"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func TestSubmissionRepositoryImpl_FindAllProblemSubmissionIds(t *testing.T) {
	cfg, err := config.Load()
	log := config.GetLogger()
	client, err := database.GetMongoDbClient(cfg.Database.Uri)
	if err != nil {
		log.Fatal().Err(err).Msgf("Can't not load mongoDB")
	}
	defer client.Disconnect(context.Background())
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		db *mongo.Database
		// Named input parameters for target function.
		problemId string
		want      []string
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			name:      "test pull submission",
			db:        client.Database("submissionjudgedb"),
			problemId: "445985",
			want:      make([]string, 0),
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := impl.NewSubmissionRepositoryImpl(tt.db)
			got, gotErr := sr.FindAllProblemSubmissionIds(context.Background(), tt.problemId)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("FindAllProblemSubmissionIds() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("FindAllProblemSubmissionIds() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			config.GetLogger().Debug().Msgf("%v", got)
			// if true {
			// 	t.Errorf("FindAllProblemSubmissionIds() = %v, want %v", got, tt.want)
			// }
		})
	}
}
