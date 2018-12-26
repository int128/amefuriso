package gateways

import (
	"context"
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/golang/mock/gomock"
	"github.com/int128/amefuriso/domain"
	"github.com/int128/amefuriso/domain/testdata"
	"github.com/int128/amefuriso/gateways/mock_infrastructure"
	"github.com/int128/go-yahoo-weather/weather"
)

func TestWeatherService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	client := mock_infrastructure.NewMockWeatherClient(ctrl)
	client.EXPECT().
		Get(ctx, "CLIENT_ID", weather.Request{
			Coordinates: []weather.Coordinates{
				{Latitude: 35.663613, Longitude: 139.732293},
				{Latitude: 41.7686738, Longitude: 140.728924},
			},
			IntervalMinutes: 5,
			PastHours:       1,
		}).
		Return([]weather.Weather{
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
		}, nil)

	service := WeatherService{Client: client}
	actual, err := service.Get(ctx,
		domain.YahooClientID("CLIENT_ID"),
		[]domain.Location{testdata.TokyoLocation, testdata.HakodateLocation},
		domain.OneHourObservation)
	if err != nil {
		t.Fatalf("Error while service.Get: %s", err)
	}
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
