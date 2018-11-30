package domain

import (
	"fmt"
	"time"
)

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

type Weather struct {
	Coordinates         Coordinates
	RainfallObservation []Rainfall
	RainfallForecast    []Rainfall
}

type RainfallMilliMeterPerHour float64

func (r RainfallMilliMeterPerHour) String() string {
	return fmt.Sprintf("%.2f mm/h", r)
}

type Rainfall struct {
	Time   time.Time
	Amount RainfallMilliMeterPerHour
}
