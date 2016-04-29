package controllers

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

type BlogController struct {
	BaseController

	blogFile string
}

func (bc *BlogController) Prepare() {
	bc.BaseController.Prepare()
	bc.Data["Type"] = "blog"

	// get correct file in beeblog directory
	p := bc.Ctx.Request.URL.Path
	p = strings.TrimRight(p, "/")
	if p == "/blog" {
		p = filepath.Join("beeblog", bc.Lang, "blog.md")
	} else {
		p = strings.Replace(p, "/blog", "beeblog/"+bc.Lang, -1)
	}
	if filepath.Ext(p) == "" {
		p += ".md"
	}
	bc.blogFile = p

	// serve static file
	if filepath.Ext(p) != ".md" {
		http.ServeFile(bc.Ctx.ResponseWriter, bc.Ctx.Request, p)
		bc.StopRun()
	}

	// render blog doc file
	if err := bc.renderBlog(); err != nil {
		bc.CustomAbort(503, err.Error())
		return
	}
}

func (bc *BlogController) renderBlog() error {
	bytes, err := ioutil.ReadFile(bc.blogFile)
	if err != nil {
		return err
	}
	bc.Data["BlogFile"] = bc.blogFile
	bc.Data["BlogContent"] = template.HTML(markdown(bytes))
	return nil
}

func (bc *BlogController) Get() {}
