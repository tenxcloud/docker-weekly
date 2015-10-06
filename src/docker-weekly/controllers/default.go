package controllers

import (
  //"fmt"
  "strconv"
	"github.com/astaxie/beego"
  "docker-weekly/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Weekly() {
  method := "Weekly"

  c.Data["Website"] = "tenxcloud.com"
  c.Data["Email"] = "service@tenxcloud.com"
  c.TplNames = "weekly.tpl"

  id, _ := strconv.Atoi(c.Ctx.Input.Param(":Id"))

  //fmt.Println(id)
  if id == 0 {
    articles, err := models.SearchArticles()
    if err != nil {
      beego.Error("Error", method, err)
    }
    id = articles[0].Number
  }
  apps, err := models.QueryApps(id)
  article, err := models.QueryArticle(id)
  if err != nil {
    beego.Error("Error", method, err)
  }
  //fmt.Println(article[0])
  c.Data["apps"] = apps
  c.Data["article"] = article[0]
}
