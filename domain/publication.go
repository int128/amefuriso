package domain

type Publication struct {
	SlackWebhookURL string
}

func (s Publication) IsZero() bool {
	return s.SlackWebhookURL == ""
}
