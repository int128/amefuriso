package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/int128/amefuriso/core/domain"
	"github.com/int128/amefuriso/core/usecases"
	"github.com/int128/amefuriso/externals"
	"github.com/int128/go-yahoo-weather/weather"
)

type subscriptionRepository domain.Subscription

func (r subscriptionRepository) FindAll(ctx context.Context) ([]domain.Subscription, error) {
	return []domain.Subscription{domain.Subscription(r)}, nil
}

func run(ctx context.Context, subscription domain.Subscription) error {
	u := usecases.PollWeathers{
		SubscriptionRepository: subscriptionRepository(subscription),
		WeatherService: &externals.WeatherService{
			Client: &weather.Client{ClientID: os.Getenv("YAHOO_CLIENT_ID")},
		},
		NotificationService: &notificationService{},
	}
	return u.Do(ctx)
}

func main() {
	var subscription domain.Subscription
	flag.StringVar(&subscription.Location.Name, "name", "Roppongi", "Location Name")
	flag.Float64Var(&subscription.Location.Coordinates.Latitude, "lat", 35.663613, "Latitude")
	flag.Float64Var(&subscription.Location.Coordinates.Longitude, "lon", 139.732293, "Longitude")
	flag.Parse()

	ctx := context.Background()
	if err := run(ctx, subscription); err != nil {
		log.Fatalf("Error: %s", err)
	}
}
