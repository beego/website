package controllers

type IndexController struct {
	BaseController
}

func (ic *IndexController) Get() {
	ic.Data["Title"] = "beego website"
	ic.Data["Type"] = "index"
}
