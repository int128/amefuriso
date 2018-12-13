package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/int128/amefurisobot/externals"
	"github.com/int128/amefurisobot/usecases"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func PollWeathers(w http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	httpClient := urlfetch.Client(ctx)

	u := usecases.PollWeathers{
		SubscriptionRepository: &externals.SubscriptionRepository{},
		WeatherService: &externals.WeatherService{
			Client:   httpClient,
			ClientID: os.Getenv("YAHOO_CLIENT_ID"), //TODO
		},
		PNGRepository: &externals.PNGRepository{},
		PNGURL: func(id string) string {
			return baseURL(req) + "/png?id=" + id
		},
		NotificationService: &externals.NotificationService{Client: httpClient},
	}

	if err := u.Do(ctx); err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 500)
		log.Errorf(ctx, "Error: %s", err)
	}
}
