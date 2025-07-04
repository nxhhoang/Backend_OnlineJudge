package models

type Problem struct {
	Name      string
	ShortName string
	url       string

	Generators []executable
	Validators []Validator
	Checker    Checker
	Solutions  []Solutions
}

type executable struct {
	Source     string
	SourceType string
	Binary     string
	// In Polygon package, binaries have type properties (exe.win32,...) but in our system,
	// We would decide it ourselves
}

type Validator struct {
	executable
	tests Testset
}

type Checker struct {
	executable
	tests Testset
}

type Solutions struct {
	Tags string
	executable
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
	Cmd        string
	InputName  string
	OutputName string
	Point      float64
	Group      string
}

type Group struct {
	FeedbackPolict string
	Name           string
	PointsPolicy   string
	Points         float64
}
