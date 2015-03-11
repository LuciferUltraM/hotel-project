package routers

import (
	"github.com/astaxie/beego"
	"github.com/code-mobi/hotel/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/room", &controllers.RoomController{})
	beego.Router("/roombooking", &controllers.RoomBookingController{}, "get:List")
	beego.Router("/roombooking/:id([0-9]+", &controllers.RoomBookingController{}, "get:Show")
	beego.Router("/roombooking/:id([0-9]+", &controllers.RoomBookingController{}, "post:Update")
}
