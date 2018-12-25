package gateways

import (
	"context"
	"fmt"

	"github.com/int128/amefurisobot/domain"
	"github.com/int128/slack"
	"github.com/pkg/errors"
)

type NotificationService struct {
	Client SlackClient
}

func (s *NotificationService) SendForecastMessage(ctx context.Context, recipient domain.Recipient, message domain.ForecastMessage) error {
	text := "予報はありません"
	if start := message.Forecast.RainWillStart; start != nil {
		text = fmt.Sprintf("%s で %s から雨が降る見込みです。",
			message.Forecast.Location.Name,
			start.Time.Format("15:04"))
	}
	if stop := message.Forecast.RainWillStop; stop != nil {
		text = fmt.Sprintf("%s で %s に雨が止む見込みです。",
			message.Forecast.Location.Name,
			stop.Time.Format("15:04"))
	}

	attachment := slack.Attachment{
		ImageURL: message.ImageURL,
		Fallback: text,
		Text:     fmt.Sprintf("%s <最新の予報|%s>", text, message.WeatherURL),
	}
	err := s.Client.Send(ctx, recipient.SlackWebhookURL, slack.Message{
		Attachments: []slack.Attachment{attachment},
	})
	if err != nil {
		return errors.Wrapf(err, "error while sending notification")
	}
	return nil
}
