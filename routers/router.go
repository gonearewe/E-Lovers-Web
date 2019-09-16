package routers

import (
	"github.com/astaxie/beego"
	"github.com/gonearewe/E-Lovers-Web/controllers"
	controllers_file "github.com/gonearewe/E-Lovers-Web/controllers/file"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/exit", &controllers.ExitController{})
	beego.Router("/about", &controllers.AboutController{})

	beego.Router("/file/view", &controllers_file.FileViewController{})
	beego.Router("/file/download", &controllers_file.FileDownloadController{})

}
