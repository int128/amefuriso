package gateways

import (
	"context"

	"github.com/int128/go-yahoo-weather/weather"
	"github.com/int128/slack"
)

//go:generate mockgen -destination mock_infrastructure/mock_infrastructure.go -package mock_infrastructure github.com/int128/amefurisobot/gateways WeatherClient,SlackClient

type WeatherClient interface {
	Get(ctx context.Context, clientID string, req weather.Request) ([]weather.Weather, error)
}

type SlackClient interface {
	Send(ctx context.Context, webhookURL string, message slack.Message) error
}
