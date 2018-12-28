package usecases

import (
	"context"

	"github.com/int128/amefuriso/domain"
)

//go:generate mockgen -destination mock_usecases/mock_usecases.go -package mock_usecases github.com/int128/amefuriso/usecases/interfaces GetWeather,GetImage,PollWeathers,CleanupImages,Setup

type GetWeather interface {
	Do(ctx context.Context, userID domain.UserID, subscriptionID domain.SubscriptionID) (*domain.Weather, error)
}

type GetImage interface {
	Do(ctx context.Context, id domain.ImageID) (*domain.Image, error)
}

type ImageURLProvider func(id domain.ImageID) string
type WeatherURLProvider func(userID domain.UserID, subscriptionID domain.SubscriptionID) string
type URLProviders struct {
	ImageURLProvider   ImageURLProvider
	WeatherURLProvider WeatherURLProvider
}

type PollWeathers interface {
	Do(ctx context.Context, urlProviders URLProviders) error
}

type CleanupImages interface {
	Do(ctx context.Context) error
}

type Setup interface {
	Do(ctx context.Context) (*domain.User, error)
}
