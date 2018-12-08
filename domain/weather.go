package domain

import (
	"fmt"
	"time"
)

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

type Location struct {
	Name        string
	Coordinates Coordinates
}

type Weather struct {
	Location            Location
	RainfallObservation []Rainfall
	RainfallForecast    []Rainfall
}

// IsRainingNow returns true if the last observation is positive.
func (w Weather) IsRainingNow() bool {
	if len(w.RainfallObservation) < 1 {
		return false
	}
	return w.RainfallObservation[len(w.RainfallObservation)-1].Amount > 0
}

// WillRainLater returns true if forecast have positive.
func (w Weather) WillRainLater() bool {
	for _, rainfall := range w.RainfallForecast {
		if rainfall.Amount > 0 {
			return true
		}
	}
	return false
}

type RainfallMilliMeterPerHour float64

func (r RainfallMilliMeterPerHour) String() string {
	return fmt.Sprintf("%.2f mm/h", r)
}

type Rainfall struct {
	Time   time.Time
	Amount RainfallMilliMeterPerHour
}
