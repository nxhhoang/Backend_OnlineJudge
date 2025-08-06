package judge

import (
	"context"

	isolateservice "github.com/bibimoni/Online-judge/submission-judge/src/service/isolate"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/problem"
)

type JudgeService interface {
	Judge(ctx context.Context, req *isolateservice.SubmissionRequest, problemInfo *problem.ProblemServiceGetOutput) error
}
