package domain

import (
	"context"
	"time"
)

//go:generate mockgen -destination mock_gateways/mock_gateways.go -package mock_gateways github.com/int128/amefuriso/domain UserRepository,SubscriptionRepository,PNGRepository,WeatherService,NotificationService

type UserRepository interface {
	FindById(ctx context.Context, id UserID) (*User, error)
	FindAll(ctx context.Context) ([]User, error)
	Save(ctx context.Context, user User) error
}

type SubscriptionRepository interface {
	FindBySubscriptionID(ctx context.Context, userID UserID, subscriptionID SubscriptionID) (*Subscription, error)
	FindByUserID(ctx context.Context, userID UserID) ([]Subscription, error)
	Save(ctx context.Context, userID UserID, subscriptions []Subscription) error
}

type PNGRepository interface {
	FindById(ctx context.Context, id ImageID) (*Image, error)
	RemoveOlderThan(ctx context.Context, t time.Time) (int, error)
	Save(ctx context.Context, image Image) error
}

type ObservationOption int

const (
	NoObservation      = ObservationOption(0)
	OneHourObservation = ObservationOption(1)
)

type WeatherService interface {
	Get(ctx context.Context, clientID YahooClientID, locations []Location, observationOption ObservationOption) ([]Weather, error)
}

type ForecastMessage struct {
	Forecast   Forecast
	ImageURL   string
	WeatherURL string
}

type NotificationService interface {
	SendForecastMessage(ctx context.Context, recipient Recipient, message ForecastMessage) error
}
