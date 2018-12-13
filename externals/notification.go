package externals

import (
	"net/http"

	"github.com/int128/amefurisobot/domain"
	"github.com/int128/slack"
	"github.com/pkg/errors"
)

type NotificationService struct {
	Client *http.Client
}

func (s *NotificationService) Send(notification domain.Notification, message domain.Message) error {
	c := slack.Client{
		HTTPClient: s.Client,
		WebhookURL: notification.SlackWebhookURL,
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
