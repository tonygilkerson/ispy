package ui

import (
	"log"
	"os"
	"text/template"

	"github.com/tonygilkerson/ispy/internal/k8s"
	"k8s.io/client-go/kubernetes"
)

type HandlerContext struct {
	wwwRoot  string
	pageTemplates map[string]*template.Template
	clientset     *kubernetes.Clientset
}

func NewHandlerContext() *HandlerContext {

	wwwRoot, exists := os.LookupEnv("ISPY_WWW_ROOT")
	if exists {
		log.Printf("Using environment variable ISPY_WWW_ROOT: %v", wwwRoot)
	} else {
		wwwRoot, _ = os.Getwd()
		log.Printf("ISPY_WWW_ROOT environment variable not set, using default value: %v", wwwRoot)
	}

	var hc HandlerContext
	hc.wwwRoot = wwwRoot
	hc.pageTemplates = make(map[string]*template.Template)
	hc.clientset = k8s.GetClientSet()

	return &hc
}
