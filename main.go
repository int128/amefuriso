package main

import (
	"fmt"
	"log"
	"os"

	"github.com/int128/amefuriso/yolpweather"
)

func run() error {
	clientID := os.Getenv("YOLP_CLIENT_ID")

	roppongi := yolpweather.Coordinates{
		Latitude:  139.732293,
		Longitude: 35.663613,
	}

	c := yolpweather.New(clientID)
	resp, err := c.Get(&yolpweather.Request{
		Coordinates: []yolpweather.Coordinates{roppongi},
	})
	if err != nil {
		return fmt.Errorf("error while getting weather: %s", err)
	}
	log.Printf("%+v", resp)
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error: %s", err)
	}
}
