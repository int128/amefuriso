package domain

type Recipient struct {
	SlackWebhookURL string
}

func (s Recipient) IsZero() bool {
	return s.SlackWebhookURL == ""
}
