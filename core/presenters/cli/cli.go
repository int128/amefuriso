package cli

import (
	"fmt"
	"io"
	"math"
	"strings"
	"time"

	"github.com/int128/amefuriso/core/domain"
	"github.com/pkg/errors"
)

func Draw(w io.Writer, weather domain.Weather) error {
	for _, rainfall := range weather.RainfallObservation {
		t := rainfall.Time.Format("15:04")
		mark := strings.Repeat("🌧 ", int(math.Ceil(float64(rainfall.Amount))))
		_, err := fmt.Fprintf(w, "| %s |         | %5.2f mm/h | %s\n", t, rainfall.Amount, mark)
		if err != nil {
			return errors.Wrapf(err, "error while writing string")
		}
	}
	for _, rainfall := range weather.RainfallForecast {
		t := rainfall.Time.Format("15:04")
		d := -time.Since(rainfall.Time).Minutes()
		mark := strings.Repeat("🌧 ", int(math.Ceil(float64(rainfall.Amount))))
		_, err := fmt.Fprintf(w, "| %s | %+3.0f min | %5.2f mm/h | %s\n", t, d, rainfall.Amount, mark)
		if err != nil {
			return errors.Wrapf(err, "error while writing string")
		}
	}
	return nil
}
