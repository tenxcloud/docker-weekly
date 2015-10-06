package models

import (
  //"errors"
  //"fmt"
  //"github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
  //"strconv"
  //"time"
)

//model for article
type Article struct {
  Id             int       `orm:"pk"`
  Number         int
  Title          string
  Created        string
}

func (m *Article) TableName() string {
  return "articles"
}

func SearchArticles() (articles []*Article, err error) {
  o := orm.NewOrm()
  _, err = o.QueryTable("articles").OrderBy("-Number").All(&articles)
  return
}

func QueryArticle(number int) (articles []*Article, err error) {
  o := orm.NewOrm()
  _, err = o.QueryTable("articles").Filter("Number", number).All(&articles)
  return
}