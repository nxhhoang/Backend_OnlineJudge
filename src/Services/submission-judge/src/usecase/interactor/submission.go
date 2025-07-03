package interactor

import (
	"context"

	"github.com/bibimoni/Online-judge/submission-judge/src/domain/repository"
	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
	"github.com/bibimoni/Online-judge/submission-judge/src/usecase"
)

type SubmissionInteractor struct {
	repository repository.SubmissionRepository
}

func NewSubmissionInteractor(repository repository.SubmissionRepository) *SubmissionInteractor {
	return &SubmissionInteractor{
		repository: repository,
	}
}

func (si *SubmissionInteractor) SubmitSubmission(ctx context.Context, input *usecase.SubmitSubmissionInput) (output *usecase.SubmitSubmissionResponse, err error) {
	log := config.GetLogger()
	log.Info().Msgf("Submitting")

	return &usecase.SubmitSubmissionResponse{
		ID: "test",
	}, nil
}
