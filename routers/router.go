package routers

import (
	"github.com/code-mobi/hotel/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
