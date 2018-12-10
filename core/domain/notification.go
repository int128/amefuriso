package domain

type Slack struct {
	WebhookURL string
	Channel    string
}

func (s Slack) IsZero() bool {
	return s.WebhookURL == "" && s.Channel == ""
}
