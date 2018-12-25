package infrastructure

import (
	"context"

	"github.com/int128/slack"
	"github.com/pkg/errors"
	"google.golang.org/appengine/urlfetch"
)

type SlackClient struct{}

func (c *SlackClient) Send(ctx context.Context, webhookURL string, message slack.Message) error {
	client := slack.Client{
		HTTPClient: urlfetch.Client(ctx),
		WebhookURL: webhookURL,
	}
	if err := client.Send(&message); err != nil {
		return errors.Wrapf(err, "error while sending Slack message")
	}
	return nil
}
