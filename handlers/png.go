package handlers

import (
	"context"
	"net/http"

	"github.com/int128/amefurisobot/domain"
	"google.golang.org/appengine/log"
)

type GetPNGImageUsecase interface {
	Do(ctx context.Context, id domain.ImageID) (*domain.Image, error)
}

type GetPNGImage struct {
	ContextProvider ContextProvider
	Usecase         GetPNGImageUsecase
}

func (h *GetPNGImage) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing parameter", 400)
		return
	}

	ctx := h.ContextProvider(req)
	image, err := h.Usecase.Do(ctx, domain.ImageID(id))
	if err != nil {
		if domain.IsErrNoSuchImage(err) {
			http.Error(w, "not found", 404)
			return
		}
		http.Error(w, "server error", 500)
		log.Errorf(ctx, "error while getting image: %s", err)
		return
	}

	w.Header().Set("content-type", string(image.ContentType))
	if _, err := w.Write(image.Bytes); err != nil {
		log.Errorf(ctx, "error while writing image: %s", err)
	}
}
