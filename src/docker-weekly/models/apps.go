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
type Apps struct {
  Id             int       `orm:"pk"`
  Title          string
  App_link          string
  Content          string
  Number         int
  Photo          string
  Image_link          string
  Image_name          string
  Image_user          string
  Article_num         int
}


func QueryApps(number int) (apps []*Apps, err error) {
  o := orm.NewOrm()
  _, err = o.QueryTable("apps").Filter("Article_num", number).OrderBy("Number").All(&apps)
  return
}