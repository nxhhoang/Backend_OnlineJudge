package judge

import (
	"context"
	"errors"

	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	"github.com/bibimoni/Online-judge/submission-judge/src/pkg"
	"github.com/bibimoni/Online-judge/submission-judge/src/pkg/memory"
	isolateservice "github.com/bibimoni/Online-judge/submission-judge/src/service/isolate"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/problem"
)

type JudgeService interface {
	Judge(ctx context.Context, req *isolateservice.SubmissionRequest, problemInfo *problem.ProblemServiceGetOutput) error
	JudgeStart(ctx context.Context, i *domain.Isolate, lang pkg.Language, req *isolateservice.SubmissionRequest, problemInfo *problem.ProblemServiceGetOutput) error
	Prep(ctx context.Context, i *domain.Isolate, lang pkg.Language, req *isolateservice.SubmissionRequest, problemInfo *problem.ProblemServiceGetOutput) error
}

type RunVerdict struct {
	Status         string        `json:"status"`
	ExitCode       int           `json:"exitcode"`
	Message        string        `json:"message"`
	Time           float64       `json:"time"`
	TimeWall       float64       `json:"time-wall"`
	CgMem          memory.Memory `json:"cg-mem"`
	CgMemSw        memory.Memory `json:"cg-mem-sw"`
	MaxRss         int           `json:"max-rss"`
	Csw            int           `json:"csw"`
	CswForced      int           `json:"csw-forced"`
	CgOomKilled    int           `json:"cg-oom-killed"`
	ExitedNormally bool          `json:"exited-normally"`
	KilledBySignal int           `json:"killed"`
}

var CompilationError = errors.New("Compilation Error")
var JugdgementFailed = errors.New("Something went wrong when trying to run the source code")
