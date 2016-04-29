package controllers

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday"
)

type DocController struct {
	BaseController

	docFile string
}

func (dc *DocController) Prepare() {
	dc.BaseController.Prepare()
	dc.Data["Type"] = "doc"

	// get correct file in doc directory
	p := dc.Ctx.Request.URL.Path
	p = strings.Replace(p, "/docs", "beedoc/"+dc.Lang, -1)
	p = strings.TrimRight(p, "/")
	if filepath.Ext(p) == "" {
		p += ".md"
	}
	dc.docFile = p

	// serve static file
	if filepath.Ext(p) != ".md" {
		http.ServeFile(dc.Ctx.ResponseWriter, dc.Ctx.Request, p)
		dc.StopRun()
	}

	// render md doc file
	if err := dc.renderDoc(); err != nil {
		dc.CustomAbort(503, err.Error())
		return
	}

}

func (dc *DocController) renderDoc() error {
	bytes, err := ioutil.ReadFile(dc.docFile)
	if err != nil {
		return err
	}
	dc.Data["DocFile"] = dc.docFile
	dc.Data["DocContent"] = template.HTML(markdown(bytes))
	return nil
}

func (dc *DocController) Get() {}

var (
	tab    = []byte("\t")
	spaces = []byte("    ")
)

type markdownRender struct {
	blackfriday.Renderer
}

// BlockCode overrides code block
func (mr *markdownRender) BlockCode(out *bytes.Buffer, text []byte, lang string) {
	var tmp bytes.Buffer
	mr.Renderer.BlockCode(&tmp, text, strings.ToLower(lang))
	out.Write(bytes.Replace(tmp.Bytes(), tab, spaces, -1))
}

func markdown(raw []byte) []byte {
	htmlFlags := 0 |
		blackfriday.HTML_USE_XHTML |
		blackfriday.HTML_USE_SMARTYPANTS |
		blackfriday.HTML_SMARTYPANTS_FRACTIONS |
		blackfriday.HTML_SMARTYPANTS_LATEX_DASHES

	renderer := &markdownRender{
		Renderer: blackfriday.HtmlRenderer(htmlFlags, "", ""),
	}

	extensions := 0 |
		blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
		blackfriday.EXTENSION_TABLES |
		blackfriday.EXTENSION_FENCED_CODE |
		blackfriday.EXTENSION_AUTOLINK |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_AUTO_HEADER_IDS |
		blackfriday.EXTENSION_HEADER_IDS

	return blackfriday.Markdown(raw, renderer, extensions)
}
