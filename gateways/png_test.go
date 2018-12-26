package gateways

import (
	"testing"
	"time"

	"github.com/favclip/testerator"
	"github.com/int128/amefurisobot/domain"
	"google.golang.org/appengine/datastore"
)

func TestPNGRepository_RemoveOlderThan(t *testing.T) {
	baseTime := time.Date(2018, 12, 26, 9, 58, 0, 0, time.UTC)
	_, ctx, err := testerator.SpinUp()
	if err != nil {
		t.Fatalf("error while spin-up of test instance: %s", err)
	}
	defer testerator.SpinDown()
	setup := func() {
		keys := []*datastore.Key{
			newPNGKey(ctx, domain.ImageID("ID1")),
			newPNGKey(ctx, domain.ImageID("ID2")),
			newPNGKey(ctx, domain.ImageID("ID3")),
		}
		entities := []*pngEntity{
			{Time: baseTime.Add(-1 * 24 * time.Hour)},
			{Time: baseTime.Add(-2 * 24 * time.Hour)},
			{Time: baseTime.Add(-3 * 24 * time.Hour)},
		}
		if _, err := datastore.PutMulti(ctx, keys, entities); err != nil {
			t.Fatalf("error while saving PNG entities: %s", err)
		}
	}
	getCount := func() int {
		count, err := datastore.NewQuery(pngKind).Count(ctx)
		if err != nil {
			t.Errorf("datastore query returned error: %s", err)
		}
		return count
	}
	var r PNGRepository

	t.Run("0 day ago", func(t *testing.T) {
		setup()
		removedCount, err := r.RemoveOlderThan(ctx, baseTime)
		if err != nil {
			t.Errorf("RemoveOlderThan returned error: %s", err)
		}
		if removedCount != 3 {
			t.Errorf("removedCount wants 3 but %d", removedCount)
		}
		if remain := getCount(); remain != 0 {
			t.Errorf("remain wants 0 but %d", remain)
		}
	})
	t.Run("1 day ago", func(t *testing.T) {
		setup()
		removedCount, err := r.RemoveOlderThan(ctx, baseTime.Add(-1*24*time.Hour))
		if err != nil {
			t.Errorf("RemoveOlderThan returned error: %s", err)
		}
		if removedCount != 2 {
			t.Errorf("removedCount wants 2 but %d", removedCount)
		}
		if remain := getCount(); remain != 1 {
			t.Errorf("remain wants 1 but %d", remain)
		}
	})
	t.Run("3 day ago", func(t *testing.T) {
		setup()
		removedCount, err := r.RemoveOlderThan(ctx, baseTime.Add(-3*24*time.Hour))
		if err != nil {
			t.Errorf("RemoveOlderThan returned error: %s", err)
		}
		if removedCount != 0 {
			t.Errorf("removedCount wants 0 but %d", removedCount)
		}
		if remain := getCount(); remain != 3 {
			t.Errorf("remain wants 3 but %d", remain)
		}
	})
}
