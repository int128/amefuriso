package chart

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/golang/freetype/truetype"
	"github.com/int128/amefuriso/domain"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	"golang.org/x/image/font/gofont/goregular"
)

var regularFont = draw2d.FontData{Name: "goregular"}

func init() {
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic(fmt.Errorf("error while loading goregular.TTF: %s", err))
	}
	draw2d.RegisterFont(regularFont, font)
}

const (
	canvasWidth = boxLeft + boxWidth + boxRight

	boxWidth  = 500
	boxTop    = 100
	boxBottom = 100
	boxLeft   = 150
	boxRight  = 150

	labelLeft = 75
	barScale  = 30
	barHeight = 30
)

var (
	observationColor = color.RGBA{R: 0, G: 0, B: 0xff, A: 0xff}
	forecastColor    = color.RGBA{R: 0, G: 0x30, B: 0xff, A: 0xff}
)

func Draw(w domain.Weather) image.Image {
	canvasHeight := boxTop + barHeight*float64(len(w.RainfallObservation)+len(w.RainfallForecast)) + boxBottom
	img := image.NewRGBA(image.Rect(0, 0, canvasWidth, int(canvasHeight)))
	gc := draw2dimg.NewGraphicContext(img)
	gc.SetFontData(regularFont)
	gc.SetFontSize(18)

	drawAxis(gc, canvasHeight-boxTop-boxBottom)

	gc.Translate(0, boxTop)
	drawRainfalls(gc, w.RainfallObservation, observationColor)
	gc.Translate(0, barHeight*float64(len(w.RainfallObservation)))
	drawRainfalls(gc, w.RainfallForecast, forecastColor)

	return img
}

func drawAxis(gc *draw2dimg.GraphicContext, boxHeight float64) {
	gc.SetStrokeColor(color.Black)

	gc.MoveTo(boxLeft, boxTop)
	gc.LineTo(boxLeft, boxTop+boxHeight)
	gc.FillStroke()

	gc.MoveTo(boxLeft, boxTop)
	gc.LineTo(canvasWidth-boxRight, boxTop)
	gc.FillStroke()

	gc.MoveTo(boxLeft, boxTop+boxHeight)
	gc.LineTo(canvasWidth-boxRight, boxTop+boxHeight)
	gc.FillStroke()
}

func drawRainfalls(gc *draw2dimg.GraphicContext, rainfalls []domain.Rainfall, barColor color.Color) {
	for i, rainfall := range rainfalls {
		barWidth := float64(rainfall.Amount) * barScale
		y := float64(i) * barHeight
		baseline := y + barHeight*3/4

		gc.SetFillColor(color.Black)
		gc.FillStringAt(rainfall.Time.Format("15:04"), labelLeft, baseline)
		if rainfall.Amount > 0 {
			x := math.Min(boxLeft+barWidth+10, boxLeft+boxWidth)
			gc.FillStringAt(fmt.Sprintf("%.2f mm/h", rainfall.Amount), x, baseline)
		}

		gc.SetFillColor(barColor)
		gc.SetStrokeColor(barColor)
		gc.SetLineWidth(1)
		draw2dkit.Rectangle(gc,
			boxLeft,
			y+barHeight*1/4,
			boxLeft+barWidth,
			y+barHeight*3/4)
		gc.FillStroke()
	}
}
