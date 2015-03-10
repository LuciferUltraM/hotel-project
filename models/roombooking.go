package models

import (
	"errors"
	"time"
)

type RoomBooking struct {
	RoomBookingNo    int32
	CustomerName     string
	CardID           string
	ContactNo        string
	CheckInDate      time.Time
	CheckOutDate     time.Time
	BookingDate      time.Time
	NumberOfCustomer int
	Status           string
	Rooms            []*Room
	ExtraBeds        []bool
	Receipt          *Receipt
	Refund           *Refund
	amount           float32
	vat              float32
	grandTotal       float32
}

func (rb *RoomBooking) GetAmount() float32 {
	return rb.amount
}

func (rb *RoomBooking) GetVat() float32 {
	return rb.vat
}

func (rb *RoomBooking) GetGrandTotal() float32 {
	return rb.grandTotal
}

func (rb *RoomBooking) ReserveRoom(
	extraBedRate float32,
	vatRate float32,
	rooms []*Room,
	extraBeds []bool,
	checkInDate time.Time,
	checkOutDate time.Time) {

	rb.CheckInDate = checkInDate
	rb.CheckOutDate = checkOutDate

	var totalPrice float32

	for _, room := range rooms {
		totalPrice += room.RoomType.Rate
	}

	for _, extraBed := range extraBeds {
		if extraBed {
			totalPrice += extraBedRate
		}
	}

	rb.setAmount(totalPrice, vatRate)
}

func (rb *RoomBooking) countDay(checkInDate time.Time, checkOutDate time.Time) (countDay int, err error) {
	if checkInDate.Before(checkOutDate) {
		for d := checkInDate; d.Before(checkOutDate); d = d.AddDate(0, 0, 1) {
			countDay++
		}
	} else {
		err = errors.New("checkInDate should before checkOutDate")
	}
	return
}

func (rb *RoomBooking) setAmount(amount float32, vatRate float32) {
	rb.amount = amount
	rb.calculateVat(vatRate)
	rb.sumGrandTotal()
}

func (rb *RoomBooking) calculateVat(vatRate float32) {
	rb.vat = rb.amount * vatRate / 100
}

func (rb *RoomBooking) sumGrandTotal() {
	rb.grandTotal = rb.amount + rb.vat
}
