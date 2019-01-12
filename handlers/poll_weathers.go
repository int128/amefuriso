package handlers

import (
	"fmt"
	"net/http"

	"github.com/int128/amefuriso/usecases/interfaces"
	"google.golang.org/appengine/log"
)

type PollWeathers struct {
	Usecase usecases.PollWeathers
}

func (h *PollWeathers) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	urlProviders := usecases.URLProviders{
		ImageURLProvider:   getImageURLFunc(req),
		WeatherURLProvider: getWeatherURLFunc(req),
	}
	if err := h.Usecase.Do(ctx, urlProviders); err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 500)
		log.Errorf(ctx, "Error: %s", err)
	}
}
