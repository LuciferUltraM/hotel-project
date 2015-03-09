package models

import (
	"fmt"
	"math/rand"
	"strconv"
)

type HotelSystem struct {
	RoomTypes []*RoomType
	Rooms     []*Room
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
}

func (hotel *HotelSystem) InitSampleRoomTypes() []*RoomType {
	return []*RoomType{
		&RoomType{"Superior Rooms", 3000, "Hiso Superior"},
		&RoomType{"Excusive Rooms", 4000, "Hiso Excusive"},
		&RoomType{"Jacuzzi Room", 5000, "Hiso Jacuzzi"},
	}
}

func (hotel *HotelSystem) InitSampleRooms(roomTypes []*RoomType) []*Room {
	roomTypesLength := len(roomTypes)
	rooms := []*Room{}
	for f := 1; f < 6; f++ {
		for n := 1; n < 11; n++ {
			rooms = append(rooms, &Room{fmt.Sprintf("%d%02d", f, n), strconv.Itoa(f), roomTypes[rand.Intn(roomTypesLength)]})
		}
	}
	return rooms
}
