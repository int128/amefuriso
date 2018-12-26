package handlers

import (
	"fmt"
	"net/http"

	"github.com/int128/amefurisobot/usecases"
	"google.golang.org/appengine/log"
)

type CleanupImages struct {
	ContextProvider ContextProvider
	Usecase         usecases.ICleanupImages
}

func (h *CleanupImages) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := h.ContextProvider(req)
	if err := h.Usecase.Do(ctx); err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 500)
		log.Errorf(ctx, "Error: %s", err)
	}
}
