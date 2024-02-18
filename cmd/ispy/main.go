package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tonygilkerson/ispy/internal/ui"
)

func main() {

	// Log to the console with date, time and filename prepended
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//
	// Create Server
	//
	port, exists := os.LookupEnv("ISPY_SERVER_PORT")
	if exists {
		log.Printf("Using environment variable ISPY_SERVER_PORT: %v", port)
	} else {
		port = "8080"
		log.Printf("ISPY_SERVER_PORT environment variable not set, using default value: %v", port)
	}
	mux := http.NewServeMux()
	svr := &http.Server{Addr: ":"+port, Handler: mux}

	//
	// Define routes
	//
	hctx := ui.NewHandlerContext()
	mux.HandleFunc("/", hctx.HomeHandler)
	mux.HandleFunc("/static/", hctx.StaticHandler)
	mux.HandleFunc("/say-hi", hctx.SayHiHandler)
	mux.HandleFunc("/say-hi-response", hctx.SayHiResponseHandler)
	mux.HandleFunc("/pod-list", hctx.PodListHandler)
	mux.HandleFunc("/ns-list", hctx.NsListHandler)

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
