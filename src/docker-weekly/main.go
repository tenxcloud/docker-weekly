package main

import (
	_ "docker-weekly/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

