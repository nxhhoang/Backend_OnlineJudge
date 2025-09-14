package interactor

import (
	"errors"

	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
)

type InteractorService interface {
	RunInteractor(interactorAddr string, inputAddr string, outputAddr string, answerAddr string) (domain.Verdict, int, string, error)
}

var FileNotExits = errors.New("File not exists on the system")
