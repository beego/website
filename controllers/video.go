package controllers

import (
	"github.com/beego/website/models"
	"strings"
)

type VideoController struct {
	BaseController
}

func (vc *VideoController) Get() {
	vc.Data["Video"] = models.Videos[strings.ToLower(vc.Lang)]
}
