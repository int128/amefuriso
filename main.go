package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"github.com/int128/amefuriso/domain"
	"github.com/int128/amefuriso/externals"
	"github.com/int128/amefuriso/usecases"
	"github.com/int128/go-yahoo-weather/weather"
)

type notificationService struct{}

func (e *notificationService) Send(ctx context.Context, notification domain.Slack, weather domain.Weather) error {
	for _, rainfall := range weather.RainfallObservation {
		t := rainfall.Time.Format("15:04")
		mark := strings.Repeat("🌧 ", int(math.Ceil(float64(rainfall.Amount))))
		fmt.Printf("| %s |         | %5.2f mm/h | %s\n", t, rainfall.Amount, mark)
	}
	for _, rainfall := range weather.RainfallForecast {
		t := rainfall.Time.Format("15:04")
		d := -time.Since(rainfall.Time).Minutes()
		mark := strings.Repeat("🌧 ", int(math.Ceil(float64(rainfall.Amount))))
		fmt.Printf("| %s | %+3.0f min | %5.2f mm/h | %s\n", t, d, rainfall.Amount, mark)
	}
	return nil
}

type subscriptionRepository domain.Subscription

func (r subscriptionRepository) FindAll(ctx context.Context) ([]domain.Subscription, error) {
	return []domain.Subscription{domain.Subscription(r)}, nil
}

func run(ctx context.Context, subscription domain.Subscription) error {
	u := usecases.PollWeathers{
		SubscriptionRepository: subscriptionRepository(subscription),
		WeatherService: externals.WeatherService{
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
