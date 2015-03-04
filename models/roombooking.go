package models

import "time"

type RoomBooking struct {
	RoomBookingNo    int
	Customer         *Customer
	CheckInDate      time.Time
	CheckOutDate     time.Time
	BookingDate      time.Time
	NumberOfCustomer int
	Status           string
	Receipt          *Receipt
	Refund           *Refund
	RoomBookingItems []*RoomBookingItem
}
