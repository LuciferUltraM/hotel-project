package models

import (
	"errors"
	"fmt"
	"time"
)

type RoomBooking struct {
	RoomBookingNo  string
	CustomerName   string
	CardID         string
	ContactNo      string
	CheckInDate    time.Time
	CheckOutDate   time.Time
	BookingDate    time.Time
	NightAmount    int
	CustomerAmount int
	Status         string
	Rooms          []*Room
	ExtraBeds      []bool
	Receipt        *Receipt
	Refund         *Refund
	Amount         float32
	Vat            float32
	GrandTotal     float32
}

func (rb *RoomBooking) GetAmount() float32 {
	return rb.Amount
}

func (rb *RoomBooking) GetVat() float32 {
	return rb.Vat
}

func (rb *RoomBooking) GetGrandTotal() float32 {
	return rb.GrandTotal
}

func (rb *RoomBooking) setRoomBookingNo(t time.Time) {
	rb.RoomBookingNo = fmt.Sprintf("%d", t.UnixNano())
}

func (rb *RoomBooking) ReserveRoom(
	extraBedRate float32,
	vatRate float32,
	rooms []*Room,
	extraBeds []bool,
	checkInDate time.Time,
	checkOutDate time.Time) error {

	rb.Rooms = rooms
	rb.BookingDate = time.Now()
	rb.setRoomBookingNo(rb.BookingDate)
	rb.CheckInDate = checkInDate
	rb.CheckOutDate = checkOutDate

	diffDay, err := rb.diffDay(rb.CheckInDate, rb.CheckOutDate)
	if err != nil {
		return err
	}
	rb.NightAmount = diffDay

	var totalPrice float32
	for _, room := range rooms {
		totalPrice += room.RoomType.Rate
	}

	for _, extraBed := range extraBeds {
		if extraBed {
			totalPrice += extraBedRate
		}
	}

	rb.setAmount(totalPrice*float32(diffDay), vatRate)
	return nil
}

func (rb *RoomBooking) diffDay(checkInDate time.Time, checkOutDate time.Time) (diffDay int, err error) {
	if checkInDate.Before(checkOutDate) {
		for d := checkInDate; d.Before(checkOutDate); d = d.AddDate(0, 0, 1) {
			diffDay++
		}
	} else {
		err = errors.New("checkInDate should before checkOutDate")
	}
	return
}

func (rb *RoomBooking) setAmount(amount float32, vatRate float32) {
	rb.Amount = amount
	rb.calculateVat(vatRate)
	rb.sumGrandTotal()
}

func (rb *RoomBooking) calculateVat(vatRate float32) {
	rb.Vat = rb.Amount * vatRate / 100
}

func (rb *RoomBooking) sumGrandTotal() {
	rb.GrandTotal = rb.Amount + rb.Vat
}
