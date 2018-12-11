package main

import (
	"flag"
	"log"
	"os"

	"github.com/int128/amefuriso/core/domain"
	"github.com/int128/amefuriso/core/presenters/cli"
	"github.com/int128/amefuriso/externals"
	"github.com/pkg/errors"
)

type options struct {
	clientID string
}

func (o *options) run(locations []domain.Location) error {
	weatherService := externals.WeatherService{ClientID: o.clientID}
	weathers, err := weatherService.Get(locations)
	if err != nil {
		return errors.Wrapf(err, "error while getting weather")
	}
	for _, weather := range weathers {
		if err := cli.Draw(os.Stdout, weather); err != nil {
			return errors.Wrapf(err, "error while drawing weather")
		}
	}
	return nil
}

func main() {
	var location domain.Location
	flag.Float64Var(&location.Coordinates.Latitude, "lat", 35.663613, "Latitude")
	flag.Float64Var(&location.Coordinates.Longitude, "lon", 139.732293, "Longitude")
	var o options
	flag.StringVar(&o.clientID, "client-id", os.Getenv("YAHOO_CLIENT_ID"), "Yahoo Weather API Client ID")
	flag.Parse()

	if err := o.run([]domain.Location{location}); err != nil {
		log.Fatalf("Error: %s", err)
	}
}
