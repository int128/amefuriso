package externals

import (
	"context"

	"github.com/int128/amefuriso/core/domain"
)

type SubscriptionRepository interface {
	FindAll(ctx context.Context) ([]domain.Subscription, error)
}

type WeatherService interface {
	Get(locations []domain.Location) ([]domain.Weather, error)
}

type NotificationService interface {
	Send(ctx context.Context, notification domain.Slack, weather domain.Weather) error
}
