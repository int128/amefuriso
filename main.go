package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"github.com/int128/amefuriso/domain"
	"github.com/int128/amefuriso/externals"
	"github.com/int128/go-yahoo-weather/weather"
)

type options struct {
	clientID  string
	latitude  float64
	longitude float64
}

func run(o options) error {
	weatherService := externals.WeatherService{
		Client: weather.NewClient(o.clientID),
	}
	weathers, err := weatherService.Get([]domain.Location{{
		Coordinates: domain.Coordinates{
			Latitude:  o.latitude,
			Longitude: o.longitude,
		},
	}})
	if err != nil {
		return fmt.Errorf("error while getting weather: %s", err)
	}
	for _, w := range weathers {
		for _, rainfall := range w.RainfallObservation {
			t := rainfall.Time.Format("15:04")
			mark := strings.Repeat("🌧 ", int(math.Ceil(float64(rainfall.Amount))))
			fmt.Printf("| %s |         | %5.2f mm/h | %s\n", t, rainfall.Amount, mark)
		}
		for _, rainfall := range w.RainfallForecast {
			t := rainfall.Time.Format("15:04")
			d := -time.Since(rainfall.Time).Minutes()
			mark := strings.Repeat("🌧 ", int(math.Ceil(float64(rainfall.Amount))))
			fmt.Printf("| %s | %+3.0f min | %5.2f mm/h | %s\n", t, d, rainfall.Amount, mark)
		}
	}
	return nil
}

func main() {
	var o options
	flag.StringVar(&o.clientID, "client-id", os.Getenv("YAHOO_CLIENT_ID"), "Yahoo API Client ID")
	flag.Float64Var(&o.latitude, "lat", 35.663613, "Latitude")
	flag.Float64Var(&o.longitude, "lon", 139.732293, "Longitude")
	flag.Parse()
	if o.clientID == "" {
		log.Fatalf("Run with -client-id option")
	}
	if err := run(o); err != nil {
		log.Fatalf("Error: %s", err)
	}
}
