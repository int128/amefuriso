package main

import (
	"net/http"

	"github.com/int128/amefurisobot/handlers"
	"google.golang.org/appengine"
)

func router() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/png", handlers.PNG)
	m.HandleFunc("/internal/poll-weather", handlers.PollWeathers)
	m.HandleFunc("/internal/setup", handlers.Setup)
	return m
}

func main() {
	http.Handle("/", router())
	appengine.Main()
}
