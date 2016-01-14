package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

type BaseController struct {
	beego.Controller
	Lang string

	// static file cache support
	cacheFile   string
	cacheEnable bool
	cacheBody   []byte
}

func (bc *BaseController) Prepare() {
	if bc.Lang = bc.Ctx.GetCookie("lang"); bc.Lang == "" {
		bc.Lang = "en-Us"
	}
	bc.cacheFile = path.Join(beego.AppConfig.String("cache_dir"), bc.Lang, bc.Ctx.Input.URL())
	if !strings.HasSuffix(bc.cacheFile, ".html") {
		bc.cacheFile += ".html"
	}
	if fi, _ := os.Stat(bc.cacheFile); fi != nil && !fi.IsDir() {
		http.ServeFile(bc.Ctx.ResponseWriter, bc.Ctx.Request, bc.cacheFile)
		bc.StopRun()
	}
}

func (bc *BaseController) Render() error {
	if renderBytes, _ := bc.RenderBytes(); len(renderBytes) > 0 {
		bc.cacheBody = renderBytes
		bc.cacheEnable = true
	}
	return bc.Controller.Render()
}

func (bc *BaseController) Finish() {
	if bc.cacheEnable {
		os.MkdirAll(path.Dir(bc.cacheFile), os.ModePerm)
		ioutil.WriteFile(bc.cacheFile, bc.cacheBody, os.ModePerm) // todo : error handle
	}
}
