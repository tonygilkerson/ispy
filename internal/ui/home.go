package ui

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/tonygilkerson/ispy/internal/util"
)

type Home struct {
	Heading string
	Intro   string
}

func (ctx *HandlerContext) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")

	pageValues := Home{
		Heading: "Welcome, from ISpy!",
		Intro:   "The ISpy home page",
	}

	// Reuse template if possible
	tmpl, exists := ctx.PageTemplates["HomePage"]
	if !exists {
		tmplFile := ctx.templateRoot + "/templates/home.gotmpl"
		log.Printf("Create template from: %v", tmplFile)

		tmplStr, err := os.ReadFile(tmplFile)
		util.DoOrDie(err)

		tmpl, err = template.New("HomePage").Parse(string(tmplStr))
		util.DoOrDie(err)

		ctx.PageTemplates["HomePage"] = tmpl
	}

	log.Printf("Execute template: %v", tmpl.Name())
	tmpl.Execute(w, pageValues)

}
