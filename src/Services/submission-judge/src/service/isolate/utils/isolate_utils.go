package utils

import (
	"fmt"
	"os"
	"path/filepath"

	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/isolate/impl"
)

func CreateSubmissionSourceFile(i *domain.Isolate, sourceCode string, submissionId string, sourceCodeName string) (*os.File, error) {
	sourceCodeAddr := impl.GetIsolateDir(i) + "/" + submissionId + "/" + sourceCodeName
	dir := filepath.Dir(sourceCodeAddr)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	file, err := os.Create(sourceCodeAddr)
	if err != nil {
		return nil, fmt.Errorf("Failed to create source code in: %s", sourceCodeAddr)
	}
	_, err = file.WriteString(sourceCode)
	if err != nil {
		return nil, err
	}
	return file, nil
}
