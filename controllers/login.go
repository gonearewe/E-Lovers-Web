package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Data["title"] = "Log In"
	c.TplName = "login.html"
}

// func (c *LoginController) Post() {
// 	olduser := new(models.User)
// 	olduser.Username = c.GetString("username")
// 	olduser.Password = tools.Md5(c.GetString("password"))
// 	if olduser.Exist() == false {
// 		c.Data["json"] = map[string]interface{}{"code": 4, "message": "用户名不存在"}
// 		c.ServeJSON()
// 		return
// 	}

// 	if actualuser, err := olduser.GetByName(); err != nil {
// 		fmt.Println(err)
// 		c.Data["json"] = map[string]interface{}{"code": 5, "message": "数据库查询出错"}
// 		c.ServeJSON()
// 		return
// 	} else if actualuser.Password != olduser.Password {
// 		fmt.Println(*olduser, " Password Incorrect !!!")
// 		c.Data["json"] = map[string]interface{}{"code": 6, "message": "密码错误"}
// 		c.ServeJSON()
// 		return
// 	} else {
// 		fmt.Println(*actualuser, " User Login")
// 		c.SetSession("LoginUser", olduser.Username)
// 		c.Data["json"] = map[string]interface{}{"code": 0, "message": "登录成功"}
// 		c.ServeJSON()
// 		return
// 	}
// }
