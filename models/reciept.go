package models

import "time"

type Receipt struct {
	ReceiptNo   int
	ReceiptDate time.Time
	Amount      float32
	CreateDate  time.Time
}
