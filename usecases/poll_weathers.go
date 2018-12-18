package usecases

import (
	"context"

	"github.com/int128/amefurisobot/domain"
	"github.com/int128/amefurisobot/presenters/chart"
	"github.com/pkg/errors"
)

type ImageURLProvider func(id domain.ImageID) string

type PollWeathers struct {
	UserRepository         domain.UserRepository
	SubscriptionRepository domain.SubscriptionRepository
	WeatherService         domain.WeatherService
	PNGRepository          domain.PNGRepository
	NotificationService    domain.NotificationService
}

func (u *PollWeathers) Do(ctx context.Context, imageURL ImageURLProvider) error {
	users, err := u.UserRepository.FindAll(ctx)
	if err != nil {
		return errors.Wrapf(err, "error while getting users")
	}
	for _, user := range users {
		if err := u.doUser(ctx, user, imageURL); err != nil {
			return errors.Wrapf(err, "error while polling weathers of user %s", user.ID)
		}
	}
	return nil
}

func (u *PollWeathers) doUser(ctx context.Context, user domain.User, imageURL ImageURLProvider) error {
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
	weathers, err := u.WeatherService.Get(ctx, user.YahooClientID, locations)
	if err != nil {
		return errors.Wrapf(err, "error while getting weather")
	}
	for i, subscription := range subscriptions {
		if subscription.Notification.IsZero() {
			continue
		}
		weather := weathers[i]
		if !weather.IsRainingNow() && !weather.WillRainLater() {
			continue
		}

		b, err := chart.DrawPNG(weather)
		if err != nil {
			return errors.Wrapf(err, "error while drawing rainfall chart")
		}
		image := domain.NewPNGImage(b)
		if err := u.PNGRepository.Save(ctx, image); err != nil {
			return errors.Wrapf(err, "error while saving the image")
		}
		message := domain.Message{
			Text:     weather.Location.Name,
			ImageURL: imageURL(image.ID),
		}
		if err := u.NotificationService.Send(ctx, subscription.Notification, message); err != nil {
			return errors.Wrapf(err, "error while sending the message")
		}
	}
	return nil
}
