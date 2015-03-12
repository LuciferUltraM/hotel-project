package models

import "time"

type CheckIn struct {
	CheckInNo   int64
	CheckInDate time.Time
	Deposit     float32
}
