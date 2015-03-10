package models

import (
	"fmt"
	"strconv"
	"time"
)

type HotelSystem struct {
	RoomTypes    []*RoomType
	Rooms        map[string]*Room
	OptionRates  map[string]*OptionRate
	RoomBookings []*RoomBooking
}

var instantiated *HotelSystem = nil

func GetInstance() *HotelSystem {
	if instantiated == nil {
		instantiated = new(HotelSystem)
		instantiated.InitInstance()
	}
	return instantiated
}

func (hotel *HotelSystem) InitInstance() {
	hotel.RoomTypes = hotel.InitSampleRoomTypes()
	hotel.Rooms = hotel.InitSampleRooms(hotel.RoomTypes)
	hotel.OptionRates = hotel.InitSampleOptionRate()
	hotel.RoomBookings = []*RoomBooking{}
}

func (hotel *HotelSystem) InitSampleRoomTypes() []*RoomType {
	return []*RoomType{
		&RoomType{"Superior Rooms", 3000, "Hiso Superior"},
		&RoomType{"Excusive Rooms", 4000, "Hiso Excusive"},
		&RoomType{"Jacuzzi Room", 5000, "Hiso Jacuzzi"},
	}
}

func (hotel *HotelSystem) InitSampleRooms(roomTypes []*RoomType) map[string]*Room {
	// roomTypesLength := len(roomTypes)
	rooms := make(map[string]*Room)
	for f := 1; f < 6; f++ {
		for n := 1; n < 11; n++ {
			roomNo := fmt.Sprintf("%d%02d", f, n)
			rooms[roomNo] = &Room{roomNo, strconv.Itoa(f), roomTypes[0]}
		}
	}
	return rooms
}

func (hotel *HotelSystem) InitSampleOptionRate() map[string]*OptionRate {
	optionRates := make(map[string]*OptionRate)
	optionRates["extra_bed"] = &OptionRate{"extra_bed", 1200}
	optionRates["vat_rate"] = &OptionRate{"vat_rate", 7}
	return optionRates
}

func (hotel *HotelSystem) FindOptionRate(optionName string) *OptionRate {
	return hotel.OptionRates[optionName]
}

func (hotel *HotelSystem) GetRoomCheckIn(checkInDate time.Time, checkOutDate time.Time) {

}

func (hotel *HotelSystem) ReserveRoom(
	rooms []*Room,
	extraBeds []bool,
	checkInDate time.Time,
	checkOutDate time.Time) *RoomBooking {
	extraBedRate := hotel.FindOptionRate("extra_bed")
	vatRate := hotel.FindOptionRate("vat_rate")
	roomBooking := &RoomBooking{}
	roomBooking.ReserveRoom(extraBedRate.GetRate(), vatRate.GetRate(), rooms, extraBeds, checkInDate, checkOutDate)

	return roomBooking
}
