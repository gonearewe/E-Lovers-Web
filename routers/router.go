package routers

import (
	"github.com/astaxie/beego"
	"github.com/gonearewe/E-Lovers-Web/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/exit", &controllers.ExitController{})

}
