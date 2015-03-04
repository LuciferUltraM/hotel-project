package models

import "time"

type RoomBookingItem struct {
	ID              int
	BookingItemDate time.Time
	Status          string
}
