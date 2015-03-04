package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.code-mobi.com"
	c.Data["Email"] = "luciferultram@gmail.com"
	c.TplNames = "index.tpl"
}
