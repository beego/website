package routers

import (
	"github.com/astaxie/beego"
	"github.com/beego/website/controllers"
)

func Init() {
	beego.Router("/", new(controllers.IndexController))
	beego.Router("/*", new(controllers.MainController))
	beego.Router("/docs", new(controllers.DocController))
	beego.Router("/docs/*", new(controllers.DocController))
	beego.Router("/blog", new(controllers.BlogController))
	beego.Router("/blog/*", new(controllers.BlogController))
	beego.Router("/video", new(controllers.VideoController))
}
