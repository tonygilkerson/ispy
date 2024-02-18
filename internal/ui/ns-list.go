package ui

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/tonygilkerson/ispy/internal/k8s"
	"github.com/tonygilkerson/ispy/internal/util"
)

type NsList struct {
	Heading    string
	Intro      string
	Namespaces []string
}

func (ctx *HandlerContext) NsListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")

	pageValues := NsList{
		Heading:    "Namespaces",
		Intro:      "Below is a list of namespaces",
		Namespaces: []string{},
	}

	// Reuse template if possible
	tmpl, exists := ctx.pageTemplates["NsList"]
	if !exists {
		tmplFile := ctx.wwwRoot + "/www/templates/ns-list.gotmpl"
		log.Printf("Create template from: %v", tmplFile)

		tmplStr, err := os.ReadFile(tmplFile)
		util.DoOrDie(err)

		tmpl, err = template.New("NsList").Parse(string(tmplStr))
		util.DoOrDie(err)

		ctx.pageTemplates["NsList"] = tmpl
	}

	//ListNamespaces function call returns a list of namespaces in the kubernetes cluster
	namespaces, err := k8s.GetNamespaces(ctx.clientset)
	util.DoOrDie(err)

	for _, namespace := range namespaces.Items {
		pageValues.Namespaces = append(pageValues.Namespaces, namespace.Name)
	}

	log.Printf("Execute template: %v", tmpl.Name())
	tmpl.Execute(w, pageValues)

}
