package controllers

import (
	"errors"
	"fmt"

	"github.com/code-mobi/hotel/models"
)

type RoomBookingController struct {
	BaseController
}

func (c *RoomBookingController) List() {
	c.GetUserLogin()
	c.TplNames = "roombookings/index.tpl"

	hotelSystem := models.GetInstance()

	c.Data["RoomBookings"] = hotelSystem.RoomBookings
}

func (c *RoomBookingController) Show() {
	user := c.GetUserLogin()
	id := c.Ctx.Input.Param(":id")
	hotelSystem := models.GetInstance()
	roomBooking := hotelSystem.FindRoomBooking(id)
	if roomBooking != nil {
		if roomBooking.Status == "New" {
			if user != nil {
				c.TplNames = "roombookings/booking.tpl"
			} else {
				c.TplNames = "roombookings/show.tpl"
			}
			c.Data["RoomBooking"] = roomBooking
		} else if roomBooking.Status == "Success" {
			c.NewPayment()
		} else {
			c.Redirect("/", 302)
		}
	} else {
		c.Redirect("/", 302)
	}
}

func (c *RoomBookingController) Update() {
	c.GetUserLogin()
	id := c.Ctx.Input.Param(":id")
	action := c.GetString("action")
	hotelSystem := models.GetInstance()
	roomBooking := hotelSystem.FindRoomBooking(id)
	if roomBooking != nil {
		if action == "confirm" {
			c.TplNames = "roombookings/booking.tpl"
			roomBooking.ConfirmBooking(c.GetString("FirstName"), c.GetString("LastName"), c.GetString("CardID"), c.GetString("ContactNo"))
			redirectUrl := fmt.Sprintf("/roombooking/%s/payment", roomBooking.RoomBookingNo)
			c.Redirect(redirectUrl, 302)
		} else if action == "cancel" {
			hotelSystem.DeleteRoomBooking(id)
			c.Redirect("/", 302)
		}
	} else {
		c.Redirect("/roombooking", 302)
	}
}

func (c *RoomBookingController) GetShow() (hotelSystem *models.HotelSystem, roomBooking *models.RoomBooking, err error) {
	c.GetUserLogin()
	c.TplNames = "roombookings/show.tpl"
	id := c.Ctx.Input.Param(":id")
	hotelSystem = models.GetInstance()
	roomBooking = hotelSystem.FindRoomBooking(id)
	if roomBooking != nil {
		c.Data["RoomBooking"] = roomBooking
		c.Data["RoomBookingStatus"] = hotelSystem.GetRoomBookingStatus(roomBooking)
	} else {
		c.Redirect("/", 302)
		err = errors.New("Room Booking Not Found")
	}
	return
}

func (c *RoomBookingController) NewPayment() {
	_, _, err := c.GetShow()
	if err != nil {
		return
	}
}

func (c *RoomBookingController) SavePayment() {
	c.GetUserLogin()
	id := c.Ctx.Input.Param(":id")
	hotelSystem := models.GetInstance()
	paymentOption := c.GetString("PaymentOption")
	if paymentOption != "" {
		receipt := hotelSystem.PaymentRoomBooking(id, paymentOption)
		url := fmt.Sprintf("/receipt/%s", receipt.ReceiptNo)
		c.Redirect(url, 302)
	} else {
		c.Redirect("/", 302)
	}
}

func (c *RoomBookingController) CheckIn() {
	hotelSystem, roomBooking, err := c.GetShow()
	if err != nil {
		return
	}
	hotelSystem.CheckIn(roomBooking)
	c.Data["RoomBookingStatus"] = hotelSystem.GetRoomBookingStatus(roomBooking)
}

func (c *RoomBookingController) CheckOut() {
	hotelSystem, _, err := c.GetShow()
	if err != nil {
		return
	}
	c.TplNames = "roombookings/checkout.tpl"
	c.Data["Equipments"] = hotelSystem.Equipments
}

func (c *RoomBookingController) SaveCheckOut() {
	c.GetUserLogin()
	id := c.Ctx.Input.Param(":id")
	hotelSystem := models.GetInstance()
	roomBooking := hotelSystem.FindRoomBooking(id)
	if roomBooking != nil {
		c.Data["RoomBooking"] = roomBooking
	} else {
		c.Redirect("/", 302)
	}
}
