package usecases

import (
	"context"
	"github.com/int128/amefurisobot/domain"
	"github.com/pkg/errors"
)

type GetWeather struct {
	UserRepository         domain.UserRepository
	SubscriptionRepository domain.SubscriptionRepository
	WeatherService         domain.WeatherService
}

func (u *GetWeather) Do(ctx context.Context, userID domain.UserID, subscriptionID domain.SubscriptionID) (*domain.Weather, error) {
	user, err := u.UserRepository.FindById(ctx, userID)
	if err != nil {
		if domain.IsErrNoSuchUser(err) {
			return nil, err
		}
		return nil, errors.Wrapf(err, "error while finding user")
	}
	subscription, err := u.SubscriptionRepository.FindBySubscriptionID(ctx, userID, subscriptionID)
	if err != nil {
		if domain.IsErrNoSuchSubscription(err) {
			return nil, err
		}
		return nil, errors.Wrapf(err, "error while finding subscription")
	}
	weathers, err := u.WeatherService.Get(ctx, user.YahooClientID, []domain.Location{subscription.Location}, domain.OneHourObservation)
	if err != nil {
		return nil, errors.Wrapf(err, "error while fetching weather")
	}
	if len(weathers) != 1 {
		return nil, errors.Errorf("len(weathers) wants 1 but %d", len(weathers))
	}
	return &weathers[0], nil
}
