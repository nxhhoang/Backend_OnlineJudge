package pkg

import (
	"io"

	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	isolateservice "github.com/bibimoni/Online-judge/submission-judge/src/service/isolate"
)

type Language interface {
	ID() string
	DisplayName() string
	DefaultFileName() string
	ExecutableName() string
	FileExtension() string
	Run(i *domain.Isolate, rc *domain.RunConfig, req *isolateservice.SubmissionRequest) error
	Compile(i *domain.Isolate, req *isolateservice.SubmissionRequest, stderr io.Writer) error
}
