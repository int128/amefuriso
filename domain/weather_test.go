package domain_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/int128/amefurisobot/domain"
	"github.com/int128/go-yahoo-weather/weather"
)

var baseTime = time.Date(2018, 12, 12, 13, 5, 0, 0, weather.Timezone)

func TestWeather_IsRainingNow(t *testing.T) {
	for i, c := range []struct {
		Weather      domain.Weather
		IsRainingNow bool
	}{
		{
			Weather: domain.Weather{
				Observations: []domain.Event{
					{Time: baseTime.Add(0 * time.Minute), Rainfall: 0.35},
					{Time: baseTime.Add(10 * time.Minute), Rainfall: 0.45},
					{Time: baseTime.Add(20 * time.Minute), Rainfall: 1.15},
				},
			},
			IsRainingNow: true,
		}, {
			Weather: domain.Weather{
				Observations: []domain.Event{
					{Time: baseTime.Add(0 * time.Minute), Rainfall: 0.35},
					{Time: baseTime.Add(10 * time.Minute), Rainfall: 0.45},
					{Time: baseTime.Add(20 * time.Minute), Rainfall: 0},
				},
			},
			IsRainingNow: false,
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			actual := c.Weather.IsRainingNow()
			if actual != c.IsRainingNow {
				t.Errorf("IsRainingNow wants %v but %v", c.IsRainingNow, actual)
			}
		})
	}
}

func TestWeather_FindRainStarts(t *testing.T) {
	for i, c := range []struct {
		Weather domain.Weather
		Start   *domain.Event
	}{
		{
			Weather: domain.Weather{
				Forecasts: []domain.Event{
					{Time: baseTime.Add(0 * time.Minute), Rainfall: 0},
					{Time: baseTime.Add(10 * time.Minute), Rainfall: 0},
					{Time: baseTime.Add(20 * time.Minute), Rainfall: 0},
				},
			},
			Start: nil,
		}, {
			Weather: domain.Weather{
				Forecasts: []domain.Event{
					{Time: baseTime.Add(0 * time.Minute), Rainfall: 1},
					{Time: baseTime.Add(10 * time.Minute), Rainfall: 0},
					{Time: baseTime.Add(20 * time.Minute), Rainfall: 0},
				},
			},
			Start: &domain.Event{Time: baseTime.Add(0 * time.Minute), Rainfall: 1},
		}, {
			Weather: domain.Weather{
				Forecasts: []domain.Event{
					{Time: baseTime.Add(0 * time.Minute), Rainfall: 0},
					{Time: baseTime.Add(10 * time.Minute), Rainfall: 1},
					{Time: baseTime.Add(20 * time.Minute), Rainfall: 0},
				},
			},
			Start: &domain.Event{Time: baseTime.Add(10 * time.Minute), Rainfall: 1},
		}, {
			Weather: domain.Weather{
				Forecasts: []domain.Event{
					{Time: baseTime.Add(0 * time.Minute), Rainfall: 0},
					{Time: baseTime.Add(10 * time.Minute), Rainfall: 0},
					{Time: baseTime.Add(20 * time.Minute), Rainfall: 1},
				},
			},
			Start: &domain.Event{Time: baseTime.Add(20 * time.Minute), Rainfall: 1},
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			actual := c.Weather.FindRainStarts()
			if diff := deep.Equal(c.Start, actual); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestWeather_FindRainStops(t *testing.T) {
	for i, c := range []struct {
		Weather domain.Weather
		Stop    *domain.Event
	}{
		{
			Weather: domain.Weather{
				Forecasts: []domain.Event{
					{Time: baseTime.Add(0 * time.Minute), Rainfall: 0},
					{Time: baseTime.Add(10 * time.Minute), Rainfall: 0},
					{Time: baseTime.Add(20 * time.Minute), Rainfall: 0},
				},
			},
			Stop: &domain.Event{Time: baseTime.Add(0 * time.Minute), Rainfall: 0},
		}, {
			Weather: domain.Weather{
				Forecasts: []domain.Event{
					{Time: baseTime.Add(0 * time.Minute), Rainfall: 1},
					{Time: baseTime.Add(10 * time.Minute), Rainfall: 0},
					{Time: baseTime.Add(20 * time.Minute), Rainfall: 0},
				},
			},
			Stop: &domain.Event{Time: baseTime.Add(10 * time.Minute), Rainfall: 0},
		}, {
			Weather: domain.Weather{
				Forecasts: []domain.Event{
					{Time: baseTime.Add(0 * time.Minute), Rainfall: 0},
					{Time: baseTime.Add(10 * time.Minute), Rainfall: 1},
					{Time: baseTime.Add(20 * time.Minute), Rainfall: 0},
				},
			},
			Stop: &domain.Event{Time: baseTime.Add(0 * time.Minute), Rainfall: 0},
		}, {
			Weather: domain.Weather{
				Forecasts: []domain.Event{
					{Time: baseTime.Add(0 * time.Minute), Rainfall: 1},
					{Time: baseTime.Add(10 * time.Minute), Rainfall: 1},
					{Time: baseTime.Add(20 * time.Minute), Rainfall: 0},
				},
			},
			Stop: &domain.Event{Time: baseTime.Add(20 * time.Minute), Rainfall: 0},
		}, {
			Weather: domain.Weather{
				Forecasts: []domain.Event{
					{Time: baseTime.Add(0 * time.Minute), Rainfall: 1},
					{Time: baseTime.Add(10 * time.Minute), Rainfall: 1},
					{Time: baseTime.Add(20 * time.Minute), Rainfall: 1},
				},
			},
			Stop: nil,
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			actual := c.Weather.FindRainStops()
			if diff := deep.Equal(c.Stop, actual); diff != nil {
				t.Error(diff)
			}
		})
	}
}
