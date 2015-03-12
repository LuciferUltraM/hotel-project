package models

import "time"

type Receipt struct {
	ReceiptNo   int
	ReceiptDate time.Time
	Status      string
	Amount      float32
	Type        string
	RoomBooking *RoomBooking
}
