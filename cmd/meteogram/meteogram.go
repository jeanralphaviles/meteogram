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
	port = flag.String("port", ":8080", "Port to listen on.")
)

func main() {
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/meteogram/{latitude},{longitude}", h.Meteogram).
		Methods("GET")
	h := handlers.LoggingHandler(os.Stdout, r)
	log.Fatal(http.ListenAndServe(*port, h))
}
