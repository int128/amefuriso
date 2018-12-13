package domain

type Notification struct {
	SlackWebhookURL string
}

func (s Notification) IsZero() bool {
	return s.SlackWebhookURL == ""
}
