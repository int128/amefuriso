package main

import (
	"log"
	"time"

	"github.com/int128/amefurisobot/core/domain"
	"github.com/int128/amefurisobot/core/presenters/chart"
	"github.com/int128/go-yahoo-weather/weather"
	"github.com/llgcode/draw2d/draw2dimg"
)

func main() {
	w := domain.Weather{
		RainfallObservation: []domain.Rainfall{
			{Time: time.Date(2018, 12, 3, 12, 30, 0, 0, weather.Timezone), Amount: 0},
			{Time: time.Date(2018, 12, 3, 12, 35, 0, 0, weather.Timezone), Amount: 0},
			{Time: time.Date(2018, 12, 3, 12, 40, 0, 0, weather.Timezone), Amount: 0},
			{Time: time.Date(2018, 12, 3, 12, 45, 0, 0, weather.Timezone), Amount: 0},
			{Time: time.Date(2018, 12, 3, 12, 50, 0, 0, weather.Timezone), Amount: 0},
			{Time: time.Date(2018, 12, 3, 12, 55, 0, 0, weather.Timezone), Amount: 0},
			{Time: time.Date(2018, 12, 3, 13, 00, 0, 0, weather.Timezone), Amount: 0},
			{Time: time.Date(2018, 12, 3, 13, 05, 0, 0, weather.Timezone), Amount: 1.45},
			{Time: time.Date(2018, 12, 3, 13, 10, 0, 0, weather.Timezone), Amount: 1.55},
			{Time: time.Date(2018, 12, 3, 13, 15, 0, 0, weather.Timezone), Amount: 0},
			{Time: time.Date(2018, 12, 3, 13, 20, 0, 0, weather.Timezone), Amount: 0.75},
			{Time: time.Date(2018, 12, 3, 13, 25, 0, 0, weather.Timezone), Amount: 0},
		},
		RainfallForecast: []domain.Rainfall{
			{Time: time.Date(2018, 12, 3, 13, 30, 0, 0, weather.Timezone), Amount: 0},
			{Time: time.Date(2018, 12, 3, 13, 35, 0, 0, weather.Timezone), Amount: 0},
			{Time: time.Date(2018, 12, 3, 13, 40, 0, 0, weather.Timezone), Amount: 0},
			{Time: time.Date(2018, 12, 3, 13, 45, 0, 0, weather.Timezone), Amount: 0},
			{Time: time.Date(2018, 12, 3, 13, 50, 0, 0, weather.Timezone), Amount: 0},
			{Time: time.Date(2018, 12, 3, 13, 55, 0, 0, weather.Timezone), Amount: 0},
			{Time: time.Date(2018, 12, 3, 14, 00, 0, 0, weather.Timezone), Amount: 1.25},
			{Time: time.Date(2018, 12, 3, 14, 05, 0, 0, weather.Timezone), Amount: 0},
			{Time: time.Date(2018, 12, 3, 14, 10, 0, 0, weather.Timezone), Amount: 0},
			{Time: time.Date(2018, 12, 3, 14, 15, 0, 0, weather.Timezone), Amount: 0.85},
			{Time: time.Date(2018, 12, 3, 14, 20, 0, 0, weather.Timezone), Amount: 0},
			{Time: time.Date(2018, 12, 3, 14, 25, 0, 0, weather.Timezone), Amount: 0},
		},
	}

	img := chart.Draw(w)
	if err := draw2dimg.SaveToPngFile("hello.png", img); err != nil {
		log.Fatalf("Could not save PNG: %s", err)
	}
}
