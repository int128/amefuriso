package externals

import (
	"net/http"

	"github.com/int128/amefuriso/domain"
	"github.com/int128/slack"
	"github.com/pkg/errors"
)

type SlackService struct {
	Client *http.Client
}

func (s *SlackService) Send(destination domain.Slack, message domain.Message) error {
	c := slack.Client{
		HTTPClient: s.Client,
		WebhookURL: destination.WebhookURL,
	}
	err := c.Send(&slack.Message{
		Channel:   destination.Channel,
		Username:  "amefuriso",
		IconEmoji: ":umbrella_with_rain_drops:",
		Attachments: []slack.Attachment{{
			Text:     message.Text,
			ImageURL: message.ImageURL,
		}},
	})
	if err != nil {
		return errors.Wrapf(err, "error while sending a message to Slack")
	}
	return nil
}
