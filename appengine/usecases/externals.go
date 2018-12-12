package usecases

import (
	"context"

	"github.com/int128/amefurisobot/core/domain"
)

type SubscriptionRepository interface {
	FindAll(ctx context.Context) ([]domain.Subscription, error)
}

type PNGRepository interface {
	GetById(ctx context.Context, id string) ([]byte, error)
	Save(ctx context.Context, id string, b []byte) error
}

type WeatherService interface {
	Get(locations []domain.Location) ([]domain.Weather, error)
}

type SlackService interface {
	Send(destination domain.Slack, message domain.Message) error
}
