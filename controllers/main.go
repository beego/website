package controllers

type MainController struct {
	BaseController
}

func (mc *MainController) Get() {
	mc.Data["Website"] = "beego.me"
	mc.Data["Email"] = "astaxie@gmail.com"
	mc.TplName = "index.tpl"
}
