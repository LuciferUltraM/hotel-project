package controllers

import "github.com/code-mobi/hotel/models"

type EquipmentController struct {
	BaseController
}

func (c *EquipmentController) Get() {
	c.GetUserLogin()
	c.TplNames = "equipments/index.tpl"
	hotelSystem := models.GetInstance()
	c.Data["Equipments"] = hotelSystem.Equipments
}
