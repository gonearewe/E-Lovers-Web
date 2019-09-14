package controllers

import (
	_"github.com/astaxie/beego"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Get() {
	c.Data["activeHome"]=true
	c.Data["title"]="Home Page"
	c.TplName = "home.html"
}
