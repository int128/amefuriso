package adapters

import (
	"fmt"
	"github.com/int128/amefuriso/domain"
	"github.com/int128/amefuriso/yolpweather"
)

func Weathers(resp *yolpweather.Response) ([]domain.Weather, error) {
	weathers := make([]domain.Weather, 0)
	for _, f := range resp.Payload.Feature {
		var weather domain.Weather
		c, err := f.Geometry.Coordinates.Parse()
		if err != nil {
			return nil, fmt.Errorf("invalid coordinates: %s", err)
		}
		weather.Coordinates = domain.Coordinates{
			Latitude:  c.Latitude,
			Longitude: c.Longitude,
		}

		for _, w := range f.Property.WeatherList.Weather {
			t, err := w.Date.Parse()
			if err != nil {
				return nil, fmt.Errorf("invalid date: %s", err)
			}
			r := domain.Rainfall{
				Time:   t,
				Amount: domain.RainfallMilliMeterPerHour(w.Rainfall),
			}
			switch w.Type {
			case yolpweather.Observation:
				weather.RainfallObservation = append(weather.RainfallObservation, r)
			case yolpweather.Forecast:
				weather.RainfallForecast = append(weather.RainfallForecast, r)
			default:
				return nil, fmt.Errorf("unknown weather type: %s", w.Type)
			}
		}

		weathers = append(weathers, weather)
	}
	return weathers, nil
}
