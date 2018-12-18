package handlers

import (
	"github.com/gorilla/mux"
	"github.com/int128/amefurisobot/domain"
	"github.com/int128/amefurisobot/externals"
	"github.com/int128/amefurisobot/presenters/chart"
	"github.com/int128/amefurisobot/usecases"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
	"image/png"
	"net/http"
)

func GetWeather(w http.ResponseWriter, req *http.Request) {
	v := mux.Vars(req)
	userID, subscriptionID := domain.UserID(v["userID"]), domain.SubscriptionID(v["subscriptionID"])

	ctx := appengine.NewContext(req)
	httpClient := urlfetch.Client(ctx)
	u := usecases.GetWeather{
		UserRepository:         &externals.UserRepository{},
		SubscriptionRepository: &externals.SubscriptionRepository{},
		WeatherService:         &externals.WeatherService{Client: httpClient},
	}
	weather, err := u.Do(ctx, userID, subscriptionID)
	if err != nil {
		if domain.IsErrNoSuchUser(err) || domain.IsErrNoSuchSubscription(err) {
			http.Error(w, "Not Found", 404)
			return
		}
		log.Errorf(ctx, "Error: %s", err)
		http.Error(w, "Server Error", 500)
		return
	}
	// TODO: issue expires header

	img := chart.Draw(*weather)
	if err := png.Encode(w, img); err != nil {
		log.Errorf(ctx, "Error while writing body: %s", err)
	}
}
