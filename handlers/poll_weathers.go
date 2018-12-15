package handlers

import (
	"fmt"
	"github.com/int128/amefurisobot/externals"
	"github.com/int128/amefurisobot/usecases"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
	"net/http"
)

func PollWeathers(w http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	httpClient := urlfetch.Client(ctx)

	u := usecases.PollWeathers{
		UserRepository:         &externals.UserRepository{},
		SubscriptionRepository: &externals.SubscriptionRepository{},
		PNGRepository:          &externals.PNGRepository{},
		PNGURL: func(id string) string {
			return baseURL(req) + "/png?id=" + id
		},
		WeatherService:      &externals.WeatherService{Client: httpClient},
		NotificationService: &externals.NotificationService{Client: httpClient},
	}

	if err := u.Do(ctx); err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 500)
		log.Errorf(ctx, "Error: %s", err)
	}
}
