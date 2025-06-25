package models

type Problem struct {
	Code       string
	Name       string
	Statement  string
	Tags       []string
	Validators []Validator
	Checker    Checker
	Solutions  []Solutions
}

type Executable struct {
	Source     string
	SourceType string
	Binary     string
	// In Polygon package, binaries have type properties (win32,...) but in our system,
	// We would decide it ourselves
	Testset Testset
}

type Validator struct {
	Executable
}

type Checker struct {
	Executable
}

type Solutions struct {
	Tags string
	Executable
}

type Judging struct {
}

type Testset struct {
	TimeLimit         uint64
	MemoryLimit       uint64
	TestCount         uint64
	InputTestPattern  string
	OutputTestPattern string
	Tests             []Test
}

type Test struct {
	Comment string
}
