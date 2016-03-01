package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

var (
	Videos = make(map[string]*Video)
)

type Video struct {
	Desc string
	List []map[string]string
}

func ReadVideo(langMap map[string]string) {
	for lang := range langMap {
		file := "conf/video/" + lang + ".ini"
		video, err := NewVideo(file)
		if err != nil {
			panic(err)
		}
		Videos[lang] = video
	}
}

func NewVideo(file string) (*Video, error) {
	iniFile, err := config.NewConfig("ini", file)
	if err != nil {
		return nil, err
	}
	v := new(Video)
	v.Desc = iniFile.String("description")

	var (
		idx      = 1
		section  string
		videoMap map[string]string
	)
	for {
		section = fmt.Sprintf("video-%d", idx)
		videoMap, _ = iniFile.GetSection(section)
		if len(videoMap) == 0 {
			break
		}

		v.List = append(v.List, videoMap)

		idx++
		videoMap = nil
	}

	beego.Info("Model.Read.[" + file + "]")
	return v, nil
}
