package handlers

import (
	"fmt"
	"github.com/int128/amefurisobot/usecases"
	"google.golang.org/appengine/log"
	"net/http"
)

type PollWeathers struct {
	ContextProvider ContextProvider
	Usecase         usecases.IPollWeathers
}

func (h *PollWeathers) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := h.ContextProvider(req)
	if err := h.Usecase.Do(ctx, getImageURLFunc(req)); err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 500)
		log.Errorf(ctx, "Error: %s", err)
	}
}
