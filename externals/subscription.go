package externals

import (
	"context"

	"github.com/int128/amefuriso/domain"
)

type SubscriptionRepository interface {
	FindAll(ctx context.Context) ([]domain.Subscription, error)
}
