package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Image struct {
	ID          ImageID
	ContentType ContentType
	Bytes       []byte
	Time        time.Time
}

type ImageID string

type ContentType string

const PNGContentType = ContentType("image/png")

func NewPNGImage(b []byte) Image {
	return Image{
		ID:          ImageID(uuid.New().String()),
		ContentType: PNGContentType,
		Bytes:       b,
		Time:        time.Now(),
	}
}

type ErrNoSuchImage struct {
	ID ImageID
}

func (e ErrNoSuchImage) Error() string {
	return fmt.Sprintf("No such image %s", e.ID)
}

func IsErrNoSuchImage(err error) bool {
	_, ok := err.(ErrNoSuchImage)
	return ok
}
