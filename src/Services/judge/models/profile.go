package models

type Profile struct {
	Username string
	UID      string
	password string // hashed

	FirstName string
	LastName  string
	Email     string
	About     string

	Points            float64
	PerformancePoints float64
	ProblemCount      uint64
}

type Organization struct {
	Name      string
	ShortName string
	About     string
	admins    []Profile
}
