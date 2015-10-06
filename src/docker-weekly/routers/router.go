package routers

import (
	"docker-weekly/controllers"
	"github.com/astaxie/beego"
)

func init() {
  beego.Router("/", &controllers.MainController{}, "GET:Weekly")
  beego.Router("/issue/:Id", &controllers.MainController{}, "GET:Weekly")
  beego.Router("/issues", &controllers.IssueController{})
}
