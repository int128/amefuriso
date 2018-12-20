package usecases

import (
	"context"

	"github.com/int128/amefurisobot/domain"
)

//go:generate mockgen -destination mock_usecases/mock_usecases.go github.com/int128/amefurisobot/usecases IGetWeather,IGetImage,IPollWeathers,ISetup

type IGetWeather interface {
	Do(ctx context.Context, userID domain.UserID, subscriptionID domain.SubscriptionID) (*domain.Weather, error)
}

type IGetImage interface {
	Do(ctx context.Context, id domain.ImageID) (*domain.Image, error)
}

type IPollWeathers interface {
	Do(ctx context.Context, urlProviders URLProviders) error
}

type ISetup interface {
	Do(ctx context.Context) (*domain.User, error)
}
