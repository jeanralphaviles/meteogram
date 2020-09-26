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
	r.HandleFunc("/meteogram/{latitude},{longitude}", h.Meteogram).
		Methods("GET")
	h := handlers.LoggingHandler(os.Stdout, r)
	log.Printf("Starting HTTP server listening on %s...", *port)
	log.Fatal(http.ListenAndServe(*port, h))
}
