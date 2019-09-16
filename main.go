package main

import (
	"github.com/astaxie/beego"
	_ "github.com/gonearewe/E-Lovers-Web/routers"
)

func main() {

	beego.SetStaticPath("/file/static", "static")
	beego.Run()
}
