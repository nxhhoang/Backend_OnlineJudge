package interactor

import (
	"context"
	"strconv"

	scr "github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/sourcecode"
	sr "github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/submission"
	"github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/submission/impl"
	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
	// "github.com/bibimoni/Online-judge/submission-judge/src/pkg"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/problem"
	// "github.com/bibimoni/Online-judge/submission-judge/src/service/store"
	usecase "github.com/bibimoni/Online-judge/submission-judge/src/usecase/submitsubmission"
)

type SubmissionInteractor struct {
	submissionRepo sr.SubmissionRepository
	sourcecodeRepo scr.SourcecodeRepository
	problemService problem.ProblemService
}

func NewSubmissionInteractor(
	sr sr.SubmissionRepository,
	scr scr.SourcecodeRepository,
	ps problem.ProblemService,
) *SubmissionInteractor {
	return &SubmissionInteractor{
		submissionRepo: sr,
		sourcecodeRepo: scr,
		problemService: ps,
	}
}

func (si *SubmissionInteractor) SubmitSubmission(ctx context.Context, input *usecase.SubmitSubmissionInput) (output *usecase.SubmitSubmissionResponse, err error) {
	log := config.GetLogger()
	log.Info().Msgf("User %s submitted a solution in %s, for problem with problem id: %s", input.Username, input.LanguageId, input.ProblemId)

	codeId, err := si.sourcecodeRepo.CreateSourcecode(ctx, input.Code, input.LanguageId)
	if err != nil {
		return nil, err
	}

	problemInfo, err := si.problemService.Get(ctx, input.ProblemId)
	if err != nil {
		return nil, err
	}

	params := impl.CreateSubmissionInput{
		ProblemId:    strconv.FormatInt(problemInfo.ProblemId, 10),
		Username:     input.Username,
		Type:         input.SubmissionType,
		SourceCodeId: codeId,
	}

	submissionId, err := si.submissionRepo.CreateSubmission(ctx, params)
	if err != nil {
		return nil, err
	}
	// req := pkg.SubmissionRequest{
	// 	SubmissionId:   submissionId,
	// 	Username:       input.Username,
	// 	Sourcecode:     input.Code,
	// 	SubmissionType: input.SubmissionType,
	// 	ProblemId:      input.ProblemId,
	// }

	// lang, err := store.DefaultStore.Get(input.LanguageId)
	if err != nil {
		return nil, err
	}

	log.Info().Msgf("Enqueued submission successfully, id: %s", submissionId)

	return &usecase.SubmitSubmissionResponse{
		// ID: codeId,
		Message: "Submit successfully!",
	}, nil
}
