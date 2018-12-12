package handlers

import (
	"fmt"
	"net/http"
	"os"

	aeExternals "github.com/int128/amefurisobot/appengine/externals"
	"github.com/int128/amefurisobot/appengine/usecases"
	"github.com/int128/amefurisobot/externals"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func PollWeathers(w http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	httpClient := urlfetch.Client(ctx)

	u := usecases.PollWeathers{
		SubscriptionRepository: &aeExternals.SubscriptionRepository{},
		WeatherService: &externals.WeatherService{
			Client:   httpClient,
			ClientID: os.Getenv("YAHOO_CLIENT_ID"), //TODO
		},
		PNGRepository: &aeExternals.PNGRepository{},
		PNGURL: func(id string) string {
			return baseURL(req) + "/png?id=" + id
		},
		SlackService: &externals.SlackService{Client: httpClient},
	}

	if err := u.Do(ctx); err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 500)
		log.Errorf(ctx, "Error: %s", err)
	}
}
