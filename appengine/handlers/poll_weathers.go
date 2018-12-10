package handlers

import (
	"fmt"
	"net/http"
	"os"

	aeExternals "github.com/int128/amefuriso/appengine/externals"
	"github.com/int128/amefuriso/externals"
	"github.com/int128/amefuriso/usecases"
	"github.com/int128/go-yahoo-weather/weather"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

type PollWeathers struct{}

func (h *PollWeathers) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	httpClient := urlfetch.Client(ctx)
	u := usecases.PollWeathers{
		SubscriptionRepository: &aeExternals.SubscriptionRepository{},
		WeatherService: externals.WeatherService{
			Client: &weather.Client{
				ClientID: os.Getenv("YAHOO_CLIENT_ID"), //TODO
				Client:   httpClient,
			},
		},
		NotificationService: &aeExternals.NotificationService{
			BaseURL:      baseURL(req),
			SlackService: externals.SlackService{Client: httpClient},
		},
	}
	if err := u.Do(ctx); err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 500)
		log.Errorf(ctx, "Error: %s", err)
	}
}
