package domain

import (
	"context"
)

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

type WeatherService interface {
	Get(ctx context.Context, clientID YahooClientID, locations []Location) ([]Weather, error)
}

type NotificationService interface {
	Send(ctx context.Context, destination Notification, message Message) error
}