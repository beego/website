package routers

import (
	"github.com/astaxie/beego"
	"github.com/beego/website/controllers"
)

func init() {
	beego.Router("/", new(controllers.MainController))
	beego.Router("/docs", new(controllers.DocController))
	beego.Router("/docs/*", new(controllers.DocController))
}
