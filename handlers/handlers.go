package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/int128/amefuriso/domain"
	"github.com/int128/amefuriso/usecases/interfaces"
)

type Handlers struct {
	GetWeather    GetWeather
	GetImage      GetImage
	PollWeathers  PollWeathers
	CleanupImages CleanupImages
	Setup         Setup
}

func (h *Handlers) NewRouter() http.Handler {
	m := mux.NewRouter()
	m.Methods("GET").Path("/{userID}/{subscriptionID}/weather").Handler(&h.GetWeather)
	m.Methods("GET").Path("/images/{ID}.png").Handler(&h.GetImage)
	m.Methods("GET").Path("/internal/poll-weather").Handler(&h.PollWeathers)
	m.Methods("GET").Path("/internal/cleanup-images").Handler(&h.CleanupImages)
	m.Methods("GET").Path("/internal/setup").Handler(&h.Setup)
	return m
}

func getImageURLFunc(r *http.Request) usecases.ImageURLProvider {
	return func(id domain.ImageID) string {
		return baseURL(r) + fmt.Sprintf("/images/%s.png", id)
	}
}

func getWeatherURLFunc(r *http.Request) usecases.WeatherURLProvider {
	return func(userID domain.UserID, subscriptionID domain.SubscriptionID) string {
		return baseURL(r) + fmt.Sprintf("/%s/%s/weather", userID, subscriptionID)
	}
}

func baseURL(r *http.Request) string {
	scheme := "http"
	if r.Header.Get("X-AppEngine-Https") == "on" {
		scheme = "https"
	}
	return scheme + "://" + r.Host
}
