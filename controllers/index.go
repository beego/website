package controllers

type IndexController struct {
	BaseController
}

func (ic *IndexController) Get() {
	ic.Layout = "layout.html"
}
