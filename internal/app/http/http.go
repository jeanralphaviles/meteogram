// Package http contains HTTP handlers for Meteogram.
package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/icodealot/noaa"
	"github.com/jeanralphaviles/meteogram/internal/builder"
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
