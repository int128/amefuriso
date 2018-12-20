package usecases

import (
	"context"
	"github.com/int128/amefurisobot/domain"
)

type GetImage struct {
	PNGRepository domain.PNGRepository
}

func (u *GetImage) Do(ctx context.Context, id domain.ImageID) (*domain.Image, error) {
	return u.PNGRepository.FindById(ctx, domain.ImageID(id))
}
