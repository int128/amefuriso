package main

import (
	"flag"
	"fmt"
	"github.com/int128/amefuriso/yolpweather"
	"log"
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
	}
	c := yolpweather.New(o.clientID)
	resp, err := c.Get(&req)
	if err != nil {
		return fmt.Errorf("error while getting weather: %s", err)
	}

	for _, f := range resp.Payload.Feature {
		coordinates, err := f.Geometry.Coordinates.Parse()
		if err != nil {
			log.Printf("Got invalid coordinates: %s", err)
		}
		for _, w := range f.Property.WeatherList.Weather {
			t, err := w.Date.Parse()
			if err != nil {
				log.Printf("Got invalid date: %s", err)
			}
			log.Printf("@%v\t%s\t%s\t%.2f mm", coordinates, w.Type, t.Format("15:04"), w.Rainfall)
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
