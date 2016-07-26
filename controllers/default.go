package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "goer.pub"
	c.Data["Email"] = "goudanbaba123@gmail.com"
	c.TplName = "index.html"
}