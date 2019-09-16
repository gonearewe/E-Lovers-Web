package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gonearewe/E-Lovers-Web/models"
	"github.com/gonearewe/E-Lovers-Web/tools"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Data["title"] = "Log In"
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	log := tools.NewLogger()
	defer log.Close()

	user := models.NewUser(c.GetString("username"), "", tools.Md5(c.GetString("password")))
	if user.Exist() == false {
		c.Data["json"] = map[string]interface{}{"code": 4, "message": "用户名不存在"}
		c.ServeJSON()
		return
	}

	if ok, err := user.VerifyPassword(); err != nil {
		log.Error("数据库查询出错:%s", err.Error())
		c.Data["json"] = map[string]interface{}{"code": 5, "message": "数据库查询出错"}
		c.ServeJSON()
		return
	} else if !ok {
		log.Error("密码错误:%s:尝试登录", user.GetName())
		c.Data["json"] = map[string]interface{}{"code": 6, "message": "密码错误"}
		c.ServeJSON()
		return
	} else {
		log.Informational("用户登录:%s", user.GetName())
		c.SetSession("LoginUser", user.GetName())
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "登录成功"}
		c.ServeJSON()
		return
	}
}
