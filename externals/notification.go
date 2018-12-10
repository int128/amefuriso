package externals

import (
	"context"

	"github.com/int128/amefuriso/domain"
)

type NotificationService interface {
	Send(ctx context.Context, notification domain.Slack, weather domain.Weather) error
}
