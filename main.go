package main

import (
	"github.com/astaxie/beego"
	"github.com/beego/website/models"
	"github.com/beego/website/routers"
)

func main() {
	beego.BeeLogger.SetLogger("console", "")

	models.Init()
	routers.Init()
	beego.Run()
}
