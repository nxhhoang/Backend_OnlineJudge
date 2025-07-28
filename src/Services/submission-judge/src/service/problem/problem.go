package problem

// This will handle the calling process, the retrieving process of getting problem
// served by the Problem service (the name might be a little confusing)
type ProblemService interface {
	Get(problemId string) (ProblemServiceGetOutput, error)
}

type ProblemServiceGetOutput struct {
	ID          string   `json:"ID,omitempty"`
	ProblemId   int64    `json:"problem-id,omitempty"`
	Name        string   `json:"name,omitempty"`
	ShortName   string   `json:"short-name,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	TestNum     int      `json:"test-num,omitempty"`
	TimeLimit   int      `json:"time-limit,omitempty"`
	MemoryLimit int64    `json:"memory-limit,omitempty"`
}
