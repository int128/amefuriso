package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/int128/amefuriso/domain"
	"github.com/int128/amefuriso/usecases/interfaces"
	"google.golang.org/appengine/log"
)

type GetImage struct {
	Usecase usecases.GetImage
}

func (h *GetImage) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	v := mux.Vars(req)
	id := domain.ImageID(v["ID"])

	ctx := req.Context()
	image, err := h.Usecase.Do(ctx, id)
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
