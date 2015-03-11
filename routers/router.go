package routers

import (
	"github.com/astaxie/beego"
	"github.com/code-mobi/hotel/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/rooms", &controllers.RoomController{})
	beego.Router("/roombookings", &controllers.RoomBookingController{})
}
