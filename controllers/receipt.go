package controllers

import "github.com/code-mobi/hotel/models"

type ReceiptController struct {
	BaseController
}

func (c *ReceiptController) Get() {
	c.GetUserLogin()
	c.TplNames = "receipts/index.tpl"
	hotelSystem := models.GetInstance()
	c.Data["Receipts"] = hotelSystem.Receipts
}

func (c *ReceiptController) Show() {
	c.GetUserLogin()
	c.TplNames = "receipts/show.tpl"
	id := c.Ctx.Input.Param(":id")
	hotelSystem := models.GetInstance()
	receipt := hotelSystem.FindReceipt(id)
	if receipt != nil {
		c.Data["Receipt"] = receipt
	} else {
		c.Redirect("/", 302)
	}
}
