package ui

import (
	"log"
	"net/http"
)

func (ctx *HandlerContext) StaticHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css;charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")

	uri := r.RequestURI 
	staticFile := ctx.wwwRoot + "/www/" + uri
	log.Printf("Server up file: %v", staticFile)

	http.ServeFile(w, r, staticFile)

}
