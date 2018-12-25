package domain_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/int128/amefurisobot/domain"
	"github.com/int128/amefurisobot/domain/testdata"
)

func TestNewForecast(t *testing.T) {
	for i, c := range []struct {
		Weather         domain.Weather
		ForecastMessage interface{}
	}{
		{
			Weather: domain.Weather{
				Location: testdata.TokyoLocation,
				Observations: []domain.Event{
					{Time: baseTime.Add(-10 * time.Minute), Rainfall: 1},
				},
				Forecasts: []domain.Event{
					{Time: baseTime.Add(0 * time.Minute), Rainfall: 1},
					{Time: baseTime.Add(10 * time.Minute), Rainfall: 0},
					{Time: baseTime.Add(20 * time.Minute), Rainfall: 0},
				},
			},
			ForecastMessage: domain.Forecast{
				Location:     testdata.TokyoLocation,
				RainWillStop: &domain.Event{Time: baseTime.Add(10 * time.Minute), Rainfall: 0},
			},
		}, {
			Weather: domain.Weather{
				Location: testdata.TokyoLocation,
				Observations: []domain.Event{
					{Time: baseTime.Add(-10 * time.Minute), Rainfall: 0},
				},
				Forecasts: []domain.Event{
					{Time: baseTime.Add(0 * time.Minute), Rainfall: 0},
					{Time: baseTime.Add(10 * time.Minute), Rainfall: 1},
					{Time: baseTime.Add(20 * time.Minute), Rainfall: 0},
				},
			},
			ForecastMessage: domain.Forecast{
				Location:      testdata.TokyoLocation,
				RainWillStart: &domain.Event{Time: baseTime.Add(10 * time.Minute), Rainfall: 1},
			},
		}, {
			Weather: domain.Weather{
				Location: testdata.TokyoLocation,
				Observations: []domain.Event{
					{Time: baseTime.Add(-10 * time.Minute), Rainfall: 0},
				},
				Forecasts: []domain.Event{
					{Time: baseTime.Add(0 * time.Minute), Rainfall: 0},
					{Time: baseTime.Add(10 * time.Minute), Rainfall: 0},
					{Time: baseTime.Add(20 * time.Minute), Rainfall: 0},
				},
			},
			ForecastMessage: domain.Forecast{
				Location: testdata.TokyoLocation,
			},
		}, {
			Weather: domain.Weather{
				Location: testdata.TokyoLocation,
				Observations: []domain.Event{
					{Time: baseTime.Add(-10 * time.Minute), Rainfall: 1},
				},
				Forecasts: []domain.Event{
					{Time: baseTime.Add(0 * time.Minute), Rainfall: 1},
					{Time: baseTime.Add(10 * time.Minute), Rainfall: 1},
					{Time: baseTime.Add(20 * time.Minute), Rainfall: 1},
				},
			},
			ForecastMessage: domain.Forecast{
				Location: testdata.TokyoLocation,
			},
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			actual := domain.NewForecast(c.Weather)
			if diff := deep.Equal(c.ForecastMessage, actual); diff != nil {
				t.Error(diff)
			}
		})
	}
}
