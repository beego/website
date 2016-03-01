package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

var (
	Langs map[string]string
	I18ns = make(map[string]*I18n)
)

// ReadI18n read i18n data
func ReadI18n(langMap map[string]string) {
	for lang := range langMap {
		file := "conf/lang/" + lang + ".ini"
		beego.Info("Model.Read.[" + file + "]")
		i18n, err := NewI18n(lang, file)
		if err != nil {
			panic(err)
		}
		I18ns[lang] = i18n
	}
	Langs = langMap
}

// I18n object
type I18n struct {
	Lang string // language string
	cfg  config.Configer
}

// Tr converts string
func (i *I18n) Tr(str string) string {
	return i.cfg.DefaultString(str, str)
}

// Trf converts string with arguments
func (i *I18n) Trf(str string, values ...interface{}) string {
	return fmt.Sprintf(i.Tr(str), values...)
}

// NewI18n reads ini file
func NewI18n(lang string, file string) (*I18n, error) {
	iniFile, err := config.NewConfig("ini", file)
	if err != nil {
		return nil, err
	}
	return &I18n{
		Lang: lang,
		cfg:  iniFile,
	}, nil
}
