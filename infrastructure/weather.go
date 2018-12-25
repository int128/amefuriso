package infrastructure

import (
	"context"

	"github.com/int128/go-yahoo-weather/weather"
	"github.com/pkg/errors"
	"google.golang.org/appengine/urlfetch"
)

type WeatherClient struct{}

func (c *WeatherClient) Get(ctx context.Context, clientID string, req weather.Request) ([]weather.Weather, error) {
	client := weather.Client{
		Client:   urlfetch.Client(ctx),
		ClientID: clientID,
	}
	resp, err := client.Get(&req)
	if err != nil {
		return nil, errors.Wrapf(err, "error while getting weather")
	}
	weathers, err := weather.Parse(resp)
	if err != nil {
		return nil, errors.Wrapf(err, "error while parsing weather response")
	}
	return weathers, nil
}
