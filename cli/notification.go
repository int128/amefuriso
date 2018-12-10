package main

import (
	"context"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/int128/amefuriso/core/domain"
)

type notificationService struct{}

func (e *notificationService) Send(ctx context.Context, notification domain.Slack, weather domain.Weather) error {
	for _, rainfall := range weather.RainfallObservation {
		t := rainfall.Time.Format("15:04")
		mark := strings.Repeat("🌧 ", int(math.Ceil(float64(rainfall.Amount))))
		fmt.Printf("| %s |         | %5.2f mm/h | %s\n", t, rainfall.Amount, mark)
	}
	for _, rainfall := range weather.RainfallForecast {
		t := rainfall.Time.Format("15:04")
		d := -time.Since(rainfall.Time).Minutes()
		mark := strings.Repeat("🌧 ", int(math.Ceil(float64(rainfall.Amount))))
		fmt.Printf("| %s | %+3.0f min | %5.2f mm/h | %s\n", t, d, rainfall.Amount, mark)
	}
	return nil
}
