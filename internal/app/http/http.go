// Package http contains HTTP handlers for Meteogram.
package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/icodealot/noaa"
	"github.com/jeanralphaviles/meteogram/internal/builder"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"net/http"
)

// Meteogram handles requests to /meteogram, returning Meteograms.
func Meteogram(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lat, long := vars["latitude"], vars["longitude"]
	f, err := noaa.GridpointForecast(lat, long)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not retrieve forecast: %q", err), http.StatusInternalServerError)
		return
	}
	csv, err := builder.CsvMeteogram(f)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not print CSV: %q", err), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(csv))
}

// Readme renders readme.md to the user as a guide to Meteogram.
func Readme(w http.ResponseWriter, r *http.Request) {
	md, err := ioutil.ReadFile("README.md")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	html := blackfriday.MarkdownCommon(md)
	w.Write(html)
}
