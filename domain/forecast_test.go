package domain_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/int128/amefuriso/domain"
	"github.com/int128/amefuriso/domain/testdata"
)

func TestNewForecast(t *testing.T) {
	for i, c := range []struct {
		Weather  domain.Weather
		Forecast domain.Forecast
	}{
		{
			Weather: newWeather(0, 0, 0, 0),
		}, {
			Weather: newWeather(1, 0, 0, 0),
			Forecast: domain.Forecast{
				RainWillStop: &domain.Event{Time: baseTime.Add(0 * time.Minute), Rainfall: 0},
			},
		}, {
			Weather: newWeather(1, 1, 0, 0),
			Forecast: domain.Forecast{
				RainWillStop: &domain.Event{Time: baseTime.Add(10 * time.Minute), Rainfall: 0},
			},
		}, {
			Weather: newWeather(1, 1, 1, 0),
			Forecast: domain.Forecast{
				RainWillStop: &domain.Event{Time: baseTime.Add(20 * time.Minute), Rainfall: 0},
			},
		}, {
			Weather: newWeather(1, 1, 1, 1),
		}, {
			Weather: newWeather(0, 1, 1, 1),
			Forecast: domain.Forecast{
				RainWillStart: &domain.Event{Time: baseTime.Add(0 * time.Minute), Rainfall: 1},
			},
		}, {
			Weather: newWeather(0, 0, 1, 1),
			Forecast: domain.Forecast{
				RainWillStart: &domain.Event{Time: baseTime.Add(10 * time.Minute), Rainfall: 1},
			},
		}, {
			Weather: newWeather(0, 0, 0, 1),
			Forecast: domain.Forecast{
				RainWillStart: &domain.Event{Time: baseTime.Add(20 * time.Minute), Rainfall: 1},
			},
		}, {
			Weather: newWeather(0, 0, 0, 0),
		},
	} {
		// set Location here to reduce code
		c.Weather.Location = testdata.TokyoLocation
		c.Forecast.Location = testdata.TokyoLocation

		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			actual := domain.NewForecast(c.Weather)
			if diff := deep.Equal(c.Forecast, actual); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func Test_newWeather(t *testing.T) {
	want := domain.Weather{
		Observations: []domain.Event{
			{Time: baseTime.Add(-10 * time.Minute), Rainfall: 1},
		},
		Forecasts: []domain.Event{
			{Time: baseTime.Add(0 * time.Minute), Rainfall: 1},
			{Time: baseTime.Add(10 * time.Minute), Rainfall: 0},
			{Time: baseTime.Add(20 * time.Minute), Rainfall: 0},
		},
	}
	actual := newWeather(1, 1, 0, 0)
	if diff := deep.Equal(want, actual); diff != nil {
		t.Error(diff)
	}
}

func newWeather(observation float64, forecasts ...float64) domain.Weather {
	w := domain.Weather{
		Observations: []domain.Event{
			{Time: baseTime.Add(-10 * time.Minute), Rainfall: domain.RainfallMilliMeterPerHour(observation)},
		},
		Forecasts: []domain.Event{},
	}
	for i, forecast := range forecasts {
		w.Forecasts = append(w.Forecasts, domain.Event{
			Time:     baseTime.Add(10 * time.Minute * time.Duration(i)),
			Rainfall: domain.RainfallMilliMeterPerHour(forecast),
		})
	}
	return w
}
