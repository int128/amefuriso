package usecases

import (
	"context"

	"github.com/int128/amefuriso/domain"
)

//go:generate mockgen -destination mock_usecases/mock_usecases.go github.com/int128/amefuriso/usecases IGetWeather,IGetImage,IPollWeathers,ICleanupImages,ISetup

type IGetWeather interface {
	Do(ctx context.Context, userID domain.UserID, subscriptionID domain.SubscriptionID) (*domain.Weather, error)
}

type IGetImage interface {
	Do(ctx context.Context, id domain.ImageID) (*domain.Image, error)
}

type IPollWeathers interface {
	Do(ctx context.Context, urlProviders URLProviders) error
}

type ICleanupImages interface {
	Do(ctx context.Context) error
}

type ISetup interface {
	Do(ctx context.Context) (*domain.User, error)
}
