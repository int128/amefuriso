package adapters

import (
	"fmt"
	"github.com/int128/amefuriso/domain"
	"github.com/int128/go-yahoo-weather/weather"
)

func Weathers(resp *weather.Response) ([]domain.Weather, error) {
	results := make([]domain.Weather, 0)
	for _, respFeature := range resp.Body.Feature {
		c, err := respFeature.Geometry.Coordinates.Parse()
		if err != nil {
			return nil, fmt.Errorf("invalid coordinates: %s", err)
		}
		result := domain.Weather{
			Coordinates: domain.Coordinates{
				Latitude:  c.Latitude,
				Longitude: c.Longitude,
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
