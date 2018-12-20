package main

import (
	"context"
	"github.com/int128/amefurisobot/externals"
	"github.com/int128/amefurisobot/usecases"
	"net/http"

	"github.com/int128/amefurisobot/handlers"
	"google.golang.org/appengine"
)

func contextProvider(req *http.Request) context.Context {
	return appengine.NewContext(req)
}

func main() {
	h := handlers.Handlers{
		GetWeather: handlers.GetWeather{
			ContextProvider: contextProvider,
			Usecase: &usecases.GetWeather{
				UserRepository:         &externals.UserRepository{},
				SubscriptionRepository: &externals.SubscriptionRepository{},
				WeatherService:         &externals.WeatherService{},
			},
		},
		GetImage: handlers.GetImage{
			ContextProvider: contextProvider,
			Usecase: &usecases.GetImage{
				PNGRepository: &externals.PNGRepository{},
			},
		},
		PollWeathers: handlers.PollWeathers{
			ContextProvider: contextProvider,
			Usecase: &usecases.PollWeathers{
				UserRepository:         &externals.UserRepository{},
				SubscriptionRepository: &externals.SubscriptionRepository{},
				PNGRepository:          &externals.PNGRepository{},
				WeatherService:         &externals.WeatherService{},
				NotificationService:    &externals.NotificationService{},
			},
		},
		Setup: handlers.Setup{
			ContextProvider: contextProvider,
			Usecase: &usecases.Setup{
				SubscriptionRepository: &externals.SubscriptionRepository{},
				UserRepository:         &externals.UserRepository{},
			},
		},
	}
	http.Handle("/", h.NewRouter())
	appengine.Main()
}
