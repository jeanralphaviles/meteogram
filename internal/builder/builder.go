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
)

// CsvMeteogram returns a Metrogram in CSV format.
func CsvMeteogram(forecast *noaa.GridpointForecastResponse) (string, error) {
	records := [][]string{
		{"Time", "Temperature", "RelativeHumidity", "Dewpoint", "HeatIndex", "WindChill", "WindSpeed", "WindDirection", "WindGust", "SkyCover", "ProbabilityOfPrecipitation"},
	}
	start := time.Now().Truncate(time.Hour)
	end := start.Add(*forecastDuration)
	for instant := start; !instant.After(end); instant = instant.Add(time.Hour) {
		record := []string{instant.Format(time.RFC3339)}
		for _, k := range records[0][1:] {
			f := reflect.ValueOf(*forecast).FieldByName(k)
			v, err := valueAt(f.Interface().(noaa.GridpointForecastTimeSeries), instant)
			if err != nil {
				return "", err
			}
			record = append(record, fmt.Sprintf("%.04f", v))
		}
		records = append(records, record)
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
			return math.NaN(), fmt.Errorf("could not ISO8601 duration: %q", sections[1])
		}
		end, _ := duration.AddTo(start)
		// t is equal to either start or end time or lies between the two.
		if (t.Equal(start) || t.Equal(end)) || (t.After(start) && t.Before(end)) {
			return v.Value, nil
		}
	}
	return math.NaN(), fmt.Errorf("value at %q not found", t)
}
