package controllers

import (
	"github.com/astaxie/beego"
	"github.com/code-mobi/hotel/models"
)

type RoomController struct {
	beego.Controller
}

func (c *RoomController) Get() {
	hotel := models.GetInstance()

	c.TplNames = "rooms/index.tpl"

	c.Data["Rooms"] = hotel.Rooms
}
