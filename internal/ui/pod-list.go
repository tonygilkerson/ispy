package ui

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/tonygilkerson/ispy/internal/k8s"
	"github.com/tonygilkerson/ispy/internal/util"
)

type PodList struct {
	Heading string
	Intro   string
	Pods    []string
}

func (ctx *HandlerContext) PodListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")

	pageValues := PodList{
		Heading: "Pods",
		Intro:   "Below is a list of pods",
		Pods:    []string{},
	}

	// Reuse template if possible
	tmpl, exists := ctx.pageTemplates["PodList"]
	if !exists {
		tmplFile := ctx.wwwRoot + "/www/templates/pod-list.gotmpl"
		log.Printf("Create template from: %v", tmplFile)

		tmplStr, err := os.ReadFile(tmplFile)
		util.DoOrDie(err)

		tmpl, err = template.New("PodList").Parse(string(tmplStr))
		util.DoOrDie(err)

		ctx.pageTemplates["PodList"] = tmpl
	}

	// An empty string returns all namespaces
	namespace := ""
	pods, err := k8s.GetPods(namespace, ctx.clientset)
	util.DoOrDie(err)

	for _, pod := range pods.Items {
		pageValues.Pods = append(pageValues.Pods, pod.Name)
	}

	log.Printf("Execute template: %v", tmpl.Name())
	tmpl.Execute(w, pageValues)

}
