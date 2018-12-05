package handlers

import (
	"net/http"

	"github.com/int128/amefuriso/appengine/externals"
	"github.com/int128/amefuriso/domain"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type RainfallChart struct{}

func (h *RainfallChart) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing parameter", 400)
		return
	}
	ctx := appengine.NewContext(req)

	var chartRepository externals.RainfallChartRepository
	c, err := chartRepository.FindById(ctx, domain.RainfallChartID(id))
	if err != nil {
		http.Error(w, "server error", 500)
		log.Errorf(ctx, "error while finding chart image: %s", err)
		return
	}
	if c == nil {
		http.Error(w, "not found", 404)
		return
	}
	w.Header().Set("content-type", c.ContentType)
	if _, err := w.Write(c.Image); err != nil {
		log.Errorf(ctx, "error while writing image: %s", err)
	}
}
