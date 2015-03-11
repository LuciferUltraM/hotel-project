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
	RoomBookings map[string]*RoomBooking
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
	hotel.RoomBookings = make(map[string]*RoomBooking)
	hotel.RoomTypes = hotel.InitSampleRoomTypes()
	hotel.Rooms = hotel.InitSampleRooms(hotel.RoomTypes)
	hotel.OptionRates = hotel.InitSampleOptionRate()
	hotel.InitSampleRoomBooking()
}

func (hotel *HotelSystem) InitSampleRoomTypes() []*RoomType {
	return []*RoomType{
		&RoomType{0, "Superior Rooms", 3000, "Hiso Superior"},
		&RoomType{1, "Excusive Rooms", 4000, "Hiso Excusive"},
		&RoomType{2, "Jacuzzi Room", 5000, "Hiso Jacuzzi"},
	}
}

func (hotel *HotelSystem) InitSampleRooms(roomTypes []*RoomType) map[string]*Room {
	rooms := make(map[string]*Room)
	var roomType *RoomType
	for f := 1; f < 6; f++ {
		if f < 3 {
			roomType = roomTypes[0]
		} else if f < 5 {
			roomType = roomTypes[1]
		} else {
			roomType = roomTypes[2]
		}

		for n := 1; n < 11; n++ {
			roomNo := fmt.Sprintf("%d%02d", f, n)
			rooms[roomNo] = &Room{roomNo, strconv.Itoa(f), "OK", roomType}
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

func (hotel *HotelSystem) InitSampleRoomBooking() {
	selectedRooms := []string{"102", "203"}
	extraBeds := []bool{false, false}
	checkInDate := time.Date(2015, 3, 13, 0, 0, 0, 0, time.UTC)
	checkOutDate := time.Date(2015, 3, 15, 0, 0, 0, 0, time.UTC)
	hotel.ReserveRoom(selectedRooms, extraBeds, checkInDate, checkOutDate)

	selectedRooms2 := []string{"205", "504"}
	extraBeds2 := []bool{false, false}
	checkInDate2 := time.Date(2015, 3, 13, 0, 0, 0, 0, time.UTC)
	checkOutDate2 := time.Date(2015, 3, 17, 0, 0, 0, 0, time.UTC)
	hotel.ReserveRoom(selectedRooms2, extraBeds2, checkInDate2, checkOutDate2)
}

func (hotel *HotelSystem) FindRoom(roomNo string) *Room {
	return hotel.Rooms[roomNo]
}

func (hotel *HotelSystem) FindOptionRate(optionName string) *OptionRate {
	return hotel.OptionRates[optionName]
}

func (hotel *HotelSystem) GetRoomAvailable(checkInDate time.Time, checkOutDate time.Time) (rooms map[string]*Room) {
	rooms = hotel.cloneRooms()
	for _, roomBooking := range hotel.RoomBookings {
		for checkingDate := roomBooking.CheckInDate; checkingDate.Before(checkOutDate); checkingDate = checkInDate.AddDate(0, 0, 1) {
			bookedDate := roomBooking.CheckInDate
			for night := 0; night < roomBooking.NightAmount; bookedDate = roomBooking.CheckInDate.AddDate(0, 0, night) {
				if checkingDate == bookedDate {
					hotel.deleteAvailableRooms(&rooms, roomBooking.Rooms)
				}
				night++
			}
		}
	}
	return
}

func (hotel *HotelSystem) ReserveRoom(
	selectedRooms []string,
	extraBeds []bool,
	checkInDate time.Time,
	checkOutDate time.Time) *RoomBooking {
	extraBedRate := hotel.FindOptionRate("extra_bed")
	vatRate := hotel.FindOptionRate("vat_rate")
	rooms := []*Room{}
	for _, roomNo := range selectedRooms {
		rooms = append(rooms, hotel.FindRoom(roomNo))
	}
	roomBooking := &RoomBooking{}
	err := roomBooking.ReserveRoom(extraBedRate.GetRate(), vatRate.GetRate(), rooms, extraBeds, checkInDate, checkOutDate)
	if err != nil {
		panic(err)
	}

	hotel.RoomBookings[roomBooking.RoomBookingNo] = roomBooking
	return roomBooking
}

func (hotel *HotelSystem) cloneRooms() map[string]*Room {
	rooms := make(map[string]*Room)
	for k, v := range hotel.Rooms {
		rooms[k] = v
	}
	return rooms
}

func (hotel *HotelSystem) deleteAvailableRooms(rooms *map[string]*Room, deleteRooms []*Room) {
	for _, room := range deleteRooms {
		delete(*rooms, room.RoomNo)
	}
}
