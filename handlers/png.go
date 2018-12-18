package handlers

import (
	"net/http"

	"github.com/int128/amefurisobot/domain"
	"github.com/int128/amefurisobot/externals"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func PNG(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing parameter", 400)
		return
	}
	ctx := appengine.NewContext(req)

	var pngRepository externals.PNGRepository
	image, err := pngRepository.FindById(ctx, domain.ImageID(id))
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
