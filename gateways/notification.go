package gateways

import (
	"context"
	"fmt"

	"github.com/int128/amefuriso/domain"
	"github.com/int128/amefuriso/gateways/interfaces"
	"github.com/int128/amefuriso/infrastructure/interfaces"
	"github.com/int128/slack"
	"github.com/pkg/errors"
)

type NotificationService struct {
	Client infrastructure.SlackClient
}

func (s *NotificationService) SendForecastMessage(ctx context.Context, recipient domain.Recipient, message gateways.ForecastMessage) error {
	text := "予報はありません"
	if start := message.Forecast.RainWillStart; start != nil {
		text = fmt.Sprintf("%s で %s から雨が降る見込みです",
			message.Forecast.Location.Name,
			start.Time.Format("15:04"))
	}
	if stop := message.Forecast.RainWillStop; stop != nil {
		text = fmt.Sprintf("%s で %s に雨が止む見込みです",
			message.Forecast.Location.Name,
			stop.Time.Format("15:04"))
	}

	attachment := slack.Attachment{
		Fallback:  text,
		Title:     text,
		TitleLink: message.WeatherURL,
		ImageURL:  message.ImageURL,
	}
	err := s.Client.Send(ctx, recipient.SlackWebhookURL, slack.Message{
		Attachments: []slack.Attachment{attachment},
	})
	if err != nil {
		return errors.Wrapf(err, "error while sending notification")
	}
	return nil
}
