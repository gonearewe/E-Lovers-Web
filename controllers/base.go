package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
	IsLogin   bool
	LoginUser interface{}
}

func (c *BaseController) Prepare(){
	if user:=c.GetSession("LoginUser");user!=nil{
		c.IsLogin=true
		c.LoginUser=user
	}else{
		c.IsLogin=false
	}
	c.Data["IsLogin"]=c.IsLogin

}