package usecases

import (
	"context"
	"time"

	"github.com/int128/amefuriso/gateways/interfaces"
	"github.com/pkg/errors"
	"google.golang.org/appengine/log"
)

type CleanupImages struct {
	PNGRepository gateways.PNGRepository
}

func (usecase *CleanupImages) Do(ctx context.Context) error {
	weekAgo := time.Now().Add(-7 * 24 * time.Hour)
	count, err := usecase.PNGRepository.RemoveOlderThan(ctx, weekAgo)
	if err != nil {
		return errors.Wrapf(err, "error while removing images older than %s", weekAgo)
	}
	log.Infof(ctx, "Removed %d images older than %s", count, weekAgo)
	return nil
}
