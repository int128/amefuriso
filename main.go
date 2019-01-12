package main

import (
	"net/http"

	"github.com/int128/amefuriso/gateways"
	"github.com/int128/amefuriso/handlers"
	"github.com/int128/amefuriso/infrastructure"
	"github.com/int128/amefuriso/usecases"
	"google.golang.org/appengine"
)

func main() {
	h := handlers.Handlers{
		GetWeather: handlers.GetWeather{
			Usecase: &usecases.GetWeather{
				UserRepository:         &gateways.UserRepository{},
				SubscriptionRepository: &gateways.SubscriptionRepository{},
				WeatherService: &gateways.WeatherService{
					Client: &infrastructure.WeatherClient{},
				},
			},
		},
		GetImage: handlers.GetImage{
			Usecase: &usecases.GetImage{
				PNGRepository: &gateways.PNGRepository{},
			},
		},
		PollWeathers: handlers.PollWeathers{
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
			Usecase: &usecases.CleanupImages{
				PNGRepository: &gateways.PNGRepository{},
			},
		},
		Setup: handlers.Setup{
			Usecase: &usecases.Setup{
				SubscriptionRepository: &gateways.SubscriptionRepository{},
				UserRepository:         &gateways.UserRepository{},
			},
		},
	}
	http.Handle("/", h.NewRouter())
	appengine.Main()
}
