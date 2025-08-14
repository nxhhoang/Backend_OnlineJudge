package domain

import "time"

type Contestant struct {
	UserID    uint64    `bson:"user-id"`
	RealStart time.Time `bson:"real-start"`

	Submissions []uint64 `bson:"submissions"`

	TotalPoints float64   `bson:"total-points"`
	Points      []float64 `bson:"points"`
}
