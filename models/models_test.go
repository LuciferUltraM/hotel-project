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

func TestModelsTestSuite(t *testing.T) {
	suite.Run(t, new(ModelsTestSuite))
}

func (suite *ModelsTestSuite) SetupTest() {
	suite.CheckInDate = time.Date(2015, 3, 15, 0, 0, 0, 0, time.UTC)
	suite.CheckOutDate = time.Date(2015, 3, 16, 0, 0, 0, 0, time.UTC)
}

func (suite *ModelsTestSuite) MockHotelSystem() *HotelSystem {
	hotel := new(HotelSystem)
	hotel.InitInstance()
	return hotel
}

func (suite *ModelsTestSuite) TestHotelSystem() {
	hotelSystem := suite.MockHotelSystem()
	suite.Len(hotelSystem.RoomTypes, 3)
	suite.Len(hotelSystem.Rooms, 25)
	suite.Equal(hotelSystem.Rooms["101"].RoomType, hotelSystem.RoomTypes[0])
}

func (suite *ModelsTestSuite) TestReceptionist() *Receptionist {
	hotelSystem := suite.MockHotelSystem()
	receptionist := hotelSystem.Receptionists["1234"]
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

func (suite *ModelsTestSuite) TestFindReceptionist() {
	hotelSystem := suite.MockHotelSystem()
	receptionist := hotelSystem.FindReceptionist("1234")
	suite.Equal(receptionist.EmployeeNo, 1234)
	suite.Equal(receptionist.FirstName, "Natt")
	suite.Equal(receptionist.LastName, "Ton")
	suite.Equal(receptionist.Gender, "M")
	suite.Equal(receptionist.BirthDate, time.Date(1987, 5, 13, 0, 0, 0, 0, time.UTC))
	suite.Equal(receptionist.UserName, "1234")
	suite.Equal(receptionist.Password, "4321")
}

func (suite *ModelsTestSuite) TestFindOptionRate() {
	hotelSystem := suite.MockHotelSystem()
	optionRate := hotelSystem.FindOptionRate("extra_bed")
	suite.Equal(optionRate.GetName(), "extra_bed")
	suite.Equal(optionRate.GetRate(), 1200)
	optionRate = hotelSystem.FindOptionRate("vat_rate")
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
	hotelSystem := suite.MockHotelSystem()
	receptionist := "1234"
	selectedRooms := []string{"101", "301"}
	extraBeds := []bool{true, false}
	checkIn := "2015-03-15"
	checkOut := "2015-03-16"
	roomBooking := hotelSystem.ReserveRoom(receptionist, selectedRooms, extraBeds, checkIn, checkOut)
	suite.Equal(roomBooking.CreatedBy, hotelSystem.FindReceptionist(receptionist))
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

func (suite *ModelsTestSuite) TestPaymentBooking() {
	// roomBooking := suite.TestConfirmBooking()
	// hotelSystem := suite.MockHotelSystem()

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
	hotelSystem := &HotelSystem{}
	suite.Equal(hotelSystem.stringToDate("2015-03-01"), time.Date(2015, 3, 1, 0, 0, 0, 0, time.UTC))
}
