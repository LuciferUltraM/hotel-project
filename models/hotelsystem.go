package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type HotelSystem struct {
	Receptionists map[string]*Receptionist
	RoomTypes     []*RoomType
	Rooms         map[string]*Room
	OptionRates   map[string]*OptionRate
	RoomBookings  map[string]*RoomBooking
	Receipts      map[string]*Receipt
	Equipments    map[string]*Equipment
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
	hotel.InitSampleReceptionist()
	hotel.RoomBookings = make(map[string]*RoomBooking)
	hotel.RoomTypes = hotel.InitSampleRoomTypes()
	hotel.Rooms = hotel.InitSampleRooms(hotel.RoomTypes)
	hotel.OptionRates = hotel.InitSampleOptionRate()
	hotel.InitSampleRoomBooking()
	hotel.Receipts = make(map[string]*Receipt)
}

func (hotel *HotelSystem) InitSampleReceptionist() {
	hotel.Receptionists = make(map[string]*Receptionist)
	receptionist := &Receptionist{}
	receptionist.EmployeeNo = 1234
	receptionist.FirstName = "Natt"
	receptionist.LastName = "Ton"
	receptionist.Gender = "M"
	receptionist.BirthDate = time.Date(1987, 5, 13, 0, 0, 0, 0, time.UTC)
	receptionist.UserName = "1234"
	receptionist.Password = "4321"
	hotel.Receptionists[receptionist.UserName] = receptionist
	return
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

		for n := 1; n < 6; n++ {
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
	receptionist := hotel.FindReceptionist("1234")
	selectedRooms := []string{"102", "203"}
	extraBeds := []bool{true, false}
	hotel.ReserveRoom(receptionist, selectedRooms, extraBeds, "2015-03-13", "2015-03-15")

	selectedRooms2 := []string{"205", "504"}
	extraBeds2 := []bool{false, false}
	hotel.ReserveRoom(receptionist, selectedRooms2, extraBeds2, "2015-03-13", "2015-03-17")
}

func (hotel *HotelSystem) FindReceptionist(username string) *Receptionist {
	return hotel.Receptionists[username]
}

func (hotel *HotelSystem) FindRoom(roomNo string) *Room {
	return hotel.Rooms[roomNo]
}

func (hotel *HotelSystem) FindOptionRate(optionName string) *OptionRate {
	return hotel.OptionRates[optionName]
}

func (hotel *HotelSystem) FindRoomBooking(roomBookingNo string) *RoomBooking {
	return hotel.RoomBookings[roomBookingNo]
}

func (hotel *HotelSystem) DeleteRoomBooking(roomBookingNo string) {
	delete(hotel.RoomBookings, roomBookingNo)
}

func (hotel *HotelSystem) GetAvailableRoom(checkIn string, checkOut string) (rooms map[string]*Room) {
	checkInDate := hotel.stringToDate(checkIn)
	checkOutDate := hotel.stringToDate(checkOut)
	rooms = hotel.cloneRooms()
	for _, roomBooking := range hotel.RoomBookings {
		isDeleted := false
		for checkingDate := checkInDate; checkingDate.Before(checkOutDate); checkingDate = checkingDate.AddDate(0, 0, 1) {
			if isDeleted {
				break
			}

			bookedDate := roomBooking.CheckInDate
			for night := 0; night < roomBooking.NightAmount; bookedDate = roomBooking.CheckInDate.AddDate(0, 0, night) {
				night++
				if checkingDate == bookedDate {
					hotel.deleteAvailableRooms(&rooms, roomBooking.Rooms)
					isDeleted = true
					break
				}
			}
		}
	}
	return
}

func (hotel *HotelSystem) ReserveRoom(
	receptionist *Receptionist,
	selectedRooms []string,
	extraBeds []bool,
	checkIn string,
	checkOut string) *RoomBooking {
	checkInDate := hotel.stringToDate(checkIn)
	checkOutDate := hotel.stringToDate(checkOut)
	extraBedRate := hotel.FindOptionRate("extra_bed")
	vatRate := hotel.FindOptionRate("vat_rate")
	rooms := []*Room{}
	for _, roomNo := range selectedRooms {
		rooms = append(rooms, hotel.FindRoom(roomNo))
	}
	roomBooking := &RoomBooking{}
	err := roomBooking.ReserveRoom(receptionist, extraBedRate.GetRate(), vatRate.GetRate(), rooms, extraBeds, checkInDate, checkOutDate)
	if err != nil {
		panic(err)
	}

	hotel.RoomBookings[roomBooking.RoomBookingNo] = roomBooking
	return roomBooking
}

func (hotel *HotelSystem) PayForRoomBooking(roomBookingNo string, paymentOption string) *Receipt {
	receipt := &Receipt{}
	now := time.Now()
	receipt.ReceiptNo = string(now.UnixNano())
	receipt.ReceiptDate = now
	roomBooking := hotel.FindRoomBooking(roomBookingNo)
	receipt.RoomBooking = roomBooking
	receipt.Type = paymentOption
	receipt.Amount = roomBooking.GrandTotal
	receipt.Status = "Success"
	roomBooking.Status = "Success"
	hotel.Receipts[receipt.ReceiptNo] = receipt
	return receipt
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

func (hotel *HotelSystem) stringToDate(dateStr string) time.Time {
	dates := strings.Split(dateStr, "-")
	year, _ := strconv.Atoi(dates[0])
	month, _ := strconv.Atoi(dates[1])
	day, _ := strconv.Atoi(dates[2])
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
