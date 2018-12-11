package main

import (
	"net/http"

	"github.com/int128/amefuriso/appengine/handlers"
	"google.golang.org/appengine"
)

func router() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/png", handlers.PNG)
	m.HandleFunc("/internal/poll-weather", handlers.PollWeathers)
	return m
}

func main() {
	http.Handle("/", router())
	appengine.Main()
}
