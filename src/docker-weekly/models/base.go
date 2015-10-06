package models

import (
  "fmt"
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
  _ "github.com/go-sql-driver/mysql"
  "os"
  // "strconv"
)

func init() {
  //dbhost := beego.AppConfig.String("dbhost")
  dbport := beego.AppConfig.String("dbport")
  dbuser := beego.AppConfig.String("dbuser")
  dbpassword := beego.AppConfig.String("dbpassword")
  dbname := beego.AppConfig.String("dbname")
  if dbport == "" {
    dbport = "3306"
  }
  dsn := dbuser + ":" + dbpassword + "@tcp(localhost:3306)/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
  if os.Getenv("DB_USER") != "" {
    fmt.Println("DB_USER: " + os.Getenv("DB_USER"))
    dsn = os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp("+ os.Getenv("DB_HOST") +")/docker-weekly?charset=utf8&loc=Asia%2FShanghai"
  }

  fmt.Println("dsn: " + dsn)
  //orm.RegisterDriver("mysql", orm.DR_MySQL)
  dbmaxIdle, _ := beego.AppConfig.Int("dbmaxIdle")
  dbmaxConn, _ := beego.AppConfig.Int("dbmaxConn")
  orm.RegisterDataBase("default", "mysql", dsn, dbmaxIdle, dbmaxConn)

  orm.RegisterModel(new(Article))
  orm.RegisterModel(new(Apps))
  if beego.AppConfig.String("mode") != "pro" {
    orm.Debug = true
  }
}
