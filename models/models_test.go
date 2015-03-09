package models

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ModelsTestSuite struct {
	suite.Suite
	CustomerName string
	CardID       string
	Tel          string
	Nationality  string
	// BirthDate    time.Time

	NumberOfCustomer int
}

func (suite *ModelsTestSuite) SetupTest() {
	suite.CustomerName = "Nat Ton"
	suite.Tel = "0895555555"
	suite.Nationality = "Thai"
	// suite.BirthDate = time.Date(1987, time.May, 13, 0, 0, 0, 0, time.UTC)

	suite.NumberOfCustomer = 2
}

func (suite *ModelsTestSuite) MockHotelSystem() *HotelSystem {
	roomTypes := suite.MockRoomTypes()
	rooms := suite.MockRooms(roomTypes)
	return &HotelSystem{
		RoomTypes: roomTypes,
		Rooms:     rooms,
	}
}

func (suite *ModelsTestSuite) MockRoomTypes() []*RoomType {
	return []*RoomType{
		&RoomType{"Superior Rooms", 3000, "Hiso Superior"},
		&RoomType{"Excusive Rooms", 4000, "Hiso Excusive"},
		&RoomType{"Jacuzzi Room", 5000, "Hiso Jacuzzi"},
	}
}

func (suite *ModelsTestSuite) MockRooms(roomTypes []*RoomType) []*Room {
	return []*Room{
		&Room{"101", "1", roomTypes[0]},
		&Room{"201", "2", roomTypes[1]},
		&Room{"301", "3", roomTypes[2]},
	}
}

func (suite *ModelsTestSuite) MockRoomBooking() *RoomBooking {
	return &RoomBooking{
		NumberOfCustomer: suite.NumberOfCustomer,
	}
}

func (suite *ModelsTestSuite) TestHotelSystem() {
	hotelSystem := suite.MockHotelSystem()
	suite.Len(hotelSystem.RoomTypes, 3)
	suite.Len(hotelSystem.Rooms, 3)
	suite.Equal(hotelSystem.Rooms[0].RoomType, hotelSystem.RoomTypes[0])
}

func (suite *ModelsTestSuite) TestRoomType() {
	hotelSystem := suite.MockHotelSystem()
	for _, roomType := range hotelSystem.RoomTypes {
		suite.NotEmpty(roomType.Name)
		suite.NotEqual(0, roomType.Rate)
		suite.NotEmpty(roomType.Detail)
	}
}

func (suite *ModelsTestSuite) TestRoom() {
	hotelSystem := suite.MockHotelSystem()
	for _, room := range hotelSystem.Rooms {
		suite.NotEmpty(room.RoomNo)
		suite.NotEmpty(room.Floor)
		suite.NotNil(room.RoomType)
	}
}

func (suite *ModelsTestSuite) TestRoomBooking() {
	rb := suite.MockRoomBooking()
	suite.Equal(suite.NumberOfCustomer, rb.NumberOfCustomer, "they should be equal")
}

func TestModelsTestSuite(t *testing.T) {
	suite.Run(t, new(ModelsTestSuite))
}
