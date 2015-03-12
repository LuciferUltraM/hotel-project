package controllers

import (
	"github.com/astaxie/beego"
	"github.com/code-mobi/hotel/models"
)

type AuthenController struct {
	beego.Controller
}

type FormLogin struct {
	UserName string
	Password string
}

func (c *AuthenController) Post() {
	login := FormLogin{}
	if err := c.ParseForm(&login); err != nil {
		c.Redirect("/", 302)
	}
	hotelSystem := models.GetInstance()
	user := hotelSystem.FindReceptionist(login.UserName)
	if user != nil && login.Password == user.Password {
		c.SetSession("username", user.UserName)
	}
	c.Redirect("/", 302)
}

func (c *AuthenController) Delete() {
	c.DelSession("username")
	c.Redirect("/", 302)
}
