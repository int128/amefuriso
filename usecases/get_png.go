package usecases

import (
	"context"
	"github.com/int128/amefurisobot/domain"
)

type GetPNGImage struct {
	PNGRepository domain.PNGRepository
}

func (u *GetPNGImage) Do(ctx context.Context, id domain.ImageID) (*domain.Image, error) {
	return u.PNGRepository.FindById(ctx, domain.ImageID(id))
}
