package controllers

type DocController struct {
	BaseController
}

func (dc *DocController) Get() {
	dc.Data["Website"] = "Doc page"
	dc.Data["Email"] = "astaxie@gmail.com"
	dc.TplName = "index.tpl"
}
