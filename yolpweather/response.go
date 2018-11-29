package yolpweather

import (
	"fmt"
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
	var c Coordinates
	_, err := fmt.Sscanf("%f,%f", string(s), &c.Latitude, &c.Longitude)
	if err != nil {
		return Coordinates{}, fmt.Errorf("could not parse coordinates %s: %s", s, err)
	}
	return c, nil
}

// DateString represents a time in YOLP format.
type DateString string

func (t DateString) Parse() (time.Time, error) {
	return time.Parse("200601021504", string(t))
}
