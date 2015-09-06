package routers

import (
	"docker-weekly/controllers"
	"github.com/astaxie/beego"
)

func init() {
  beego.Router("/", &controllers.MainController{})
  beego.Router("/issues", &controllers.IssueController{})
}
