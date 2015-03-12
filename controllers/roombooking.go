package controllers

import (
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
	c.GetUserLogin()
	id := c.Ctx.Input.Param(":id")
	hotelSystem := models.GetInstance()
	roomBooking := hotelSystem.FindRoomBooking(id)
	if roomBooking != nil {
		if roomBooking.Status == "New" {
			c.TplNames = "roombookings/booking.tpl"
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

func (c *RoomBookingController) NewPayment() {
	c.GetUserLogin()
	c.TplNames = "roombookings/newpayment.tpl"
	id := c.Ctx.Input.Param(":id")
	hotelSystem := models.GetInstance()
	roomBooking := hotelSystem.FindRoomBooking(id)
	if roomBooking != nil {
		c.Data["RoomBooking"] = roomBooking
		c.Data["RoomBookingStatus"] = hotelSystem.GetRoomBookingStatus(roomBooking)
	} else {
		c.Redirect("/", 302)
	}
}

func (c *RoomBookingController) SavePayment() {
	c.GetUserLogin()
	id := c.Ctx.Input.Param(":id")
	hotelSystem := models.GetInstance()
	paymentOption := c.GetString("PaymentOption")
	if paymentOption != "" {
		hotelSystem.PayForRoomBooking(id, paymentOption)
		c.Data["Flash"] = "Payment Success"
		c.Redirect("/receipt", 302)
	} else {
		c.Redirect("/", 302)
	}
}

func (c *RoomBookingController) CheckIn() {
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
