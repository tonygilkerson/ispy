package ui

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/tonygilkerson/ispy/internal/util"
)

type SayHi struct {
	Heading string
	Prompt   string
}

func (ctx *HandlerContext) SayHiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")

	pageValues := SayHi{
		Heading: "Say Hi!",
		Prompt:  "Press button to say hi...",
	}

	// Reuse template if possible
	tmpl, exists := ctx.PageTemplates["SayHi"]
	if !exists {
		tmplFile := ctx.templateRoot + "/templates/say-hi.gotmpl"
		log.Printf("Create template from: %v", tmplFile)

		tmplStr, err := os.ReadFile(tmplFile)
		util.DoOrDie(err)

		tmpl, err = template.New("SayHi").Parse(string(tmplStr))
		util.DoOrDie(err)

		ctx.PageTemplates["SayHi"] = tmpl
	}

	log.Printf("Execute template: %v", tmpl.Name())
	tmpl.Execute(w, pageValues)

}

func (ctx *HandlerContext) SayHiResponseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")

	homePage := []byte("Hello from ISpy!")

	_, _ = w.Write(homePage)

}
