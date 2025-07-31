package cpp

import (
	domain "github.com/bibimoni/Online-judge/submission-judge/src/domain/entitiy"
	"github.com/bibimoni/Online-judge/submission-judge/src/pkg"
)

type Cpp struct {
	id          string
	name        string
	compileArgs []string
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

func (cpp Cpp) Run(i *domain.Isolate) error {

	return nil
}

func (cpp Cpp) Compile(i *domain.Isolate) error {
	return nil
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
