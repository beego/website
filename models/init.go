package models

import "github.com/astaxie/beego/config"

// Init init models
func Init() {
	iniFile, err := config.NewConfig("ini", "conf/nav.ini")
	if err != nil {
		panic(err)
	}
	ReadNav(iniFile)
	ReadI18n(iniFile)
}
