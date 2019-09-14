package controllers

import (
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.Data["title"] = "Register"
	c.TplName = "register.html"
}

// func (c *RegisterController) Post() {
// 	if c.GetString("password") != c.GetString("repassword") {
// 		c.Data["json"] = map[string]interface{}{"code": 1, "message": "两次输入密码不一致"}
// 		c.ServeJSON()
// 		return
// 	}

// 	newUser := new(models.User)
// 	newUser.Username = c.GetString("username")
// 	newUser.Password = tools.Md5(c.GetString("password"))

// 	if newUser.Exist() {
// 		c.Data["json"] = map[string]interface{}{"code": 2, "message": "用户名已被占用"}
// 		c.ServeJSON()
// 		return
// 	}

// 	newUser.Status = true
// 	newUser.Createtime = time.Now().Unix()
// 	if _, err := newUser.Insert(); err != nil {
// 		fmt.Println(err)
// 		c.Data["json"] = map[string]interface{}{"code": 3, "message": "注册失败"}
// 	} else {
// 		fmt.Println("User Register")
// 		fmt.Println(newUser.GetByName())
// 		c.Data["json"] = map[string]interface{}{"code": 0, "message": "注册成功"}
// 	}
// 	c.ServeJSON()

// }
