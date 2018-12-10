package externals

import (
	"context"
	"time"

	aeDomain "github.com/int128/amefuriso/appengine/domain"
	"github.com/int128/amefuriso/chart"
	"github.com/int128/amefuriso/domain"
	"github.com/int128/amefuriso/externals"
	"github.com/pkg/errors"
	"google.golang.org/appengine/log"
)

type NotificationService struct {
	BaseURL      string
	SlackService externals.SlackService
}

func (e *NotificationService) Send(ctx context.Context, notification domain.Slack, weather domain.Weather) error {
	b, err := chart.DrawPNG(weather)
	if err != nil {
		return errors.Wrapf(err, "error while drawing rainfall chart")
	}
	c := aeDomain.RainfallChart{
		ID:          aeDomain.NewRainfallChartID(),
		Image:       b,
		ContentType: "image/png",
		Time:        time.Now(),
		Coordinates: weather.Location.Coordinates,
	}
	var chartRepository RainfallChartRepository
	if err := chartRepository.Save(ctx, c); err != nil {
		return errors.Wrapf(err, "error while saving the image")
	}

	url := e.BaseURL + "/rainfall?id=" + c.ID.String()
	log.Debugf(ctx, "image is available at %s", url)
	if notification.IsZero() {
		return nil
	}
	message := domain.Message{
		Text:     "Rainfall",
		ImageURL: "", //TODO
	}
	if err := e.SlackService.Send(notification, message); err != nil {
		return errors.Wrapf(err, "error while sending the message")
	}
	return nil
}
