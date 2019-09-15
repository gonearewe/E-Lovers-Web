package controllers

import (
	"github.com/gonearewe/E-Lovers-Web/models"

	"github.com/astaxie/beego"
	"github.com/gonearewe/E-Lovers-Web/email"
	"github.com/gonearewe/E-Lovers-Web/tools"
)

type RegisterController struct {
	beego.Controller
	// step int //step 1:send verify code,step 2:register
	email      string
	verifyCode string
}

/*
if verifyCode is "" ,it means the client didn't send verify email;

save true verifyCode in the Session to be
compared with what client sends
*/
func (c *RegisterController) Prepare() {
	// if c.GetSession("step")==2{
	// 	c.step=2
	// }else {
	// 	c.step=1
	// }
	code := c.GetSession("verifyCode").(string)
	email := c.GetSession("email").(string)
	if code != "" && c.GetString("emailaddr") == email {
		c.verifyCode = code
	}

}

func (c *RegisterController) Get() {
	c.Data["title"] = "Register"
	c.TplName = "register.html"
}

func (c *RegisterController) Post() {
	log := tools.NewLogger()

	if c.verifyCode == "" {
		if models.UserEmailExist(c.GetString("emailaddr")) {
			c.Data["json"] = map[string]interface{}{"code": 2, "message": "邮箱已被占用"}
			c.ServeJSON()
			return
		}

		verifyCode, err := email.SendVerifyCodeEmail(c.GetString("emailaddr"))
		if err != nil {
			log.Informational("发送邮件失败:%s", err.Error())
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "发送邮件失败"}
			c.ServeJSON()
			return
		}
		c.SetSession("verifyCode", verifyCode)
		c.SetSession("email", c.GetString("emailaddr"))
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "邮件发送成功"}
		c.ServeJSON()
		return
	}

	if c.GetString("password") != c.GetString("repassword") {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "两次输入密码不一致"}
		c.ServeJSON()
		return
	}

	newUser := models.NewUser(
		c.GetString("username"),
		c.GetString("emailaddr"),
		tools.Md5(c.GetString("password")),
	)

	if newUser.Exist() {
		c.Data["json"] = map[string]interface{}{"code": 2, "message": "用户名已被占用"}
		c.ServeJSON()
		return
	}

	if id, err := newUser.Insert(); err != nil {
		log.Error("注册失败:%s", err.Error())
		c.Data["json"] = map[string]interface{}{"code": 3, "message": "注册失败"}
	} else {
		log.Informational("注册成功:ID=%d,name=%s", id, newUser.GetName())
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "注册成功"}
	}
	c.ServeJSON()

}
