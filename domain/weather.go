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
	Location     Location
	Observations []Event
	Forecasts    []Event
}

// IsRainingNow returns true if the last observation is positive.
func (w Weather) IsRainingNow() bool {
	if len(w.Observations) < 1 {
		return false
	}
	return w.Observations[len(w.Observations)-1].Rainfall > 0
}

// WillRainLater returns true if forecast have positive.
func (w Weather) WillRainLater() bool {
	for _, event := range w.Forecasts {
		if event.Rainfall > 0 {
			return true
		}
	}
	return false
}

type RainfallMilliMeterPerHour float64

func (r RainfallMilliMeterPerHour) String() string {
	return fmt.Sprintf("%.2f mm/h", r)
}

type Event struct {
	Time     time.Time
	Rainfall RainfallMilliMeterPerHour
}
