package externals

import (
	"github.com/go-test/deep"
	"github.com/int128/amefurisobot/domain"
	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
	"testing"
)

func TestSubscriptionRepository_FindByUserID_empty(t *testing.T) {
	ctx, shutdown, err := aetest.NewContext()
	if err != nil {
		t.Fatalf("error while initializing context: %s", err)
	}
	defer shutdown()
	var r SubscriptionRepository

	t.Run("OneEntity", func(t *testing.T) {
		if _, err := datastore.Put(ctx,
			newSubscriptionKey(ctx, "USER2", "SUBSCRIPTION1"),
			&subscriptionEntity{
				LocationName: "Tokyo",
				Coordinates:  appengine.GeoPoint{Lat: 35.663613, Lng: 139.732293},
			}); err != nil {
			t.Fatalf("error while saving subscription: %s", err)
		}

		subscriptions, err := r.FindByUserID(ctx, "USER2")
		if err != nil {
			t.Fatalf("error while finding subscriptions: %s", err)
		}
		want := []domain.Subscription{{
			ID: "SUBSCRIPTION1",
			Location: domain.Location{
				Name:        "Tokyo",
				Coordinates: domain.Coordinates{Latitude: 35.663613, Longitude: 139.732293},
			},
		}}
		if diff := deep.Equal(want, subscriptions); diff != nil {
			t.Error(diff)
		}
	})

	t.Run("MoreEntities", func(t *testing.T) {
		if _, err := datastore.Put(ctx,
			newSubscriptionKey(ctx, "USER3", "SUBSCRIPTION1"),
			&subscriptionEntity{
				LocationName: "Tokyo",
				Coordinates:  appengine.GeoPoint{Lat: 35.663613, Lng: 139.732293},
			}); err != nil {
			t.Fatalf("error while saving subscription: %s", err)
		}
		if _, err := datastore.Put(ctx,
			newSubscriptionKey(ctx, "USER3", "SUBSCRIPTION2"),
			&subscriptionEntity{
				LocationName: "Hakodate",
				Coordinates:  appengine.GeoPoint{Lat: 41.7686738, Lng: 140.728924},
			}); err != nil {
			t.Fatalf("error while saving subscription: %s", err)
		}

		subscriptions, err := r.FindByUserID(ctx, "USER3")
		if err != nil {
			t.Fatalf("error while finding subscriptions: %s", err)
		}
		want := []domain.Subscription{
			{
				ID: "SUBSCRIPTION1",
				Location: domain.Location{
					Name:        "Tokyo",
					Coordinates: domain.Coordinates{Latitude: 35.663613, Longitude: 139.732293},
				},
			}, {
				ID: "SUBSCRIPTION2",
				Location: domain.Location{
					Name:        "Hakodate",
					Coordinates: domain.Coordinates{Latitude: 41.7686738, Longitude: 140.728924},
				},
			},
		}
		if diff := deep.Equal(want, subscriptions); diff != nil {
			t.Error(diff)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		subscriptions, err := r.FindByUserID(ctx, "USER1")
		if err != nil {
			t.Fatalf("error while finding subscriptions: %s", err)
		}
		if len(subscriptions) != 0 {
			t.Errorf("subscriptions wants empty slice but %+v", subscriptions)
		}
	})
}
