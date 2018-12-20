package handlers

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type ContextProvider func(*http.Request) context.Context

type Handlers struct {
	GetWeather   GetWeather
	GetPNGImage  GetImage
	PollWeathers PollWeathers
	Setup        Setup
}

func (h *Handlers) NewRouter() http.Handler {
	m := mux.NewRouter()
	m.Path("/{userID}/{subscriptionID}/weather").Methods("GET").Handler(&h.GetWeather)
	m.Path("/png").Methods("GET").Handler(&h.GetPNGImage)
	m.Path("/internal/poll-weather").Methods("GET").Handler(&h.PollWeathers)
	m.Path("/internal/setup").Methods("GET").Handler(&h.Setup)
	return m
}
