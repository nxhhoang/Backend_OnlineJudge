package domain

import "time"

type Contestant struct {
	UserID    uint64    `bson:"user_id"`
	Points    uint64    `bson:"points"`
	RealStart time.Time `bson:"real_start"`

	Submissions []uint64 `bson:"submissions"`
}
