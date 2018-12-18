package externals

import (
	"context"
	"github.com/int128/amefurisobot/domain"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/appengine/datastore"
)

const pngKind = "PNG"

type pngEntity struct {
	Image []byte
	Time  time.Time
}

func newPNGKey(ctx context.Context, id domain.ImageID) *datastore.Key {
	return datastore.NewKey(ctx, pngKind, string(id), 0, nil)
}

type PNGRepository struct{}

func (r *PNGRepository) FindById(ctx context.Context, id domain.ImageID) (*domain.Image, error) {
	k := newPNGKey(ctx, id)
	var e pngEntity
	err := datastore.Get(ctx, k, &e)
	if err != nil {
		if err == datastore.ErrNoSuchEntity {
			return nil, domain.ErrNoSuchImage{ID: id}
		}
		return nil, errors.Wrapf(err, "error while getting the entity")
	}
	return &domain.Image{
		ID:          id,
		ContentType: domain.PNGContentType,
		Bytes:       e.Image,
		Time:        e.Time,
	}, nil
}

func (r *PNGRepository) Save(ctx context.Context, image domain.Image) error {
	k := newPNGKey(ctx, image.ID)
	e := pngEntity{
		Image: image.Bytes,
		Time:  image.Time,
	}
	_, err := datastore.Put(ctx, k, &e)
	if err != nil {
		return errors.Wrapf(err, "error while saving the entity")
	}
	return nil
}
