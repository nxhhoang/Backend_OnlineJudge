package cpp

import (
	"io"
	"time"

	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	"github.com/bibimoni/Online-judge/submission-judge/src/pkg"
	"github.com/bibimoni/Online-judge/submission-judge/src/pkg/memory"
	isolateservice "github.com/bibimoni/Online-judge/submission-judge/src/service/isolate"
	"github.com/bibimoni/Online-judge/submission-judge/src/service/isolate/utils"
)

type Cpp struct {
	id          string
	name        string
	compileArgs []string
	needCompile bool
}

func (cpp Cpp) ID() string {
	return cpp.id
}

func (cpp Cpp) DisplayName() string {
	return cpp.name
}

func (cpp Cpp) DefaultFileName() string {
	return "main.cpp"
}

func (cpp Cpp) FileExtension() string {
	return "cpp"
}

func (cpp Cpp) ExecutableName() string {
	return "main"
}

func (cpp Cpp) NeedCompile() bool {
	return true
}

func (cpp Cpp) Run(i *domain.Isolate, req *isolateservice.SubmissionRequest) error {

	return nil
}

// Compile file, make sure the file is present inside the isolate working directory
func (cpp Cpp) Compile(i *domain.Isolate, req *isolateservice.SubmissionRequest, stderr io.Writer) error {
	// compile
	rc := domain.RunConfig{
		TimeLimit:    time.Second * 10,
		MemoryLimit:  256 * memory.MiB,
		MaxProcesses: 200,
		InheritEnv:   true,
		Stdout:       stderr,
		Stderr:       stderr,
		Meta:         true,
	}

	runArgs := cpp.compileArgs
	runArgs = append(runArgs, []string{
		utils.GetMappedFileNamePath(cpp.DefaultFileName()),
		"-o",
		utils.GetMappedFileNamePath(cpp.ExecutableName()),
	}...)

	i.Logger.Info().Msgf("Start compiling source code with id: %s", req.SubmissionId)

	return req.IService.Run(
		i, rc, req, "/usr/bin/g++", runArgs...,
	)
}

var DefaultCompileArgs = []string{"-O2", "-static", "-DONLINE_JUDGE"}

var cpp11 = Cpp{
	id:          "cpp11",
	name:        "C++11",
	compileArgs: append(DefaultCompileArgs, "-std=c++11"),
}

var cpp14 = Cpp{
	id:          "cpp14",
	name:        "C++14",
	compileArgs: append(DefaultCompileArgs, "-std=c++14"),
}

var cpp17 = Cpp{
	id:          "cpp17",
	name:        "C++17",
	compileArgs: append(DefaultCompileArgs, "-std=c++17"),
}

var cpp20 = Cpp{
	id:          "cpp20",
	name:        "C++20",
	compileArgs: append(DefaultCompileArgs, "-std=c++20"),
}

func GetAllOptions() []pkg.Language {
	return []pkg.Language{cpp11, cpp14, cpp17, cpp20}
}
