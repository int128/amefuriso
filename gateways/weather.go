package gateways

import (
	"context"

	"github.com/int128/amefuriso/domain"
	"github.com/int128/amefuriso/gateways/interfaces"
	"github.com/int128/amefuriso/infrastructure/interfaces"
	"github.com/int128/go-yahoo-weather/weather"
	"github.com/pkg/errors"
)

type WeatherService struct {
	Client infrastructure.WeatherClient
}

func (s *WeatherService) Get(ctx context.Context, clientID domain.YahooClientID, locations []domain.Location, observationOption gateways.ObservationOption) ([]domain.Weather, error) {
	req := weather.Request{
		IntervalMinutes: 5,
		PastHours:       int(observationOption),
	}
	for _, location := range locations {
		req.Coordinates = append(req.Coordinates, weather.Coordinates{
			Latitude:  location.Coordinates.Latitude,
			Longitude: location.Coordinates.Longitude,
		})
	}
	weathers, err := s.Client.Get(ctx, string(clientID), req)
	if err != nil {
		return nil, errors.Wrapf(err, "error while getting weather")
	}
	results := make([]domain.Weather, 0)
	for i, w := range weathers {
		result := domain.Weather{
			Location: locations[i],
		}
		for _, event := range w.Events {
			rainfall := domain.Event{
				Time:     event.Time,
				Rainfall: domain.RainfallMilliMeterPerHour(event.Rainfall),
			}
			if event.Forecast {
				result.Forecasts = append(result.Forecasts, rainfall)
			} else {
				result.Observations = append(result.Observations, rainfall)
			}
		}
		results = append(results, result)
	}
	return results, nil
}
