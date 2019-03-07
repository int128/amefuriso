package domain_test

import (
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/int128/amefuriso/domain"
	"github.com/int128/amefuriso/domain/testdata"
)

func TestNewForecast(t *testing.T) {
	t.Run("NoRain", func(t *testing.T) {
		weather := domain.Weather{
			Location: testdata.TokyoLocation,
			Observations: []domain.Event{
				{Time: baseTime.Add(-10 * time.Minute), Rainfall: 0},
			},
			Forecasts: []domain.Event{
				{Time: baseTime.Add(0 * time.Minute), Rainfall: 0},
				{Time: baseTime.Add(10 * time.Minute), Rainfall: 0},
				{Time: baseTime.Add(20 * time.Minute), Rainfall: 0},
			},
		}
		want := domain.Forecast{
			Location: testdata.TokyoLocation,
		}
		actual := domain.NewForecast(weather)
		if diff := deep.Equal(want, actual); diff != nil {
			t.Error(diff)
		}
	})

	t.Run("RainStart", func(t *testing.T) {
		weather := domain.Weather{
			Location: testdata.TokyoLocation,
			Observations: []domain.Event{
				{Time: baseTime.Add(-10 * time.Minute), Rainfall: 0},
			},
			Forecasts: []domain.Event{
				{Time: baseTime.Add(0 * time.Minute), Rainfall: 0},
				{Time: baseTime.Add(10 * time.Minute), Rainfall: 1},
				{Time: baseTime.Add(20 * time.Minute), Rainfall: 1},
			},
		}
		want := domain.Forecast{
			Location: testdata.TokyoLocation,
			RainWillStart: &domain.Event{
				Time: baseTime.Add(10 * time.Minute), Rainfall: 1,
			},
		}
		actual := domain.NewForecast(weather)
		if diff := deep.Equal(want, actual); diff != nil {
			t.Error(diff)
		}
	})

	t.Run("RainStop", func(t *testing.T) {
		weather := domain.Weather{
			Location: testdata.TokyoLocation,
			Observations: []domain.Event{
				{Time: baseTime.Add(-10 * time.Minute), Rainfall: 1},
			},
			Forecasts: []domain.Event{
				{Time: baseTime.Add(0 * time.Minute), Rainfall: 1},
				{Time: baseTime.Add(10 * time.Minute), Rainfall: 0},
				{Time: baseTime.Add(20 * time.Minute), Rainfall: 0},
			},
		}
		want := domain.Forecast{
			Location: testdata.TokyoLocation,
			RainWillStop: &domain.Event{
				Time: baseTime.Add(10 * time.Minute), Rainfall: 0,
			},
		}
		actual := domain.NewForecast(weather)
		if diff := deep.Equal(want, actual); diff != nil {
			t.Error(diff)
		}
	})

	t.Run("ForeverRain", func(t *testing.T) {
		weather := domain.Weather{
			Location: testdata.TokyoLocation,
			Observations: []domain.Event{
				{Time: baseTime.Add(-10 * time.Minute), Rainfall: 1},
			},
			Forecasts: []domain.Event{
				{Time: baseTime.Add(0 * time.Minute), Rainfall: 1},
				{Time: baseTime.Add(10 * time.Minute), Rainfall: 1},
				{Time: baseTime.Add(20 * time.Minute), Rainfall: 1},
			},
		}
		want := domain.Forecast{
			Location: testdata.TokyoLocation,
		}
		actual := domain.NewForecast(weather)
		if diff := deep.Equal(want, actual); diff != nil {
			t.Error(diff)
		}
	})
}
