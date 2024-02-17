package ui

import (
	"log"
	"os"
	"text/template"
)

type HandlerContext struct {
	templateRoot       string
	PageTemplates map[string]*template.Template
}

func NewHandlerContext() *HandlerContext {

	templateRoot, exists := os.LookupEnv("TEMPLATE_ROOT")
	if exists {
		log.Printf("Using environment variable TEMPLATE_ROOT: %v", templateRoot)
	} else {
		templateRoot, _ = os.Getwd()
		log.Printf("TEMPLATE_ROOT environment variable not set, using default value: %v", templateRoot)
	}

	var hc HandlerContext
	hc.templateRoot = templateRoot
	hc.PageTemplates = make(map[string]*template.Template)

	return &hc
}

func doOrDie(err error) {
	if err != nil {
		log.Panicf("Oops err: %v ", err)
	}
}
