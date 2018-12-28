package gateways

import (
	"context"
	"time"

	"github.com/int128/amefuriso/domain"
)

//go:generate mockgen -destination mock_gateways/mock_gateways.go -package mock_gateways github.com/int128/amefuriso/gateways/interfaces UserRepository,SubscriptionRepository,PNGRepository,WeatherService,NotificationService

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
	RemoveOlderThan(ctx context.Context, t time.Time) (int, error)
	Save(ctx context.Context, image domain.Image) error
}

type ObservationOption int

const (
	NoObservation      = ObservationOption(0)
	OneHourObservation = ObservationOption(1)
)

type WeatherService interface {
	Get(ctx context.Context, clientID domain.YahooClientID, locations []domain.Location, observationOption ObservationOption) ([]domain.Weather, error)
}

type ForecastMessage struct {
	Forecast   domain.Forecast
	ImageURL   string
	WeatherURL string
}

type NotificationService interface {
	SendForecastMessage(ctx context.Context, recipient domain.Recipient, message ForecastMessage) error
}
