package handlers

import (
	"fmt"
	"net/http"

	"github.com/int128/amefuriso/usecases/interfaces"
	"google.golang.org/appengine/log"
)

type Setup struct {
	ContextProvider ContextProvider
	Usecase         usecases.Setup
}

func (h *Setup) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := h.ContextProvider(req)
	user, err := h.Usecase.Do(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 500)
		log.Errorf(ctx, "Error: %s", err)
	}
	if _, err := fmt.Fprintf(w, "Created user %s", user); err != nil {
		log.Errorf(ctx, "Error while writing body: %s", err)
	}
}
