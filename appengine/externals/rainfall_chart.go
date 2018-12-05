package externals

import (
	"context"
	"time"

	"github.com/int128/amefuriso/domain"
	"github.com/pkg/errors"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type rainfallChart struct {
	Image       []byte
	ContentType string
	Time        time.Time
	Coordinates appengine.GeoPoint
}

func newChartKey(ctx context.Context, id domain.RainfallChartID) *datastore.Key {
	return datastore.NewKey(ctx, "RainfallChart", id.String(), 0, nil)
}

type RainfallChartRepository struct{}

func (r *RainfallChartRepository) FindById(ctx context.Context, id domain.RainfallChartID) (*domain.RainfallChart, error) {
	k := newChartKey(ctx, id)
	var e rainfallChart
	err := datastore.Get(ctx, k, &e)
	if err != nil {
		if err == datastore.ErrNoSuchEntity {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "error while getting the RainfallChart entity")
	}
	return &domain.RainfallChart{
		ID:          id,
		Image:       e.Image,
		ContentType: e.ContentType,
		Time:        e.Time,
		Coordinates: domain.Coordinates{Latitude: e.Coordinates.Lat, Longitude: e.Coordinates.Lng},
	}, nil
}

func (r *RainfallChartRepository) Save(ctx context.Context, chart domain.RainfallChart) error {
	k := newChartKey(ctx, chart.ID)
	e := rainfallChart{
		Image:       chart.Image,
		ContentType: chart.ContentType,
		Time:        chart.Time,
		Coordinates: appengine.GeoPoint{Lat: chart.Coordinates.Latitude, Lng: chart.Coordinates.Longitude},
	}
	_, err := datastore.Put(ctx, k, &e)
	if err != nil {
		return errors.Wrapf(err, "error while saving the RainfallChart entity")
	}
	return nil
}
