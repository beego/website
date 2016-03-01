package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

// Init init models
func Init() {
	iniFile, err := config.NewConfig("ini", "conf/nav.ini")
	if err != nil {
		panic(err)
	}
	beego.Info("Model.Read.[conf/nav.ini]")
	ReadNav(iniFile)

	langMap, _ := iniFile.GetSection("lang")
	ReadI18n(langMap)
	ReadVideo(langMap)
}
