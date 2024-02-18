package ui

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/tonygilkerson/ispy/internal/k8s"
	"github.com/tonygilkerson/ispy/internal/util"
)

type GetPods struct {
	Heading string
	Intro   string
	Pods    []string
}

func (ctx *HandlerContext) GetPodHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")

	pageValues := GetPods{
		Heading: "Pods",
		Intro:   "Below is a list of pods",
		Pods:    []string{"aaa","bbb","ccc"},
	}

	// Reuse template if possible
	tmpl, exists := ctx.PageTemplates["GetPods"]
	if !exists {
		tmplFile := ctx.templateRoot + "/templates/get-pods.gotmpl"
		log.Printf("Create template from: %v", tmplFile)

		tmplStr, err := os.ReadFile(tmplFile)
		util.DoOrDie(err)

		tmpl, err = template.New("GetPods").Parse(string(tmplStr))
		util.DoOrDie(err)


		ctx.PageTemplates["GetPods"] = tmpl
	}

	log.Printf("Execute template: %v", tmpl.Name())
	tmpl.Execute(w, pageValues)

	k8s.PodListWrapper()

}


