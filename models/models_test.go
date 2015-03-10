package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ModelsTestSuite struct {
	suite.Suite
	CustomerName     string
	CardID           string
	Tel              string
	Nationality      string
	CheckInDate      time.Time
	CheckOutDate     time.Time
	NumberOfCustomer int
}

func (suite *ModelsTestSuite) SetupTest() {
	suite.CustomerName = "Nat Ton"
	suite.CardID = "1709912345678"
	suite.Tel = "0895555555"
	suite.Nationality = "Thai"
	suite.CheckInDate = time.Date(2015, 3, 10, 0, 0, 0, 0, time.UTC)
	suite.CheckOutDate = time.Date(2015, 3, 12, 0, 0, 0, 0, time.UTC)

	suite.NumberOfCustomer = 2
}

func (suite *ModelsTestSuite) MockHotelSystem() *HotelSystem {
	hotel := new(HotelSystem)
	hotel.InitInstance()
	return hotel
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

func (suite *ModelsTestSuite) TestHotelSystem() {
	hotelSystem := suite.MockHotelSystem()
	suite.Len(hotelSystem.RoomTypes, 3)
	suite.Len(hotelSystem.Rooms, 50)
	suite.Equal(hotelSystem.Rooms["101"].RoomType, hotelSystem.RoomTypes[0])
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

func (suite *ModelsTestSuite) TestFindOptionRate() {
	hotelSystem := suite.MockHotelSystem()
	optionRate := hotelSystem.FindOptionRate("extra_bed")
	suite.Equal(optionRate.GetName(), "extra_bed")
	suite.Equal(optionRate.GetRate(), 1200)
}

func (suite *ModelsTestSuite) TestBooking() {
	var hotel *HotelSystem
	hotel = suite.MockHotelSystem()
	rooms := []*Room{hotel.Rooms["101"], hotel.Rooms["102"]}
	extraBeds := []bool{true, false}
	roomBooking := hotel.ReserveRoom(rooms, extraBeds, suite.CheckInDate, suite.CheckOutDate)

	suite.Equal(roomBooking.CheckInDate, suite.CheckInDate)
	suite.Equal(roomBooking.CheckOutDate, suite.CheckOutDate)

	amount := float32(7200)
	suite.Equal(roomBooking.GetAmount(), amount)
	suite.Equal(roomBooking.GetVat(), amount*7/100)
	suite.Equal(roomBooking.GetGrandTotal(), amount+roomBooking.GetVat())
}

func (suite *ModelsTestSuite) TestRoomBookingCountDay() {
	rb := &RoomBooking{}
	countDay, err := rb.countDay(suite.CheckInDate, suite.CheckOutDate)
	suite.Equal(countDay, 2)
	suite.Nil(err)
}

func (suite *ModelsTestSuite) TestRoomBookingCountDayError() {
	rb := &RoomBooking{}
	countDay, err := rb.countDay(suite.CheckOutDate, suite.CheckInDate)
	suite.Equal(countDay, 0)
	suite.NotNil(err)
}

func TestModelsTestSuite(t *testing.T) {
	suite.Run(t, new(ModelsTestSuite))
}
