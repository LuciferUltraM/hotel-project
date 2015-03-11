package controllers

import (
	"github.com/astaxie/beego"
	"github.com/code-mobi/hotel/models"
)

type RoomBookingController struct {
	beego.Controller
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
	c.Data["RoomBooking"] = hotel.FindRoomBooking(id)

}

func (c *RoomBookingController) Update() {
	c.TplNames = "roombookings/booking.tpl"

	id := c.Ctx.Input.Param(":id")
	hotel := models.GetInstance()
	c.Data["RoomBooking"] = hotel.FindRoomBooking(id)

}
