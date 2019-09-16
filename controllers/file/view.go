package controllers_file

import (
	"github.com/astaxie/beego"
)

type FileViewController struct {
	beego.Controller
}

func (c *FileViewController) Get() {
	c.Data["activeFile"] = true
	c.Data["title"] = "File"
	c.TplName = "file/view.html"
}
