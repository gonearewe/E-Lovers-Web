package controllers

import (
	"github.com/astaxie/beego"
)

type FileViewController struct {
	beego.Controller
}

func (c *FileViewController) Get() {
	c.Data["title"] = "File"
	c.TplName = "file/view.html"
}
