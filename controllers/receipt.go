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

}
