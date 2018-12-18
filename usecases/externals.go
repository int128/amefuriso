package usecases

import (
	"context"

	"github.com/int128/amefurisobot/domain"
)

type UserRepository interface {
	FindById(ctx context.Context, id domain.UserID) (*domain.User, error)
	FindAll(ctx context.Context) ([]domain.User, error)
	Save(ctx context.Context, user domain.User) error
}

type SubscriptionRepository interface {
	FindBySubscriptionID(ctx context.Context, userID domain.UserID, subscriptionID domain.SubscriptionID) (*domain.Subscription, error)
	FindByUserID(ctx context.Context, userID domain.UserID) ([]domain.Subscription, error)
	Save(ctx context.Context, userID domain.UserID, subscriptions []domain.Subscription) error
}

type PNGRepository interface {
	FindById(ctx context.Context, id domain.ImageID) (*domain.Image, error)
	Save(ctx context.Context, image domain.Image) error
}

type WeatherService interface {
	Get(clientID domain.YahooClientID, locations []domain.Location) ([]domain.Weather, error)
}

type NotificationService interface {
	Send(destination domain.Notification, message domain.Message) error
}
