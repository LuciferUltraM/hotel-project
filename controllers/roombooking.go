package controllers

import (
	"github.com/astaxie/beego"
	"github.com/code-mobi/hotel/models"
)

type RoomBookingController struct {
	beego.Controller
}

type RoomBookingStatus struct {
	Booking string
	Confirm string
}

func NewRoomBookingStatus() *RoomBookingStatus {
	rbs := new(RoomBookingStatus)
	rbs.Booking = "Booking"
	rbs.Confirm = "Confirm"
	return rbs
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
		c.Data["RoomBookingStatus"] = NewRoomBookingStatus()
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
			roomBooking.ConfirmBooking(c.GetString("Firstname"), c.GetString("Lastname"), c.GetString("CardID"))
			c.Data["RoomBooking"] = roomBooking
			c.Data["RoomBookingStatus"] = NewRoomBookingStatus()
		} else if action == "cancel" {
			hotel.DeleteRoomBooking(id)
			c.List()
		}
	} else {
		c.List()
	}
}
