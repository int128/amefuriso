package externals

import (
	"testing"

	"github.com/favclip/testerator"
	"github.com/go-test/deep"
	"github.com/int128/amefurisobot/domain"
	"github.com/int128/amefurisobot/domain/testdata"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

var TokyoGeoPoint = appengine.GeoPoint{Lat: 35.663613, Lng: 139.732293}
var HakodateGeoPoint = appengine.GeoPoint{Lat: 41.7686738, Lng: 140.728924}

func TestSubscriptionRepository_FindBySubscriptionID(t *testing.T) {
	_, ctx, err := testerator.SpinUp()
	if err != nil {
		t.Fatalf("error while spin-up of test instance: %s", err)
	}
	defer testerator.SpinDown()
	var r SubscriptionRepository

	t.Run("ExactOne", func(t *testing.T) {
		if _, err := datastore.Put(ctx,
			newSubscriptionKey(ctx, "USER2", "SUBSCRIPTION1"),
			&subscriptionEntity{LocationName: "Tokyo", Coordinates: TokyoGeoPoint}); err != nil {
			t.Fatalf("error while saving subscription: %s", err)
		}

		subscription, err := r.FindBySubscriptionID(ctx, "USER2", "SUBSCRIPTION1")
		if err != nil {
			t.Fatalf("error while finding subscriptions: %s", err)
		}
		want := &domain.Subscription{
			ID:       "SUBSCRIPTION1",
			Location: testdata.TokyoLocation,
		}
		if diff := deep.Equal(want, subscription); diff != nil {
			t.Error(diff)
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		subscription, err := r.FindBySubscriptionID(ctx, "USER2", "SUBSCRIPTION2")
		if subscription != nil {
			t.Errorf("subscription wants nil but %+v", subscription)
		}
		if !domain.IsErrNoSuchSubscription(err) {
			t.Errorf("IsErrNoSuchSubscription(error) wants true but false: err=%+v", err)
		}
	})
}

func TestSubscriptionRepository_FindByUserID_empty(t *testing.T) {
	_, ctx, err := testerator.SpinUp()
	if err != nil {
		t.Fatalf("error while spin-up of test instance: %s", err)
	}
	defer testerator.SpinDown()
	var r SubscriptionRepository

	t.Run("OneEntity", func(t *testing.T) {
		if _, err := datastore.Put(ctx,
			newSubscriptionKey(ctx, "USER2", "SUBSCRIPTION1"),
			&subscriptionEntity{LocationName: "Tokyo", Coordinates: TokyoGeoPoint}); err != nil {
			t.Fatalf("error while saving subscription: %s", err)
		}

		subscriptions, err := r.FindByUserID(ctx, "USER2")
		if err != nil {
			t.Fatalf("error while finding subscriptions: %s", err)
		}
		want := []domain.Subscription{{
			ID:       "SUBSCRIPTION1",
			Location: testdata.TokyoLocation,
		}}
		if diff := deep.Equal(want, subscriptions); diff != nil {
			t.Error(diff)
		}
	})

	t.Run("MoreEntities", func(t *testing.T) {
		if _, err := datastore.Put(ctx,
			newSubscriptionKey(ctx, "USER3", "SUBSCRIPTION1"),
			&subscriptionEntity{LocationName: "Tokyo", Coordinates: TokyoGeoPoint}); err != nil {
			t.Fatalf("error while saving subscription: %s", err)
		}
		if _, err := datastore.Put(ctx,
			newSubscriptionKey(ctx, "USER3", "SUBSCRIPTION2"),
			&subscriptionEntity{LocationName: "Hakodate", Coordinates: HakodateGeoPoint}); err != nil {
			t.Fatalf("error while saving subscription: %s", err)
		}

		subscriptions, err := r.FindByUserID(ctx, "USER3")
		if err != nil {
			t.Fatalf("error while finding subscriptions: %s", err)
		}
		want := []domain.Subscription{
			{
				ID:       "SUBSCRIPTION1",
				Location: testdata.TokyoLocation,
			}, {
				ID:       "SUBSCRIPTION2",
				Location: testdata.HakodateLocation,
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
