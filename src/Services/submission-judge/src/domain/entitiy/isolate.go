package domain

import (
	"time"

	"github.com/bibimoni/Online-judge/submission-judge/src/pkg/memory"
	"github.com/rs/zerolog"
)

// This will be a sandbox implementation, it will call isolate command.
// It needs isolate installed in the system to work
type Isolate struct {
	ID       int
	BoxDir   string
	MetaFile string
	Logger   *zerolog.Logger
	Inited   bool
}

// This will be the arguments passed to isolate
type RunConfig struct {
	TimeLimit   time.Duration
	MemoryLimit memory.Memory
}
