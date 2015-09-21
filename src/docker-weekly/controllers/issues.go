package controllers

import (
  "github.com/astaxie/beego"
)

type IssueController struct {
  beego.Controller
}

func (c *IssueController) Get() {
  c.TplNames = "issues.tpl"
  //c.Data["maps0"] = ReadSql()
}
