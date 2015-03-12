package controllers

import "github.com/code-mobi/hotel/models"

type RoomController struct {
	BaseController
}

func (c *RoomController) Get() {
	hotel := models.GetInstance()

	c.TplNames = "rooms/index.tpl"

	c.Data["Rooms"] = hotel.Rooms
}
