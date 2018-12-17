package handlers

import (
	"fmt"
	"github.com/int128/amefurisobot/externals"
	"github.com/int128/amefurisobot/usecases"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func Setup(w http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := usecases.Setup{
		SubscriptionRepository: &externals.SubscriptionRepository{},
		UserRepository:         &externals.UserRepository{},
	}
	user, err := u.Do(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 500)
		log.Errorf(ctx, "Error: %s", err)
	}
	if _, err := fmt.Fprintf(w, "Created user %s", user); err != nil {
		log.Errorf(ctx, "Error while writing body: %s", err)
	}
}
