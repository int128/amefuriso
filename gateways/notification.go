package gateways

import (
	"context"

	"github.com/int128/amefurisobot/domain"
	"github.com/int128/slack"
	"github.com/pkg/errors"
	"google.golang.org/appengine/urlfetch"
)

type NotificationService struct{}

func (s *NotificationService) Send(ctx context.Context, publication domain.Publication, message domain.Message) error {
	c := slack.Client{
		HTTPClient: urlfetch.Client(ctx),
		WebhookURL: publication.SlackWebhookURL,
	}
	err := c.Send(&slack.Message{
		Attachments: []slack.Attachment{{
			Text:     message.Text,
			Fallback: message.Text,
			ImageURL: message.ImageURL,
		}},
	})
	if err != nil {
		return errors.Wrapf(err, "error while sending Slack notification")
	}
	return nil
}
