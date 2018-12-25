package gateways

import (
	"context"

	"github.com/int128/amefurisobot/domain"
	"github.com/int128/slack"
	"github.com/pkg/errors"
)

type NotificationService struct {
	Client SlackClient
}

func (s *NotificationService) Send(ctx context.Context, recipient domain.Recipient, message domain.Message) error {
	err := s.Client.Send(ctx, recipient.SlackWebhookURL, slack.Message{
		Attachments: []slack.Attachment{{
			Text:     message.Text,
			Fallback: message.Text,
			ImageURL: message.ImageURL,
		}},
	})
	if err != nil {
		return errors.Wrapf(err, "error while sending notification")
	}
	return nil
}
