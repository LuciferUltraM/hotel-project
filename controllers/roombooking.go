package controllers

import "github.com/code-mobi/hotel/models"

type RoomBookingController struct {
	BaseController
}

type RoomBookingStatus struct {
	Booking string
	Confirm string
}

func (c *RoomBookingController) List() {
	c.TplNames = "roombookings/index.tpl"

	hotel := models.GetInstance()

	c.Data["RoomBookings"] = hotel.RoomBookings
}

func (c *RoomBookingController) Show() {
	c.TplNames = "roombookings/booking.tpl"
	id := c.Ctx.Input.Param(":id")
	hotel := models.GetInstance()
	roomBooking := hotel.FindRoomBooking(id)
	if roomBooking != nil {
		c.Data["RoomBooking"] = roomBooking
	} else {
		c.List()
	}
}

func (c *RoomBookingController) Update() {
	id := c.Ctx.Input.Param(":id")
	action := c.GetString("action")
	hotel := models.GetInstance()
	roomBooking := hotel.FindRoomBooking(id)
	if roomBooking != nil {
		if action == "confirm" {
			c.TplNames = "roombookings/booking.tpl"
			roomBooking.ConfirmBooking(c.GetString("FirstName"), c.GetString("LastName"), c.GetString("CardID"), c.GetString("ContactNo"))
			c.Data["RoomBooking"] = roomBooking
		} else if action == "cancel" {
			hotel.DeleteRoomBooking(id)
			c.Redirect("/", 302)
		}
	} else {
		c.Redirect("/roombooking", 302)
	}
}
