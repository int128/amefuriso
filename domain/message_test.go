package domain_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/int128/amefurisobot/domain"
)

func TestNewForecastMessage(t *testing.T) {
	for i, c := range []struct {
		Weather         domain.Weather
		ForecastMessage interface{}
	}{
		{
			Weather: domain.Weather{
				Observations: []domain.Event{
					{Time: baseTime.Add(-10 * time.Minute), Rainfall: 1},
				},
				Forecasts: []domain.Event{
					{Time: baseTime.Add(0 * time.Minute), Rainfall: 1},
					{Time: baseTime.Add(10 * time.Minute), Rainfall: 0},
					{Time: baseTime.Add(20 * time.Minute), Rainfall: 0},
				},
			},
			ForecastMessage: domain.RainWillStopMessage{
				Event: domain.Event{Time: baseTime.Add(10 * time.Minute), Rainfall: 0},
			},
		}, {
			Weather: domain.Weather{
				Observations: []domain.Event{
					{Time: baseTime.Add(-10 * time.Minute), Rainfall: 0},
				},
				Forecasts: []domain.Event{
					{Time: baseTime.Add(0 * time.Minute), Rainfall: 0},
					{Time: baseTime.Add(10 * time.Minute), Rainfall: 1},
					{Time: baseTime.Add(20 * time.Minute), Rainfall: 0},
				},
			},
			ForecastMessage: domain.RainWillStartMessage{
				Event: domain.Event{Time: baseTime.Add(10 * time.Minute), Rainfall: 1},
			},
		}, {
			Weather: domain.Weather{
				Observations: []domain.Event{
					{Time: baseTime.Add(-10 * time.Minute), Rainfall: 0},
				},
				Forecasts: []domain.Event{
					{Time: baseTime.Add(0 * time.Minute), Rainfall: 0},
					{Time: baseTime.Add(10 * time.Minute), Rainfall: 0},
					{Time: baseTime.Add(20 * time.Minute), Rainfall: 0},
				},
			},
			ForecastMessage: domain.NoForecastMessage,
		}, {
			Weather: domain.Weather{
				Observations: []domain.Event{
					{Time: baseTime.Add(-10 * time.Minute), Rainfall: 1},
				},
				Forecasts: []domain.Event{
					{Time: baseTime.Add(0 * time.Minute), Rainfall: 1},
					{Time: baseTime.Add(10 * time.Minute), Rainfall: 1},
					{Time: baseTime.Add(20 * time.Minute), Rainfall: 1},
				},
			},
			ForecastMessage: domain.NoForecastMessage,
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			actual := domain.NewForecastMessage(c.Weather)
			if diff := deep.Equal(c.ForecastMessage, actual); diff != nil {
				t.Error(diff)
			}
		})
	}
}
