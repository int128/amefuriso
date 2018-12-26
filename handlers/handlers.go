package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/int128/amefurisobot/domain"
	"github.com/int128/amefurisobot/usecases"

	"github.com/gorilla/mux"
)

type ContextProvider func(*http.Request) context.Context

type Handlers struct {
	GetWeather    GetWeather
	GetImage      GetImage
	PollWeathers  PollWeathers
	CleanupImages CleanupImages
	Setup         Setup
}

func (h *Handlers) NewRouter() http.Handler {
	m := mux.NewRouter()
	m.Path("/{userID}/{subscriptionID}/weather").Methods("GET").Handler(&h.GetWeather)
	m.Path("/images/{ID}.png").Methods("GET").Handler(&h.GetImage)
	m.Path("/internal/poll-weather").Methods("GET").Handler(&h.PollWeathers)
	m.Path("/internal/cleanup-images").Methods("GET").Handler(&h.CleanupImages)
	m.Path("/internal/setup").Methods("GET").Handler(&h.Setup)
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
