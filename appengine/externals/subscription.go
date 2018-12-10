package externals

import (
	"context"

	"github.com/int128/amefuriso/core/domain"
	"github.com/pkg/errors"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

const subscriptionKind = "Subscription"

type subscriptionEntity struct {
	LocationName string
	Coordinates  appengine.GeoPoint
	SlackWebhook string
	SlackChannel string
}

type SubscriptionRepository struct{}

func (r *SubscriptionRepository) FindAll(ctx context.Context) ([]domain.Subscription, error) {
	q := datastore.NewQuery(subscriptionKind).Limit(10)

	var entities []subscriptionEntity
	if _, err := q.GetAll(ctx, &entities); err != nil {
		return nil, errors.Wrapf(err, "error while getting entities")
	}

	var ret []domain.Subscription
	for _, e := range entities {
		ret = append(ret, domain.Subscription{
			Location: domain.Location{
				Name:        e.LocationName,
				Coordinates: domain.Coordinates{Latitude: e.Coordinates.Lat, Longitude: e.Coordinates.Lng},
			},
			Notification: domain.Slack{
				WebhookURL: e.SlackWebhook,
				Channel:    e.SlackChannel,
			},
		})
	}
	return ret, nil
}
