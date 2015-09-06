package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "tenxcloud.com"
	c.Data["Email"] = "service@tenxcloud.com"
	c.TplNames = "index.tpl"
}
