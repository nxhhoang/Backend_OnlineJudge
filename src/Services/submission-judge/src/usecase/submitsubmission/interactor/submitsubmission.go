package interactor

import (
	"context"

	scr "github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/sourcecode"
	sr "github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/submission"
	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
	"github.com/bibimoni/Online-judge/submission-judge/src/usecase/submitsubmission"
)

type SubmissionInteractor struct {
	submissionRepo sr.SubmissionRepository
	sourcecodeRepo scr.SourcecodeRepository
}

func NewSubmissionInteractor(sr sr.SubmissionRepository, scr scr.SourcecodeRepository) *SubmissionInteractor {
	return &SubmissionInteractor{
		submissionRepo: sr,
		sourcecodeRepo: scr,
	}
}

func (si *SubmissionInteractor) SubmitSubmission(ctx context.Context, input *usecase.SubmitSubmissionInput) (output *usecase.SubmitSubmissionResponse, err error) {
	log := config.GetLogger()
	log.Info().Msgf("User %s submitted a solution in %s, for problem with problem id: %s", input.Username, input.LanguageId, input.ProblemId)

	codeId, err := si.sourcecodeRepo.CreateSourcecode(ctx, input.Code, input.LanguageId)
	if err != nil {
		log.Debug().Msgf("An error occured: %v", err)
		return nil, err
	}

	return &usecase.SubmitSubmissionResponse{
		ID: codeId,
	}, nil
}
