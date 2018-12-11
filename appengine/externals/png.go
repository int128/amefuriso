package externals

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/appengine/datastore"
)

const pngKind = "PNG"

type pngEntity struct {
	Image []byte
	Time  time.Time
}

func newPNGKey(ctx context.Context, id string) *datastore.Key {
	return datastore.NewKey(ctx, pngKind, id, 0, nil)
}

type PNGRepository struct{}

func (r *PNGRepository) GetById(ctx context.Context, id string) ([]byte, error) {
	k := newPNGKey(ctx, id)
	var e pngEntity
	err := datastore.Get(ctx, k, &e)
	if err != nil {
		if err == datastore.ErrNoSuchEntity {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "error while getting the entity")
	}
	return e.Image, nil
}

func (r *PNGRepository) Save(ctx context.Context, id string, b []byte) error {
	k := newPNGKey(ctx, id)
	e := pngEntity{
		Image: b,
		Time:  time.Now(),
	}
	_, err := datastore.Put(ctx, k, &e)
	if err != nil {
		return errors.Wrapf(err, "error while saving the entity")
	}
	return nil
}
