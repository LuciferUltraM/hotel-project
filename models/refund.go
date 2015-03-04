package models

import "time"

type Refund struct {
	RefundNo   int
	RefundDate time.Time
	Status     string
	Amount     float32
	CreateDate time.Time
}
