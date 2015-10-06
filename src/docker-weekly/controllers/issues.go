package controllers

import (
  //"fmt"
  "github.com/astaxie/beego"
  "docker-weekly/models"
  //"time"
)

type IssueController struct {
  beego.Controller
}

func (c *IssueController) Get() {
  method := "IssueController"

  c.TplNames = "issues.tpl"
  articles, err := models.SearchArticles()

  if err != nil {
    beego.Error("Error", method, err)
  }
  //tm2, _ := time.Parse("2006-01-02", &articles[0].created)
  //fmt.Println(articles)

  c.Data["articles"] = articles
}
