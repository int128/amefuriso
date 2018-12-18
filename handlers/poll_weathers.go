package handlers

import (
	"context"
	"fmt"
	"github.com/int128/amefurisobot/domain"
	"github.com/int128/amefurisobot/usecases"
	"google.golang.org/appengine/log"
	"net/http"
)

type PollWeathersUsecase interface {
	Do(ctx context.Context, imageURL usecases.ImageURLProvider) error
}

type PollWeathers struct {
	ContextProvider ContextProvider
	Usecase         PollWeathersUsecase
}

func (h *PollWeathers) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := h.ContextProvider(req)
	imageURL := func(id domain.ImageID) string {
		return baseURL(req) + "/png?id=" + string(id)
	}
	if err := h.Usecase.Do(ctx, imageURL); err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 500)
		log.Errorf(ctx, "Error: %s", err)
	}
}
