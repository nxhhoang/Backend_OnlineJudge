package utils_test

import (
	"fmt"
	"judge/utils"
	"os"
	"testing"
)

func TestParseProblem(t *testing.T) {
	f, err := os.Open("problem.xml")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	problem, err := utils.ParseProblemStruct(f)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v", problem)
}
