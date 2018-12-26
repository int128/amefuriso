package main

import (
	"context"
	"net/http"

	"github.com/int128/amefuriso/gateways"
	"github.com/int128/amefuriso/handlers"
	"github.com/int128/amefuriso/infrastructure"
	"github.com/int128/amefuriso/usecases"
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
				UserRepository:         &gateways.UserRepository{},
				SubscriptionRepository: &gateways.SubscriptionRepository{},
				WeatherService: &gateways.WeatherService{
					Client: &infrastructure.WeatherClient{},
				},
			},
		},
		GetImage: handlers.GetImage{
			ContextProvider: contextProvider,
			Usecase: &usecases.GetImage{
				PNGRepository: &gateways.PNGRepository{},
			},
		},
		PollWeathers: handlers.PollWeathers{
			ContextProvider: contextProvider,
			Usecase: &usecases.PollWeathers{
				UserRepository:         &gateways.UserRepository{},
				SubscriptionRepository: &gateways.SubscriptionRepository{},
				PNGRepository:          &gateways.PNGRepository{},
				WeatherService: &gateways.WeatherService{
					Client: &infrastructure.WeatherClient{},
				},
				NotificationService: &gateways.NotificationService{
					Client: &infrastructure.SlackClient{},
				},
			},
		},
		CleanupImages: handlers.CleanupImages{
			ContextProvider: contextProvider,
			Usecase: &usecases.CleanupImages{
				PNGRepository: &gateways.PNGRepository{},
			},
		},
		Setup: handlers.Setup{
			ContextProvider: contextProvider,
			Usecase: &usecases.Setup{
				SubscriptionRepository: &gateways.SubscriptionRepository{},
				UserRepository:         &gateways.UserRepository{},
			},
		},
	}
	http.Handle("/", h.NewRouter())
	appengine.Main()
}
