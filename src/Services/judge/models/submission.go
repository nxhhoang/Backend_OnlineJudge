package models

type Submission struct {
	ID uint64
	UserID 	uint64
	ProblemID uint64
	SubmissionDate uint64
	submissionPath string // stores submission-related files
	JudgeStatus []uint8
}
