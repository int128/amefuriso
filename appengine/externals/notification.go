package externals

import (
	"context"
	"strconv"
	"time"

	"github.com/int128/amefuriso/core/chart"
	"github.com/int128/amefuriso/core/domain"
	"github.com/int128/amefuriso/externals"
	"github.com/pkg/errors"
)

type NotificationService struct {
	BaseURL      string
	SlackService externals.SlackService
}

func (e *NotificationService) Send(ctx context.Context, notification domain.Slack, weather domain.Weather) error {
	if notification.IsZero() {
		return nil
	}

	b, err := chart.DrawPNG(weather)
	if err != nil {
		return errors.Wrapf(err, "error while drawing rainfall chart")
	}
	id := strconv.FormatInt(time.Now().UnixNano(), 36)
	var pngRepository PNGRepository
	if err := pngRepository.Save(ctx, id, b); err != nil {
		return errors.Wrapf(err, "error while saving the image")
	}
	url := e.BaseURL + "/png?id=" + id
	message := domain.Message{
		Text:     weather.Location.Name,
		ImageURL: url,
	}
	if err := e.SlackService.Send(notification, message); err != nil {
		return errors.Wrapf(err, "error while sending the message")
	}
	return nil
}
