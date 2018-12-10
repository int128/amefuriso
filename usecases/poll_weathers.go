package usecases

import (
	"context"

	"github.com/int128/amefuriso/domain"
	"github.com/int128/amefuriso/externals"
	"github.com/pkg/errors"
)

type PollWeathers struct {
	SubscriptionRepository externals.SubscriptionRepository
	WeatherService         externals.WeatherService
	NotificationService    externals.NotificationService
}

func (u *PollWeathers) Do(ctx context.Context) error {
	subscriptions, err := u.SubscriptionRepository.FindAll(ctx)
	if err != nil {
		return errors.Wrapf(err, "error while getting subscriptions")
	}
	if len(subscriptions) == 0 {
		return errors.New("no subscription found")
	}
	var locations []domain.Location
	for _, subscription := range subscriptions {
		locations = append(locations, subscription.Location)
	}
	weathers, err := u.WeatherService.Get(locations)
	if err != nil {
		return errors.Wrapf(err, "error while getting weather")
	}
	for i, subscription := range subscriptions {
		weather := weathers[i]
		if !weather.IsRainingNow() && !weather.WillRainLater() {
			continue
		}
		if err := u.NotificationService.Send(ctx, subscription.Notification, weather); err != nil {
			return errors.Wrapf(err, "error while sending weather notification")
		}
	}
	return nil
}
