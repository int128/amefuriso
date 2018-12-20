package usecases

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/int128/amefurisobot/domain"
	"github.com/int128/amefurisobot/domain/mock_externals"
	"github.com/int128/amefurisobot/domain/testdata"
)

func TestPollWeathers_Do_WithoutNotification(t *testing.T) {
	ctx := context.Background()
	imageURLProvider := func(id domain.ImageID) string {
		return "/" + string(id)
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepository := mock_externals.NewMockUserRepository(ctrl)
	userRepository.EXPECT().FindAll(ctx).Return([]domain.User{
		{
			ID:            "USER1",
			YahooClientID: "CLIENT1",
		},
	}, nil)
	subscriptionRepository := mock_externals.NewMockSubscriptionRepository(ctrl)
	subscriptionRepository.EXPECT().FindByUserID(ctx, domain.UserID("USER1")).Return([]domain.Subscription{
		{
			ID:           "SUBSCRIPTION1",
			Location:     testdata.TokyoLocation,
			Notification: domain.Notification{ /* NONE */ },
		},
	}, nil)
	weatherService := mock_externals.NewMockWeatherService(ctrl)
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

	if err := usecase.Do(ctx, imageURLProvider); err != nil {
		t.Fatalf("Do returned error: %s", err)
	}
}
