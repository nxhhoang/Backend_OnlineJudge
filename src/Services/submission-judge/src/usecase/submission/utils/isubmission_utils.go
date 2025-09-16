package isubmission_utils

import (
	"context"

	erepository "github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/evaluation"
	screpository "github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/sourcecode"
	srepository "github.com/bibimoni/Online-judge/submission-judge/src/domain/repository/submission"
	"github.com/bibimoni/Online-judge/submission-judge/src/infrastructure/config"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/store"
	usecase "github.com/bibimoni/Online-judge/submission-judge/src/usecase/submission"
	usecasews "github.com/bibimoni/Online-judge/submission-judge/src/usecase/wssubmission"
)

func GetSubmission(
	ctx context.Context,
	evalRepo erepository.EvaluationRepository,
	sourcecodeRepo screpository.SourcecodeRepository,
	submissionRepo srepository.SubmissionRepository,
	submissionId string,
) (*usecase.GetSubmissionOutput, error) {
	log := config.GetLogger()
	sub, err := submissionRepo.FindSubmission(ctx, submissionId)
	if err != nil {
		log.Debug().Msgf("Find submission, error: %v", err)
		return nil, err
	}

	source, err := sourcecodeRepo.GetSourceBySubmissionId(ctx, (*sub).Id)
	if err != nil {
		log.Debug().Msgf("Find sourcecode, error: %v", err)
		return nil, err
	}

	eval, err := evalRepo.GetEvalBySubmissionId(ctx, (*sub).Id)
	if err != nil {
		log.Debug().Msgf("Find eval, error: %v", err)
		return nil, err
	}

	lang, err := store.DefaultStore.Get((*source).LanguageId)
	if err != nil {
		return nil, err
	}

	returnVal := usecase.GetSubmissionOutput{
		ProblemId:       (*sub).ProblemId,
		Verdict:         (*eval).Verdict,
		VerdictCase:     (*eval).VerdictCase,
		CpuTime:         (*eval).CpuTime,
		CpuTimeCase:     (*eval).CpuTimeCase,
		MemoryUsage:     (*eval).MemoryUsage,
		MemoryUsageCase: (*eval).MemoryUsageCase,
		NSuccess:        (*eval).NSuccess,
		Outputs:         (*eval).Outputs,
		Message:         (*eval).Message,
		Points:          (*eval).Points,
		PointsCase:      (*eval).PointsCase,
		NCases:          (*eval).NCases,
		TL:              (*eval).TL,
		ML:              (*eval).ML,
		Username:        (*sub).Username,
		Timestamp:       (*sub).Timestamp,
		Type:            (*sub).Type,
		Language:        lang.DisplayName(),
		SourceCode:      (*source).SourceCode,
		EvalStatus:      (*eval).EvalStatus,
	}

	return &returnVal, nil
}

func GetSubmissionWithoutSourceCode(
	ctx context.Context,
	evalRepo erepository.EvaluationRepository,
	sourcecodeRepo screpository.SourcecodeRepository,
	submissionRepo srepository.SubmissionRepository,
	submissionId string,
) (*usecasews.WSSubmissionResponse, error) {

	sub, err := GetSubmission(
		ctx,
		evalRepo,
		sourcecodeRepo,
		submissionRepo,
		submissionId,
	)

	if err != nil {
		return nil, err
	}
	returnVal := usecasews.WSSubmissionResponse{
		Username:        sub.Username,
		SubmissionId:    submissionId,
		ProblemId:       sub.ProblemId,
		Timestamp:       sub.Timestamp,
		Language:        sub.Language,
		Verdict:         sub.Verdict,
		VerdictCase:     sub.VerdictCase,
		CpuTime:         sub.CpuTime,
		CpuTimeCase:     sub.CpuTimeCase,
		MemoryUsage:     sub.MemoryUsage,
		MemoryUsageCase: sub.MemoryUsageCase,
		NSuccess:        sub.NSuccess,
		Outputs:         sub.Outputs,
		Points:          sub.Points,
		PointsCase:      sub.PointsCase,
		Message:         sub.Message,
		EvalStatus:      sub.EvalStatus,
	}

	return &returnVal, nil
}
