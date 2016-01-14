package routers

import (
	"github.com/astaxie/beego"
	"github.com/beego/website/controllers"
)

func init() {
	beego.Router("/", new(controllers.MainController))
	beego.Router("/doc", new(controllers.DocController))
	beego.Router("/doc/*", new(controllers.DocController))
}
