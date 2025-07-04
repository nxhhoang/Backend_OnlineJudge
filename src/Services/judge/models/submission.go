package models

type Submission struct {
	ID             uint64
	UserID         uint64
	ProblemID      uint64
	SubmissionDate uint64
	JudgeStatus    []uint8
}

type judgeStatus string

// "AC", "WA", "RTE", "IR", "OLE", "MLE", "TLE", "IE"
const (
	AC  judgeStatus = "AC"
	WA  judgeStatus = "WA"
	RTE judgeStatus = "RTE"
	IR  judgeStatus = "IR"
	OLE judgeStatus = "OLE"
	MLE judgeStatus = "MLE"
	TLE judgeStatus = "TLE"
	IE  judgeStatus = "IE"
)
