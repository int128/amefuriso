package domain

import (
	"context"
)

//go:generate mockgen -destination mock_externals/mock_externals.go -package mock_externals github.com/int128/amefurisobot/domain UserRepository,SubscriptionRepository,PNGRepository,WeatherService,NotificationService

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

type NotificationService interface {
	Send(ctx context.Context, destination Notification, message Message) error
}
