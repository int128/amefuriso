package usecases

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/int128/amefurisobot/domain"
	"github.com/int128/amefurisobot/domain/mock_gateways"
	"github.com/int128/amefurisobot/domain/testdata"
)

func TestPollWeathers_Do_WithoutNotification(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	urlProviders := URLProviders{
		ImageURLProvider: func(id domain.ImageID) string {
			return "/" + string(id)
		},
		WeatherURLProvider: func(userID domain.UserID, subscriptionID domain.SubscriptionID) string {
			return "/" + string(userID) + "/" + string(subscriptionID) + "/weather"
		},
	}
	userRepository := mock_gateways.NewMockUserRepository(ctrl)
	userRepository.EXPECT().FindAll(ctx).Return([]domain.User{
		{
			ID:            "USER1",
			YahooClientID: "CLIENT1",
		},
	}, nil)
	subscriptionRepository := mock_gateways.NewMockSubscriptionRepository(ctrl)
	subscriptionRepository.EXPECT().FindByUserID(ctx, domain.UserID("USER1")).Return([]domain.Subscription{
		{
			ID:        "SUBSCRIPTION1",
			Location:  testdata.TokyoLocation,
			Recipient: domain.Recipient{ /* NONE */ },
		},
	}, nil)
	weatherService := mock_gateways.NewMockWeatherService(ctrl)
	weatherService.EXPECT().
		Get(ctx, domain.YahooClientID("CLIENT1"), []domain.Location{testdata.TokyoLocation}, domain.NoObservation).
		Return([]domain.Weather{
			{
				Location:     testdata.TokyoLocation,
				Observations: []domain.Event{},
				Forecasts:    []domain.Event{},
			},
		}, nil)
	usecase := PollWeathers{
		UserRepository:         userRepository,
		SubscriptionRepository: subscriptionRepository,
		WeatherService:         weatherService,
	}

	if err := usecase.Do(ctx, urlProviders); err != nil {
		t.Fatalf("Do returned error: %s", err)
	}
}
