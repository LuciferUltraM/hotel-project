package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ModelsTestSuite struct {
	suite.Suite
	FirstName   string
	LastName    string
	Tel         string
	Nationality string
	BirthDate   time.Time

	NumberOfCustomer int
}

func (suite *ModelsTestSuite) SetupTest() {
	suite.FirstName = "Nat"
	suite.LastName = "Ton"
	suite.Tel = "0895555555"
	suite.Nationality = "Thai"
	suite.BirthDate = time.Date(1987, time.May, 13, 0, 0, 0, 0, time.UTC)

	suite.NumberOfCustomer = 2
}

func (suite *ModelsTestSuite) MockCustomer() *Customer {
	return &Customer{
		FirstName:   suite.FirstName,
		LastName:    suite.LastName,
		Tel:         suite.Tel,
		Nationality: suite.Nationality,
		BirthDate:   suite.BirthDate,
	}
}

func (suite *ModelsTestSuite) MockRoomBooking(customer *Customer) *RoomBooking {
	return &RoomBooking{
		Customer:         customer,
		NumberOfCustomer: suite.NumberOfCustomer,
	}
}

func (suite *ModelsTestSuite) TestCustomer() {
	c := suite.MockCustomer()
	suite.Equal(c.FirstName, suite.FirstName, "they should be equal")
	suite.Equal(c.LastName, suite.LastName, "they should be equal")
	suite.Equal(c.Tel, suite.Tel, "they should be equal")
	suite.Equal(c.Nationality, suite.Nationality, "they should be equal")
	suite.Equal(c.BirthDate, suite.BirthDate, "they should be equal")
}

func (suite *ModelsTestSuite) TestRoomBooking() {
	customer := suite.MockCustomer()
	rb := suite.MockRoomBooking(customer)
	suite.Equal(customer, rb.Customer, "they should be equal")
	suite.Equal(suite.NumberOfCustomer, rb.NumberOfCustomer, "they should be equal")
}

func TestModelsTestSuite(t *testing.T) {
	suite.Run(t, new(ModelsTestSuite))
}
