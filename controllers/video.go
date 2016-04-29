package controllers

import (
	"strings"

	"github.com/beego/website/models"
)

type VideoController struct {
	BaseController
}

func (vc *VideoController) Get() {
	vc.Data["Type"] = "video"
	vc.Data["Video"] = models.Videos[strings.ToLower(vc.Lang)]
}
