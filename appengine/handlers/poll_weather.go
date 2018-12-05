package handlers

import (
	"bytes"
	"context"
	"fmt"
	"image/png"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/int128/amefuriso/adapters"
	"github.com/int128/amefuriso/appengine/externals"
	"github.com/int128/amefuriso/chart"
	"github.com/int128/amefuriso/domain"
	"github.com/int128/go-yahoo-weather/weather"
	"github.com/int128/slack"
	"github.com/pkg/errors"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

type PollWeather struct{}

func (h *PollWeather) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	lat, err := strconv.ParseFloat(q.Get("lat"), 64)
	if err != nil {
		http.Error(w, "invalid parameter: lat", 400)
		return
	}
	lon, err := strconv.ParseFloat(q.Get("lon"), 64)
	if err != nil {
		http.Error(w, "invalid parameter: lon", 400)
		return
	}
	ctx := appengine.NewContext(req)
	if err := h.serve(ctx, req, lat, lon); err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), 500)
		log.Errorf(ctx, "Error: %s", err)
	}
}

func (h *PollWeather) serve(ctx context.Context, req *http.Request, lat, lon float64) error {
	httpClient := urlfetch.Client(ctx)

	client := weather.NewClient(os.Getenv("YAHOO_CLIENT_ID"))
	client.Client = urlfetch.Client(ctx)
	resp, err := client.Get(&weather.Request{
		Coordinates:     []weather.Coordinates{{Latitude: lat, Longitude: lon}},
		IntervalMinutes: 5,
		PastHours:       1,
	})
	if err != nil {
		return errors.Wrapf(err, "error while getting weather")
	}
	weathers, err := adapters.Weathers(resp)
	if err != nil {
		return errors.Wrapf(err, "error while parsing response")
	}
	w := weathers[0]
	if !w.IsRainingNow() && !w.WillRainLater() {
		log.Debugf(ctx, "nothing to notify, exit")
		return nil
	}

	img := chart.Draw(w)
	var b bytes.Buffer
	if err := png.Encode(&b, img); err != nil {
		return errors.Wrapf(err, "error while encoding PNG")
	}
	c := domain.RainfallChart{
		ID:          domain.NewRainfallChartID(),
		Image:       b.Bytes(),
		ContentType: "image/png",
		Time:        time.Now(),
		Coordinates: w.Coordinates,
	}
	var chartRepository externals.RainfallChartRepository
	if err := chartRepository.Save(ctx, c); err != nil {
		return errors.Wrapf(err, "error while saving the image")
	}
	url := baseURL(req) + "/rainfall?id=" + c.ID.String()
	log.Debugf(ctx, "image is available at %s", url)

	notification := domain.Notification{
		Text:     "Rainfall",
		ImageURL: url,
	}
	slackService := externals.SlackService{
		Client: &slack.Client{
			HTTPClient: httpClient,
			WebhookURL: os.Getenv("SLACK_WEBHOOK"),
		},
	}
	if err := slackService.Send(notification); err != nil {
		return errors.Wrapf(err, "error while sending the notification")
	}
	log.Debugf(ctx, "sent notification")
	return nil
}
