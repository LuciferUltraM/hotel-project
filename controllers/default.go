package controllers

import (
	"github.com/astaxie/beego"
	"github.com/code-mobi/hotel/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplNames = "index.tpl"

	hotel := models.GetInstance()
	c.Data["Rooms"] = hotel.Rooms
}
