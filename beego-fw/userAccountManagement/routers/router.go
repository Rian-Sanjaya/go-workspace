package routers

import (
	"beego-fw/userAccountManagement/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &controllers.MainController{})
	// both GET and POST are handled by 'Login' function.
	// the :back is a request parameter, it tells what page to return after successful login
	beego.Router("/user/login/:back", &controllers.MainController{}, "get,post:Login")
	beego.Router("/user/logout", &controllers.MainController{}, "get:Logout")
	beego.Router("/user/register", &controllers.MainController{}, "get,post:Register")
	beego.Router("/user/profile", &controllers.MainController{}, "get,post:Profile")
	// :uuid is a request parameter that must match the regular expression for a Version 4 UUID
	// the GET operation is handled by the 'verify' function
	beego.Router("/user/verify/:uuid({[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}})", &controllers.MainController{}, "get:verify")
	beego.Router("/user/remove", &controllers.MainController{}, "get,post:Remove")
	beego.Router("/notice", &controllers.MainController{}, "get:Notice")
}
