package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (main *MainController) Get() {
	main.Data["Website"] = "beego.me"
	main.Data["Email"] = "astaxie@gmail.com"
	main.TplName = "index.tpl"
}

func (main *MainController) HelloSitepoint() {
	// initialised 3 template variables by storing them in a field in the Controller, called Data, which is a map
	main.Data["Website"] = "My Website"
	main.Data["Email"] = "your.email.address@example.com"
	main.Data["EmailName"] = "Your Name"
	main.Data["Id"] = main.Ctx.Input.Param(":id")
	// specified the template file name, by default beego will look in the Views directory
	main.TplName = "default/hello-sitepoint.tpl"
}
