package handlers

import (
	"net/http"

	"github.com/int128/amefuriso/appengine/externals"
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
	b, err := pngRepository.GetById(ctx, id)
	if err != nil {
		http.Error(w, "server error", 500)
		log.Errorf(ctx, "error while getting image: %s", err)
		return
	}
	if b == nil {
		http.Error(w, "not found", 404)
		return
	}

	w.Header().Set("content-type", "image/png")
	if _, err := w.Write(b); err != nil {
		log.Errorf(ctx, "error while writing image: %s", err)
	}
}
