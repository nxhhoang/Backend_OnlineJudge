package domain

import (
	"time"

	"github.com/bibimoni/Online-judge/submission-judge/src/pkg/memory"
)

// This will be the arguments passed to isolate
type RunConfig struct {
	TimeLimit        time.Duration
	MemoryLimit      memory.Memory
	WorkingDirectory string
	DirectoryMaps    []DirectoryMap
	Env              []string
	InheritEnv       bool
	MaxProcesses     int
}

type DirectoryMap struct {
	Inside  string
	Outside string
	Options []DirectoryMapOption
}

type DirectoryMapOption string

const (
	AllowSpecial   DirectoryMapOption = "dev"
	MountFS        DirectoryMapOption = "fs"
	Maybe          DirectoryMapOption = "maybe"
	NoExec         DirectoryMapOption = "noexec"
	AllowReadWrite DirectoryMapOption = "rw"
	Temporary      DirectoryMapOption = "tmp"
)
