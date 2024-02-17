package main

import (
	"log"
	"net/http"

	"github.com/tonygilkerson/ispy/internal/ui"
)

func main() {

	// Log to the console with date, time and filename prepended
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//
	// Create Server
	//
	mux := http.NewServeMux()
	svr := &http.Server{Addr: ":8080", Handler: mux}

	//
	// Init handler context
	//
	hctx := ui.NewHandlerContext()

	//
	// Define routes
	//
	mux.HandleFunc("/", hctx.HomeHandler)
	mux.HandleFunc("/say-hi", hctx.SayHiHandler)
	mux.HandleFunc("/say-hi-response", hctx.SayHiResponseHandler)

	// For the live and readness probes
	mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})

	log.Printf("Listening on: %v", svr.Addr)
	log.Fatal(svr.ListenAndServe())
}
