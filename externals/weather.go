package externals

import (
	"fmt"
	"net/http"

	"github.com/int128/amefurisobot/domain"
	"github.com/int128/go-yahoo-weather/weather"
	"github.com/pkg/errors"
)

type WeatherService struct {
	Client   *http.Client
	ClientID string
}

func (s *WeatherService) Get(locations []domain.Location) ([]domain.Weather, error) {
	req := weather.Request{
		IntervalMinutes: 5,
		PastHours:       1,
	}
	for _, location := range locations {
		req.Coordinates = append(req.Coordinates, weather.Coordinates{
			Latitude:  location.Coordinates.Latitude,
			Longitude: location.Coordinates.Longitude,
		})
	}
	c := weather.Client{
		Client:   s.Client,
		ClientID: s.ClientID,
	}
	resp, err := c.Get(&req)
	if err != nil {
		return nil, errors.Wrapf(err, "error while getting weather")
	}
	return weatherAdaptor(resp, locations)
}

func weatherAdaptor(resp *weather.Response, locations []domain.Location) ([]domain.Weather, error) {
	results := make([]domain.Weather, 0)
	for i, respFeature := range resp.Body.Feature {
		c, err := respFeature.Geometry.Coordinates.Parse()
		if err != nil {
			return nil, fmt.Errorf("invalid coordinates: %s", err)
		}
		result := domain.Weather{
			Location: domain.Location{
				Name: locations[i].Name,
				Coordinates: domain.Coordinates{
					Latitude:  c.Latitude,
					Longitude: c.Longitude,
				},
			},
		}
		for _, respWeather := range respFeature.Property.WeatherList.Weather {
			t, err := respWeather.Date.Parse()
			if err != nil {
				return nil, fmt.Errorf("invalid date: %s", err)
			}
			rainfall := domain.Rainfall{
				Time:   t,
				Amount: domain.RainfallMilliMeterPerHour(respWeather.Rainfall),
			}
			switch respWeather.Type {
			case "observation":
				result.RainfallObservation = append(result.RainfallObservation, rainfall)
			case "forecast":
				result.RainfallForecast = append(result.RainfallForecast, rainfall)
			default:
				return nil, fmt.Errorf("unknown weather type: %s", respWeather.Type)
			}
		}
		results = append(results, result)
	}
	return results, nil
}
