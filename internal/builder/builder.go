// Package builder contains functions to build a Meteogram.
package builder

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/icodealot/noaa"
	"github.com/rickb777/date/period"
	"math"
	"reflect"
	"strings"
	"time"
)

var (
	forecastDuration = flag.Duration("forecast_duration", 48*time.Hour, "How far into the future should a forecast be.")
	labels           = []string{
		"Time", "Temperature", "Dewpoint", "MaxTemperature", "MinTemperature",
		"RelativeHumidity", "ApparentTemperature", "HeatIndex", "WindChill",
		"SkyCover", "WindDirection", "WindSpeed", "WindGust",
		"ProbabilityOfPrecipitation", "QuantitativePrecipitation",
	}
	weatherConditions = []string{
		"Coverage", "Weather", "Intensity",
	}
)

// CsvMeteogram returns a Metrogram in CSV format.
func CsvMeteogram(forecast *noaa.GridpointForecastResponse) (string, error) {
	records := [][]string{append(labels, weatherConditions...)}
	start := time.Now().Truncate(time.Hour)
	end := start.Add(*forecastDuration)
	for instant := start; !instant.After(end); instant = instant.Add(time.Hour) {
		record := []string{instant.Format(time.RFC3339)}
		// Excludes "Time"
		for _, k := range labels[1:] {
			f := reflect.ValueOf(*forecast).FieldByName(k).Interface().(noaa.GridpointForecastTimeSeries)
			v, err := valueAt(f, instant)
			if err != nil {
				return "", err
			}
			if math.IsNaN(v) {
				record = append(record, "")
			} else {
				record = append(record, fmt.Sprintf("%.04f", v))
			}
		}
		v, err := conditionsAt(*forecast, instant)
		for _, k := range weatherConditions {
			if err != nil {
				record = append(record, "")
			} else {
				record = append(record, reflect.ValueOf(*v).FieldByName(k).String())
			}
		}
		records = append(records, record)
		// Append units to CSV header.
		if instant == start {
			for i, k := range labels[1:] {
				f := reflect.ValueOf(*forecast).FieldByName(k).Interface().(noaa.GridpointForecastTimeSeries)
				if f.Uom != "" {
					unit := humanizeUnit(f.Uom)
					records[0][i+1] = fmt.Sprintf("%s (%s)", records[0][i+1], unit)
				}
			}
		}
	}
	b := new(bytes.Buffer)
	w := csv.NewWriter(b)
	if err := w.WriteAll(records); err != nil {
		return "", fmt.Errorf("could not write CSV: %q", err)
	}
	return b.String(), nil
}

func valueAt(g noaa.GridpointForecastTimeSeries, t time.Time) (float64, error) {
	for _, v := range g.Values {
		sections := strings.Split(v.ValidTime, "/")
		if len(sections) != 2 {
			return math.NaN(), fmt.Errorf("not a valid ISO8601 interval: %q", v.ValidTime)
		}
		start, err := time.Parse(time.RFC3339, sections[0])
		if err != nil {
			return math.NaN(), fmt.Errorf("could not parse time: %q", sections[0])
		}
		duration, err := period.Parse(sections[1], true)
		if err != nil {
			return math.NaN(), fmt.Errorf("could not parse ISO8601 duration: %q", sections[1])
		}
		end, _ := duration.AddTo(start)
		// t is equal to either start or end time or lies between the two.
		if (t.Equal(start) || t.Equal(end)) || (t.After(start) && t.Before(end)) {
			return v.Value, nil
		}
	}
	return math.NaN(), nil
}

type Conditions struct {
	Coverage  string
	Weather   string
	Intensity string
}

func conditionsAt(w noaa.GridpointForecastResponse, t time.Time) (*Conditions, error) {
	for _, v := range w.Weather.Values {
		sections := strings.Split(v.ValidTime, "/")
		if len(sections) != 2 {
			return nil, fmt.Errorf("not a valid ISO8601 interval: %q", v.ValidTime)
		}
		start, err := time.Parse(time.RFC3339, sections[0])
		if err != nil {
			return nil, fmt.Errorf("could not parse time: %q", sections[0])
		}
		duration, err := period.Parse(sections[1], true)
		if err != nil {
			return nil, fmt.Errorf("could not parse ISO8601 duration: %q", sections[1])
		}
		end, _ := duration.AddTo(start)
		// t is equal to either start or end time or lies between the two.
		if (t.Equal(start) || t.Equal(end)) || (t.After(start) && t.Before(end)) {
			return &Conditions{
				v.Value[0].Coverage,
				v.Value[0].Weather,
				v.Value[0].Intensity,
			}, nil
		}
	}
	return nil, fmt.Errorf("weather conditions at %q not found", t)
}

// humanizeUnit converts a subset of WMO units of measure used by the
// weather.gov API into a more human readable form. See http://codes.wmo.int/common/unit.
func humanizeUnit(unit string) string {
	switch unit {
	case "wmoUnit:degC":
		return "C"
	case "wmoUnit:percent":
		return "%"
	case "wmoUnit:degree_(angle)":
		return "angle"
	case "wmoUnit:km_h-1":
		return "km/h"
	case "wmoUnit:mm":
		return "mm"
	default:
		fmt.Printf("Unknown unit: %q", unit)
		return unit
	}
}
