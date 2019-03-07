package domain_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/int128/amefuriso/domain"
	"github.com/int128/go-yahoo-weather/weather"
)

var baseTime = time.Date(2018, 12, 12, 13, 5, 0, 0, weather.Timezone)

func makeEvents(rainfalls ...domain.RainfallMilliMeterPerHour) []domain.Event {
	e := make([]domain.Event, len(rainfalls))
	for i, rainfall := range rainfalls {
		e[i] = domain.Event{
			Time:     baseTime.Add(time.Duration(i) * 10 * time.Minute),
			Rainfall: rainfall,
		}
	}
	return e
}

func TestWeather_IsRainingNow(t *testing.T) {
	for i, c := range []struct {
		Weather      domain.Weather
		IsRainingNow bool
	}{
		{
			Weather: domain.Weather{
				Observations: makeEvents(0.35, 0.45, 1.15),
			},
			IsRainingNow: true,
		}, {
			Weather: domain.Weather{
				Observations: makeEvents(0.35, 0.45, 0),
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
		Forecasts []domain.Event
		Start     *domain.Event
	}{
		{
			Forecasts: makeEvents(0, 0, 0),
		}, {
			Forecasts: makeEvents(0, 0, 1),
			Start:     &domain.Event{Time: baseTime.Add(20 * time.Minute), Rainfall: 1},
		}, {
			Forecasts: makeEvents(0, 1, 0),
			Start:     &domain.Event{Time: baseTime.Add(10 * time.Minute), Rainfall: 1},
		}, {
			Forecasts: makeEvents(0, 1, 1),
			Start:     &domain.Event{Time: baseTime.Add(10 * time.Minute), Rainfall: 1},
		}, {
			Forecasts: makeEvents(1, 1, 1),
			Start:     &domain.Event{Time: baseTime, Rainfall: 1},
		}, {
			Forecasts: makeEvents(1, 1, 0),
			Start:     &domain.Event{Time: baseTime, Rainfall: 1},
		}, {
			Forecasts: makeEvents(1, 0, 1),
			Start:     &domain.Event{Time: baseTime, Rainfall: 1},
		}, {
			Forecasts: makeEvents(1, 0, 0),
			Start:     &domain.Event{Time: baseTime, Rainfall: 1},
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			w := domain.Weather{Forecasts: c.Forecasts}
			actual := w.FindRainStarts()
			if diff := deep.Equal(c.Start, actual); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestWeather_FindRainStops(t *testing.T) {
	for i, c := range []struct {
		Forecasts []domain.Event
		Stop      *domain.Event
	}{
		{
			Forecasts: makeEvents(0, 0, 0),
			Stop:      &domain.Event{Time: baseTime},
		}, {
			Forecasts: makeEvents(0, 0, 1),
			Stop:      &domain.Event{Time: baseTime},
		}, {
			Forecasts: makeEvents(0, 1, 0),
			Stop:      &domain.Event{Time: baseTime},
		}, {
			Forecasts: makeEvents(0, 1, 1),
			Stop:      &domain.Event{Time: baseTime},
		}, {
			Forecasts: makeEvents(1, 1, 1),
		}, {
			Forecasts: makeEvents(1, 1, 0),
			Stop:      &domain.Event{Time: baseTime.Add(20 * time.Minute)},
		}, {
			Forecasts: makeEvents(1, 0, 1),
			Stop:      &domain.Event{Time: baseTime.Add(10 * time.Minute)},
		}, {
			Forecasts: makeEvents(1, 0, 0),
			Stop:      &domain.Event{Time: baseTime.Add(10 * time.Minute)},
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			w := domain.Weather{Forecasts: c.Forecasts}
			actual := w.FindRainStops()
			if diff := deep.Equal(c.Stop, actual); diff != nil {
				t.Error(diff)
			}
		})
	}
}
