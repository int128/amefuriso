package externals

import (
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/int128/amefurisobot/domain"
	"github.com/int128/amefurisobot/domain/testdata"
	"github.com/int128/go-yahoo-weather/weather"
)

func Test_weatherAdaptor(t *testing.T) {
	weathers := []weather.Weather{
		{
			Coordinates: weather.Coordinates{Latitude: 35.663613, Longitude: 139.73229},
			Events: []weather.Event{
				{Time: time.Date(2018, 12, 12, 13, 5, 0, 0, weather.Timezone)},
				{Time: time.Date(2018, 12, 12, 13, 15, 0, 0, weather.Timezone), Forecast: true},
				{Time: time.Date(2018, 12, 12, 13, 25, 0, 0, weather.Timezone), Forecast: true},
				{Time: time.Date(2018, 12, 12, 13, 35, 0, 0, weather.Timezone), Forecast: true},
				{Time: time.Date(2018, 12, 12, 13, 45, 0, 0, weather.Timezone), Forecast: true},
				{Time: time.Date(2018, 12, 12, 13, 55, 0, 0, weather.Timezone), Forecast: true},
				{Time: time.Date(2018, 12, 12, 14, 5, 0, 0, weather.Timezone), Forecast: true},
			},
		}, {
			Coordinates: weather.Coordinates{Latitude: 41.768674, Longitude: 140.72892},
			Events: []weather.Event{
				{Time: time.Date(2018, 12, 12, 13, 5, 0, 0, weather.Timezone), Rainfall: 0.35},
				{Time: time.Date(2018, 12, 12, 13, 15, 0, 0, weather.Timezone), Rainfall: 0.45},
				{Time: time.Date(2018, 12, 12, 13, 25, 0, 0, weather.Timezone), Rainfall: 1.15},
				{Time: time.Date(2018, 12, 12, 13, 35, 0, 0, weather.Timezone), Rainfall: 0.45},
				{Time: time.Date(2018, 12, 12, 13, 45, 0, 0, weather.Timezone), Rainfall: 1.85},
				{Time: time.Date(2018, 12, 12, 13, 55, 0, 0, weather.Timezone), Forecast: true},
				{Time: time.Date(2018, 12, 12, 14, 5, 0, 0, weather.Timezone), Forecast: true},
			},
		},
	}
	locations := []domain.Location{testdata.TokyoLocation, testdata.HakodateLocation}
	actual := weatherAdaptor(weathers, locations)
	want := []domain.Weather{
		{
			Location: testdata.TokyoLocation,
			Observations: []domain.Event{
				{Time: time.Date(2018, 12, 12, 13, 5, 0, 0, weather.Timezone)},
			},
			Forecasts: []domain.Event{
				{Time: time.Date(2018, 12, 12, 13, 15, 0, 0, weather.Timezone)},
				{Time: time.Date(2018, 12, 12, 13, 25, 0, 0, weather.Timezone)},
				{Time: time.Date(2018, 12, 12, 13, 35, 0, 0, weather.Timezone)},
				{Time: time.Date(2018, 12, 12, 13, 45, 0, 0, weather.Timezone)},
				{Time: time.Date(2018, 12, 12, 13, 55, 0, 0, weather.Timezone)},
				{Time: time.Date(2018, 12, 12, 14, 5, 0, 0, weather.Timezone)},
			},
		}, {
			Location: testdata.HakodateLocation,
			Observations: []domain.Event{
				{Time: time.Date(2018, 12, 12, 13, 5, 0, 0, weather.Timezone), Rainfall: 0.35},
				{Time: time.Date(2018, 12, 12, 13, 15, 0, 0, weather.Timezone), Rainfall: 0.45},
				{Time: time.Date(2018, 12, 12, 13, 25, 0, 0, weather.Timezone), Rainfall: 1.15},
				{Time: time.Date(2018, 12, 12, 13, 35, 0, 0, weather.Timezone), Rainfall: 0.45},
				{Time: time.Date(2018, 12, 12, 13, 45, 0, 0, weather.Timezone), Rainfall: 1.85},
			},
			Forecasts: []domain.Event{
				{Time: time.Date(2018, 12, 12, 13, 55, 0, 0, weather.Timezone)},
				{Time: time.Date(2018, 12, 12, 14, 5, 0, 0, weather.Timezone)},
			},
		},
	}
	if diff := deep.Equal(want, actual); diff != nil {
		t.Error(diff)
	}
}
