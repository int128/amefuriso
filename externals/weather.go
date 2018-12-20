package externals

import (
	"context"

	"google.golang.org/appengine/urlfetch"

	"github.com/int128/amefurisobot/domain"
	"github.com/int128/go-yahoo-weather/weather"
	"github.com/pkg/errors"
)

type WeatherService struct{}

func (s *WeatherService) Get(ctx context.Context, clientID domain.YahooClientID, locations []domain.Location, observationOption domain.ObservationOption) ([]domain.Weather, error) {
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
	c := weather.Client{Client: urlfetch.Client(ctx), ClientID: string(clientID)}
	resp, err := c.Get(&req)
	if err != nil {
		return nil, errors.Wrapf(err, "error while getting weather")
	}
	weathers, err := weather.Parse(resp)
	if err != nil {
		return nil, errors.Wrapf(err, "error while parsing weather response")
	}
	return weatherAdaptor(weathers, locations), nil
}

func weatherAdaptor(weathers []weather.Weather, locations []domain.Location) []domain.Weather {
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
	return results
}
