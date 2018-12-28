package usecases

import (
	"bytes"
	"context"
	"image/png"

	"github.com/int128/amefuriso/domain"
	"github.com/int128/amefuriso/domain/chart"
	"github.com/int128/amefuriso/gateways/interfaces"
	"github.com/int128/amefuriso/usecases/interfaces"
	"github.com/pkg/errors"
)

type PollWeathers struct {
	UserRepository         gateways.UserRepository
	SubscriptionRepository gateways.SubscriptionRepository
	WeatherService         gateways.WeatherService
	PNGRepository          gateways.PNGRepository
	NotificationService    gateways.NotificationService
}

func (u *PollWeathers) Do(ctx context.Context, urlProviders usecases.URLProviders) error {
	users, err := u.UserRepository.FindAll(ctx)
	if err != nil {
		return errors.Wrapf(err, "error while getting users")
	}
	for _, user := range users {
		if err := u.doUser(ctx, user, urlProviders); err != nil {
			return errors.Wrapf(err, "error while processing user %s", user.ID)
		}
	}
	return nil
}

func (u *PollWeathers) doUser(ctx context.Context, user domain.User, urlProviders usecases.URLProviders) error {
	subscriptions, err := u.SubscriptionRepository.FindByUserID(ctx, user.ID)
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
	weathers, err := u.WeatherService.Get(ctx, user.YahooClientID, locations, gateways.NoObservation)
	if err != nil {
		return errors.Wrapf(err, "error while getting weather")
	}
	for i, subscription := range subscriptions {
		weather := weathers[i]
		if err := u.doSubscription(ctx, user, subscription, weather, urlProviders); err != nil {
			return errors.Wrapf(err, "error while processing user %s subscription %s", user.ID, subscription.ID)
		}
	}
	return nil
}

func (u *PollWeathers) doSubscription(ctx context.Context, user domain.User, subscription domain.Subscription, weather domain.Weather, urlProviders usecases.URLProviders) error {
	recipient := subscription.Recipient
	if recipient.IsZero() {
		return nil
	}
	forecast := domain.NewForecast(weather)
	if !forecast.HasAnyTopic() {
		return nil
	}

	weatherChart := chart.Draw(weather)
	var b bytes.Buffer
	if err := png.Encode(&b, weatherChart); err != nil {
		return errors.Wrapf(err, "error while encoding PNG")
	}
	image := domain.NewPNGImage(b.Bytes())
	if err := u.PNGRepository.Save(ctx, image); err != nil {
		return errors.Wrapf(err, "error while saving the image")
	}

	forecastMessage := gateways.ForecastMessage{
		Forecast:   forecast,
		ImageURL:   urlProviders.ImageURLProvider(image.ID),
		WeatherURL: urlProviders.WeatherURLProvider(user.ID, subscription.ID),
	}
	if err := u.NotificationService.SendForecastMessage(ctx, recipient, forecastMessage); err != nil {
		return errors.Wrapf(err, "error while sending the message")
	}
	return nil
}
