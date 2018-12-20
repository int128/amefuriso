package handlers

import (
	"fmt"
	"net/http"

	"github.com/int128/amefurisobot/usecases"
	"google.golang.org/appengine/log"
)

type PollWeathers struct {
	ContextProvider ContextProvider
	Usecase         usecases.IPollWeathers
}

func (h *PollWeathers) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := h.ContextProvider(req)
	urlProviders := usecases.URLProviders{
		ImageURLProvider:   getImageURLFunc(req),
		WeatherURLProvider: getWeatherURLFunc(req),
	}
	if err := h.Usecase.Do(ctx, urlProviders); err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 500)
		log.Errorf(ctx, "Error: %s", err)
	}
}
