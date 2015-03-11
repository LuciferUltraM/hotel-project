package controllers

import (
	"github.com/astaxie/beego"
	"github.com/code-mobi/hotel/models"
)

type RoomBookingController struct {
	beego.Controller
}

func (c *RoomBookingController) Get() {
	hotel := models.GetInstance()

	c.TplNames = "roombookings/index.tpl"

	c.Data["RoomBookings"] = hotel.RoomBookings
}
