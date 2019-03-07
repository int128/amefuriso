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

// FindRainStarts returns the first event which it and next are non-zero.
func (w Weather) FindRainStarts() *Event {
	if len(w.Forecasts) < 2 {
		return nil // forecasts should have 2 or more
	}
	for i := 0; i < len(w.Forecasts)-1; i++ {
		f0 := w.Forecasts[i]
		f1 := w.Forecasts[i+1]
		if f0.Rainfall > 0 && f1.Rainfall > 0 {
			return &w.Forecasts[i]
		}
	}
	return nil
}

// FindRainStops returns the first event which it and next are zero.
func (w Weather) FindRainStops() *Event {
	if len(w.Forecasts) < 2 {
		return nil // forecasts should have 2 or more
	}
	for i := 0; i < len(w.Forecasts)-1; i++ {
		f0 := w.Forecasts[i]
		f1 := w.Forecasts[i+1]
		if f0.Rainfall == 0 && f1.Rainfall == 0 {
			return &w.Forecasts[i]
		}
	}
	return nil
}

type RainfallMilliMeterPerHour float64

func (r RainfallMilliMeterPerHour) String() string {
	return fmt.Sprintf("%.2f mm/h", r)
}

type Event struct {
	Time     time.Time
	Rainfall RainfallMilliMeterPerHour
}
