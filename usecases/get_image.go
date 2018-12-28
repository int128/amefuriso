package usecases

import (
	"context"
	"github.com/int128/amefuriso/domain"
	"github.com/int128/amefuriso/gateways/interfaces"
)

type GetImage struct {
	PNGRepository gateways.PNGRepository
}

func (u *GetImage) Do(ctx context.Context, id domain.ImageID) (*domain.Image, error) {
	return u.PNGRepository.FindById(ctx, domain.ImageID(id))
}
