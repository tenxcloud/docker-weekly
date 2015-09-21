package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Weekly3() {
	c.Data["Website"] = "tenxcloud.com"
	c.Data["Email"] = "service@tenxcloud.com"
	c.TplNames = "weekly3.tpl"
}

func (c *MainController) Weekly2() {
  c.Data["Website"] = "tenxcloud.com"
  c.Data["Email"] = "service@tenxcloud.com"
  c.TplNames = "weekly2.tpl"
}

func (c *MainController) Weekly1() {
  c.Data["Website"] = "tenxcloud.com"
  c.Data["Email"] = "service@tenxcloud.com"
  c.TplNames = "weekly1.tpl"
}
