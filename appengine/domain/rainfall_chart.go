package domain

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/int128/amefuriso/core/domain"
)

type RainfallChart struct {
	ID          RainfallChartID
	Image       []byte
	ContentType string
	Time        time.Time
	Coordinates domain.Coordinates
}

type RainfallChartID string

func (id RainfallChartID) String() string {
	return string(id)
}

func NewRainfallChartID() RainfallChartID {
	var n uint64
	if err := binary.Read(rand.Reader, binary.LittleEndian, &n); err != nil {
		panic(err)
	}
	return RainfallChartID(fmt.Sprintf("%x", n))
}
