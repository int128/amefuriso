package handlers

import (
	"github.com/gorilla/mux"
	"github.com/int128/amefuriso/domain"
	"github.com/int128/amefuriso/domain/chart"
	"github.com/int128/amefuriso/usecases"
	"google.golang.org/appengine/log"
	"image/png"
	"net/http"
)

type GetWeather struct {
	ContextProvider ContextProvider
	Usecase         usecases.IGetWeather
}

func (h *GetWeather) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	v := mux.Vars(req)
	userID, subscriptionID := domain.UserID(v["userID"]), domain.SubscriptionID(v["subscriptionID"])

	ctx := h.ContextProvider(req)
	weather, err := h.Usecase.Do(ctx, userID, subscriptionID)
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
