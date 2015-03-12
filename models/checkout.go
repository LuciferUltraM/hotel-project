package models

import "time"

type CheckOut struct {
	CheckOutNo   int64
	CheckOutDate time.Time
	Fine         float32
}
