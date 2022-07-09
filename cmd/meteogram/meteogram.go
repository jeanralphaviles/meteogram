package main

import (
	"flag"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	h "github.com/jeanralphaviles/meteogram/internal/app/http"
	"log"
	"net/http"
	"os"
)

var (
	port = flag.String("address", ":8080", "Address to listen on.")
)

func main() {
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/", h.Readme).Methods("GET")
	r.HandleFunc("/meteogram/{latitude},{longitude}", h.Meteogram).Methods("GET")
	h := handlers.LoggingHandler(os.Stdout, r)

	// App Engine recommends that we listen on the port defined by the PORT variable.
	if env, ok := os.LookupEnv("PORT"); ok {
		log.Printf("Overriding --port=%s with PORT=%s environment variable.", *port, env)
		*port = ":" + env
	}

	log.Printf("Starting HTTP server listening on %s...", *port)
	log.Fatal(http.ListenAndServe(*port, h))
}
