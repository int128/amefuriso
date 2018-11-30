package yolpweather

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Response struct {
	Payload Payload
	Expires time.Time
}

type Payload struct {
	ResultInfo struct {
		Count       int     `json:"Count"`
		Total       int     `json:"Total"`
		Start       int     `json:"Start"`
		Status      int     `json:"Status"`
		Latency     float64 `json:"Latency"`
		Description string  `json:"Description"`
		Copyright   string  `json:"Copyright"`
	} `json:"ResultInfo"`
	Feature []struct {
		ID       string `json:"Id"`
		Name     string `json:"Name"`
		Geometry struct {
			Type        string            `json:"Type"`
			Coordinates CoordinatesString `json:"Coordinates"`
		} `json:"Geometry"`
		Property struct {
			WeatherAreaCode int `json:"WeatherAreaCode"`
			WeatherList     struct {
				Weather []struct {
					Type     WeatherType `json:"Type"`
					Date     DateString  `json:"Date"`
					Rainfall float64     `json:"Rainfall"`
				} `json:"Weather"`
			} `json:"WeatherList"`
		} `json:"Property"`
	} `json:"Feature"`
}

type WeatherType string

const (
	Observation = WeatherType("observation")
	Forecast    = WeatherType("forecast")
)

// DateString represents a coordinates in YOLP format.
type CoordinatesString string

func (s CoordinatesString) Parse() (Coordinates, error) {
	p := strings.SplitN(string(s), ",", 2)
	if len(p) != 2 {
		return Coordinates{}, fmt.Errorf("wants comma separated 2 values but %d", len(p))
	}
	lat, lon := p[1], p[0]

	var c Coordinates
	var err error
	c.Latitude, err = strconv.ParseFloat(lat, 64)
	if err != nil {
		return Coordinates{}, fmt.Errorf("error while parsing latitude %s: %s", lat, err)
	}
	c.Longitude, err = strconv.ParseFloat(lon, 64)
	if err != nil {
		return Coordinates{}, fmt.Errorf("error while parsing longitude %s: %s", lon, err)
	}
	return c, nil
}

// DateString represents a time in YOLP format.
type DateString string

func (t DateString) Parse() (time.Time, error) {
	return time.Parse("200601021504", string(t))
}
