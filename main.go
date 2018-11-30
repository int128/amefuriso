package main

import (
	"flag"
	"fmt"
	"github.com/int128/amefuriso/adapters"
	"github.com/int128/amefuriso/yolpweather"
	"log"
	"math"
	"strings"
)

type options struct {
	clientID  string
	latitude  float64
	longitude float64
}

func run(o options) error {
	req := yolpweather.Request{
		Coordinates: []yolpweather.Coordinates{{
			Latitude:  o.latitude,
			Longitude: o.longitude,
		}},
		IntervalMinutes: 5,
		PastHours:       1,
	}
	c := yolpweather.New(o.clientID)
	resp, err := c.Get(&req)
	if err != nil {
		return fmt.Errorf("error while getting weather: %s", err)
	}
	weathers, err := adapters.Weathers(resp)
	if err != nil {
		return fmt.Errorf("error while parsing response: %s", err)
	}
	for _, weather := range weathers {
		fmt.Printf("Weather at (%f, %f)\n", weather.Coordinates.Longitude, weather.Coordinates.Latitude)
		for _, rainfall := range weather.RainfallObservation {
			t := rainfall.Time.Format("15:04")
			mark := strings.Repeat("🌧 ", int(math.Ceil(float64(rainfall.Amount))))
			fmt.Printf("| %s | %5.2f mm/h | %s\n", t, rainfall.Amount, mark)
		}
		fmt.Println("|-------|------------|---")
		for _, rainfall := range weather.RainfallForecast {
			t := rainfall.Time.Format("15:04")
			mark := strings.Repeat("🌧 ", int(math.Ceil(float64(rainfall.Amount))))
			fmt.Printf("| %s | %5.2f mm/h | %s\n", t, rainfall.Amount, mark)
		}
	}
	return nil
}

func main() {
	var o options
	flag.StringVar(&o.clientID, "client-id", "", "YOLP Client ID")
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
