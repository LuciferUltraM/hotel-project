package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ModelsTestSuite struct {
	suite.Suite
	hotelSystem  *HotelSystem
	CheckInDate  time.Time
	CheckOutDate time.Time
}

func TestModelsTestSuite(t *testing.T) {
	suite.Run(t, new(ModelsTestSuite))
}

func (suite *ModelsTestSuite) SetupTest() {
	suite.hotelSystem = suite.MockHotelSystem()
	suite.CheckInDate = time.Date(2015, 3, 15, 0, 0, 0, 0, time.UTC)
	suite.CheckOutDate = time.Date(2015, 3, 16, 0, 0, 0, 0, time.UTC)
}

func (suite *ModelsTestSuite) MockHotelSystem() *HotelSystem {
	hotel := new(HotelSystem)
	hotel.InitInstance()
	return hotel
}

func (suite *ModelsTestSuite) TestHotelSystem() {
	suite.Len(suite.hotelSystem.RoomTypes, 3)
	suite.Len(suite.hotelSystem.Rooms, 25)
	suite.Equal(suite.hotelSystem.Rooms["101"].RoomType, suite.hotelSystem.RoomTypes[0])
}

func (suite *ModelsTestSuite) TestReceptionist() *Receptionist {
	receptionist := suite.hotelSystem.Receptionists["1234"]
	suite.Equal(receptionist.EmployeeNo, 1234)
	suite.Equal(receptionist.FirstName, "Natt")
	suite.Equal(receptionist.LastName, "Ton")
	suite.Equal(receptionist.Gender, "M")
	suite.Equal(receptionist.BirthDate, time.Date(1987, 5, 13, 0, 0, 0, 0, time.UTC))
	suite.Equal(receptionist.UserName, "1234")
	suite.Equal(receptionist.Password, "4321")
	return receptionist
}

func (suite *ModelsTestSuite) TestRoomType() {
	for _, roomType := range suite.hotelSystem.RoomTypes {
		suite.NotEmpty(roomType.Name)
		suite.NotEqual(0, roomType.Rate)
		suite.NotEmpty(roomType.Detail)
	}
}

func (suite *ModelsTestSuite) TestRoom() {
	for _, room := range suite.hotelSystem.Rooms {
		suite.NotEmpty(room.RoomNo)
		suite.NotEmpty(room.Floor)
		suite.NotNil(room.RoomType)
	}
}

func (suite *ModelsTestSuite) TestFindReceptionist() {
	receptionist := suite.hotelSystem.FindReceptionist("1234")
	suite.Equal(receptionist.EmployeeNo, 1234)
	suite.Equal(receptionist.FirstName, "Natt")
	suite.Equal(receptionist.LastName, "Ton")
	suite.Equal(receptionist.Gender, "M")
	suite.Equal(receptionist.BirthDate, time.Date(1987, 5, 13, 0, 0, 0, 0, time.UTC))
	suite.Equal(receptionist.UserName, "1234")
	suite.Equal(receptionist.Password, "4321")
}

func (suite *ModelsTestSuite) TestFindOptionRate() {
	optionRate := suite.hotelSystem.FindOptionRate("extra_bed")
	suite.Equal(optionRate.GetName(), "extra_bed")
	suite.Equal(optionRate.GetRate(), 1200)
	optionRate = suite.hotelSystem.FindOptionRate("vat_rate")
	suite.Equal(optionRate.GetName(), "vat_rate")
	suite.Equal(optionRate.GetRate(), 7)
}

func (suite *ModelsTestSuite) TestGetAvailableRoom() {
	hotelSystem := suite.MockHotelSystem()
	rooms := hotelSystem.GetAvailableRoom("2015-03-15", "2015-03-16")
	suite.NotNil(rooms)
	suite.Len(rooms, 23)
}

func (suite *ModelsTestSuite) TestBooking() *RoomBooking {
	receptionist := "1234"
	selectedRooms := []string{"101", "301"}
	extraBeds := []bool{true, false}
	checkIn := "2015-03-15"
	checkOut := "2015-03-16"
	roomBooking := suite.hotelSystem.ReserveRoom(receptionist, selectedRooms, extraBeds, checkIn, checkOut)
	suite.Equal(roomBooking.CreatedBy, suite.hotelSystem.FindReceptionist(receptionist))
	suite.Len(roomBooking.Rooms, 2)
	suite.Len(roomBooking.ExtraBeds, 2)
	suite.Equal(roomBooking.CheckInDate, suite.CheckInDate)
	suite.Equal(roomBooking.CheckOutDate, suite.CheckOutDate)
	suite.Equal(roomBooking.NightAmount, 1)
	amount := float32(8200 * 1)
	suite.Equal(roomBooking.GetAmount(), amount)
	suite.Equal(roomBooking.GetVat(), amount*7/100)
	suite.Equal(roomBooking.GetGrandTotal(), amount+roomBooking.GetVat())
	suite.Equal(roomBooking.Status, "New")
	return roomBooking
}

func (suite *ModelsTestSuite) TestConfirmBooking() (roomBooking *RoomBooking) {
	roomBooking = suite.TestBooking()
	roomBooking.ConfirmBooking("Natt", "Ton", "1709999999999", "0895555555")
	suite.Equal(roomBooking.FirstName, "Natt")
	suite.Equal(roomBooking.LastName, "Ton")
	suite.Equal(roomBooking.CardID, "1709999999999")
	suite.Equal(roomBooking.ContactNo, "0895555555")
	return
}

func (suite *ModelsTestSuite) TestPayForRoomBooking() {
	roomBooking := suite.TestConfirmBooking()
	suite.NotNil(roomBooking)
	var receipt *Receipt
	receipt = suite.hotelSystem.PayForRoomBooking(roomBooking.RoomBookingNo, "Cash")
	suite.NotNil(receipt)
	suite.Equal(receipt.RoomBooking, roomBooking)
	suite.NotEmpty(receipt.ReceiptNo)
	suite.NotNil(receipt.ReceiptDate)
	suite.Equal(receipt.Status, "Success")
	suite.Equal(receipt.Amount, roomBooking.GrandTotal)
	suite.Equal(receipt.Type, "Cash")
	suite.Equal(roomBooking.Status, "Success")
}

func (suite *ModelsTestSuite) TestRoomBookingDiffDay() {
	rb := &RoomBooking{}
	diffDay, err := rb.diffDay(suite.CheckInDate, suite.CheckOutDate)
	suite.Equal(diffDay, 1)
	suite.Nil(err)
}

func (suite *ModelsTestSuite) TestRoomBookingDiffDayError() {
	rb := &RoomBooking{}
	diffDay, err := rb.diffDay(suite.CheckOutDate, suite.CheckInDate)
	suite.Equal(diffDay, 0)
	suite.NotNil(err)
}

func (suite *ModelsTestSuite) TestStringToDate() {
	suite.Equal(suite.hotelSystem.stringToDate("2015-03-01"), time.Date(2015, 3, 1, 0, 0, 0, 0, time.UTC))
}
