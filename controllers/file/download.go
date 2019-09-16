package controllers_file

import (
	"io/ioutil"

	"github.com/gonearewe/E-Lovers-Web/tools"

	"github.com/astaxie/beego"
)

type FileDownloadController struct {
	beego.Controller
}

func (c *FileDownloadController) Get() {
	c.Data["activeFile"] = true

	log := tools.NewLogger()
	defer log.Close()

	path := beego.AppConfig.String("filepath")
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Error("寻找要下载的文件时异常:%s", err.Error())
		return
	} else if len(files) == 0 {
		log.Notice("没有任何可供下载的文件") //:试图下载:%s
	}
	c.Ctx.Output.Download(path + "/" + files[0].Name())
}
