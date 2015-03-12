package controllers

import "github.com/code-mobi/hotel/models"

type RoomController struct {
	BaseController
}

func (c *RoomController) Get() {
	c.GetUserLogin()
	c.TplNames = "rooms/index.tpl"
	hotelSystem := models.GetInstance()
	c.Data["Rooms"] = hotelSystem.Rooms
}
