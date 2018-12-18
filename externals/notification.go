package externals

import (
	"context"

	"github.com/int128/amefurisobot/domain"
	"github.com/int128/slack"
	"github.com/pkg/errors"
	"google.golang.org/appengine/urlfetch"
)

type NotificationService struct{}

func (s *NotificationService) Send(ctx context.Context, destination domain.Notification, message domain.Message) error {
	c := slack.Client{
		HTTPClient: urlfetch.Client(ctx),
		WebhookURL: destination.SlackWebhookURL,
	}
	err := c.Send(&slack.Message{
		Attachments: []slack.Attachment{{
			Text:     message.Text,
			ImageURL: message.ImageURL,
		}},
	})
	if err != nil {
		return errors.Wrapf(err, "error while sending Slack notification")
	}
	return nil
}
