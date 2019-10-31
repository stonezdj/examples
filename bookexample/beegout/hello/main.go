package main

import (
	"bookexample/beegout/hello/controllers"
	_ "bookexample/beegout/hello/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Run()
}
