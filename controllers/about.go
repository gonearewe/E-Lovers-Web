package controllers

import (
	"github.com/astaxie/beego"
)

type AboutController struct {
	beego.Controller
}

func (c *AboutController) Get() {
	c.Data["activeAbout"] = true
	c.Data["title"] = "About"
	c.TplName = "about.html"
}
