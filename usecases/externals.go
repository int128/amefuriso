package usecases

import (
	"context"

	"github.com/int128/amefurisobot/domain"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	Save(ctx context.Context, user domain.User) error
}

type SubscriptionRepository interface {
	FindByUserID(ctx context.Context, userID domain.UserID) ([]domain.Subscription, error)
	Save(ctx context.Context, userID domain.UserID, subscriptions []domain.Subscription) error
}

type PNGRepository interface {
	GetById(ctx context.Context, id string) ([]byte, error)
	Save(ctx context.Context, id string, b []byte) error
}

type WeatherService interface {
	Get(clientID domain.YahooClientID, locations []domain.Location) ([]domain.Weather, error)
}

type NotificationService interface {
	Send(destination domain.Notification, message domain.Message) error
}
