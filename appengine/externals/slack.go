package externals

import (
	"github.com/int128/amefuriso/domain"
	"github.com/int128/slack"
	"github.com/pkg/errors"
)

type SlackService struct {
	Client *slack.Client
}

func (s *SlackService) Send(notification domain.Notification) error {
	err := s.Client.Send(&slack.Message{
		Username:  "amefuriso",
		IconEmoji: ":umbrella_with_rain_drops:",
		Attachments: []slack.Attachment{{
			Text:     notification.Text,
			ImageURL: notification.ImageURL,
		}},
	})
	if err != nil {
		return errors.Wrapf(err, "error while sending a message to Slack")
	}
	return nil
}
