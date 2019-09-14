package controllers

type ExitController struct {
	BaseController
}

func (c *ExitController) Get() {
	c.Data["title"] = "Log Out"
	c.DelSession("LoginUser")
	c.Redirect("/", 302)

}
