package routers

import (
	"github.com/astaxie/beego"
	"github.com/code-mobi/hotel/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.AuthenController{})
	beego.Router("/logout", &controllers.AuthenController{}, "get:Delete")
	beego.Router("/room", &controllers.RoomController{})
	beego.Router("/roombooking", &controllers.RoomBookingController{}, "get:List")
	beego.Router("/roombooking/:id([0-9]+", &controllers.RoomBookingController{}, "get:Show")
	beego.Router("/roombooking/:id([0-9]+", &controllers.RoomBookingController{}, "post:Update")
	beego.Router("/roombooking/:id([0-9]+/payment", &controllers.RoomBookingController{}, "get:NewPayment")
	beego.Router("/roombooking/:id([0-9]+/payment", &controllers.RoomBookingController{}, "post:SavePayment")
	beego.Router("/roombooking/:id([0-9]+/checkin", &controllers.RoomBookingController{}, "post:CheckIn")
	beego.Router("/roombooking/:id([0-9]+/checkout", &controllers.RoomBookingController{}, "get:CheckOut")
	beego.Router("/roombooking/:id([0-9]+/checkout", &controllers.RoomBookingController{}, "post:CheckOut")
	beego.Router("/receipt", &controllers.ReceiptController{})
	beego.Router("/receipt/:id([0-9]+", &controllers.ReceiptController{}, "get:Show")
	beego.Router("/equipment", &controllers.EquipmentController{})
}
