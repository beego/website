package controllers

import (
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/astaxie/beego"
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
	bc.cacheFile = path.Join(beego.AppConfig.DefaultString("cache_dir", "_cache"), bc.Lang, bc.Ctx.Input.URL())
	if !strings.HasSuffix(bc.cacheFile, ".html") {
		bc.cacheFile += ".html"
	}
	if fi, _ := os.Stat(bc.cacheFile); fi != nil && !fi.IsDir() {
		http.ServeFile(bc.Ctx.ResponseWriter, bc.Ctx.Request, bc.cacheFile)
		bc.StopRun()
	}

	bc.Layout = "layout.html"
	bc.setProperTemplateFile()
}

func (bc *BaseController) setProperTemplateFile() {
	p := bc.Ctx.Request.URL.Path
	paths := strings.Split(p, "/")
	if len(paths) >= 2 {
		bc.TplName = paths[1]
	}
	if bc.TplName == "" {
		bc.TplName = "index.html"
	}
	if !strings.HasSuffix(bc.TplName, ".html") {
		bc.TplName += ".html"
	}
}

func (bc *BaseController) Render() error {
	if renderBytes, _ := bc.RenderBytes(); len(renderBytes) > 0 {
		bc.cacheBody = renderBytes
		bc.cacheEnable = beego.AppConfig.DefaultBool("cache_flag", false)
	}
	return bc.Controller.Render()
}

func (bc *BaseController) Finish() {
	if bc.cacheEnable {
		os.MkdirAll(path.Dir(bc.cacheFile), os.ModePerm)
		ioutil.WriteFile(bc.cacheFile, bc.cacheBody, os.ModePerm) // todo : error handle
	}
}
