package controllers

import (
	"github.com/astaxie/beego"
	"github.com/code-mobi/hotel/models"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) GetUserLogin() (user *models.Receptionist) {
	username := c.GetSession("username")
	if username != nil {
		hotelSystem := models.GetInstance()
		user = hotelSystem.FindReceptionist(username.(string))
		c.Data["User"] = user
	}
	return
}
