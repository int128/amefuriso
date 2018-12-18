package main

import (
	"github.com/gorilla/mux"
	"net/http"

	"github.com/int128/amefurisobot/handlers"
	"google.golang.org/appengine"
)

func router() http.Handler {
	m := mux.NewRouter()
	m.Path("/{userID}/{subscriptionID}/weather").Methods("GET").HandlerFunc(handlers.GetWeather)
	m.Path("/png").Methods("GET").HandlerFunc(handlers.PNG)
	m.Path("/internal/poll-weather").Methods("GET").HandlerFunc(handlers.PollWeathers)
	m.Path("/internal/setup").Methods("GET").HandlerFunc(handlers.Setup)
	return m
}

func main() {
	http.Handle("/", router())
	appengine.Main()
}
