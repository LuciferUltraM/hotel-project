package models

import "time"

type RoomBooking struct {
	RoomBookingNo    int
	CheckInDate      time.Time
	CheckOutDate     time.Time
	BookingDate      time.Time
	NumberOfCustomer int
	Status           string
	Receipt          *Receipt
	Refund           *Refund
}

func (rm *RoomBooking) GetRoomCheckIn(checkInDate time.Time) {

}
