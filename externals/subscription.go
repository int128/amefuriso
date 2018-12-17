package externals

import (
	"context"

	"github.com/int128/amefurisobot/domain"
	"github.com/pkg/errors"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

const subscriptionKind = "Subscription"

func newSubscriptionKey(ctx context.Context, userID domain.UserID, id domain.SubscriptionID) *datastore.Key {
	return datastore.NewKey(ctx, subscriptionKind, string(id), 0, newUserKey(ctx, userID))
}

type subscriptionEntity struct {
	LocationName    string
	Coordinates     appengine.GeoPoint
	SlackWebhookURL string
}

type SubscriptionRepository struct{}

func (r *SubscriptionRepository) FindByUserID(ctx context.Context, userID domain.UserID) ([]domain.Subscription, error) {
	q := datastore.NewQuery(subscriptionKind).Ancestor(newUserKey(ctx, userID))
	var entities []subscriptionEntity
	keys, err := q.GetAll(ctx, &entities)
	if err != nil {
		return nil, errors.Wrapf(err, "error while getting entities")
	}
	var ret []domain.Subscription
	for i, e := range entities {
		ret = append(ret, domain.Subscription{
			ID: domain.SubscriptionID(keys[i].StringID()),
			Location: domain.Location{
				Name:        e.LocationName,
				Coordinates: domain.Coordinates{Latitude: e.Coordinates.Lat, Longitude: e.Coordinates.Lng},
			},
			Notification: domain.Notification{
				SlackWebhookURL: e.SlackWebhookURL,
			},
		})
	}
	return ret, nil
}

func (r *SubscriptionRepository) Save(ctx context.Context, userID domain.UserID, subscriptions []domain.Subscription) error {
	if userID == "" {
		return errors.Errorf("userID must not be empty")
	}
	var keys []*datastore.Key
	var entities []*subscriptionEntity
	for _, subscription := range subscriptions {
		if subscription.ID == "" {
			return errors.Errorf("Subscription.ID must not be empty")
		}
		k := newSubscriptionKey(ctx, userID, subscription.ID)
		e := subscriptionEntity{
			LocationName: subscription.Location.Name,
			Coordinates: appengine.GeoPoint{
				Lat: subscription.Location.Coordinates.Latitude,
				Lng: subscription.Location.Coordinates.Longitude,
			},
			SlackWebhookURL: subscription.Notification.SlackWebhookURL,
		}
		keys = append(keys, k)
		entities = append(entities, &e)
	}
	if _, err := datastore.PutMulti(ctx, keys, entities); err != nil {
		return errors.Wrapf(err, "error while saving entities")
	}
	return nil
}
