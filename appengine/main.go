package main

import (
	"net/http"

	"github.com/int128/amefuriso/appengine/handlers"
	"google.golang.org/appengine"
)

func router() http.Handler {
	m := http.NewServeMux()
	m.Handle("/rainfall", &handlers.RainfallChart{})
	m.Handle("/internal/poll-weather", &handlers.PollWeather{})
	return m
}

func main() {
	http.Handle("/", router())
	appengine.Main()
}
