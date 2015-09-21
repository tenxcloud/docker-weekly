package routers

import (
	"docker-weekly/controllers"
	"github.com/astaxie/beego"
)

func init() {
  beego.Router("/", &controllers.MainController{}, "GET:Weekly3")
  beego.Router("/issue/2", &controllers.MainController{}, "GET:Weekly2")
  beego.Router("/issue/1", &controllers.MainController{}, "GET:Weekly1")
  beego.Router("/issues", &controllers.IssueController{})
}
