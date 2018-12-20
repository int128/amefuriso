package chart

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"

	"github.com/golang/freetype/truetype"
	"github.com/int128/amefurisobot/domain"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	"github.com/pkg/errors"
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

	boxWidth  = 300
	boxTop    = 10
	boxBottom = 10
	boxLeft   = 90
	boxRight  = 10

	labelLeft = 10
	barScale  = 30
	barHeight = 30
)

var (
	observationColor = color.RGBA{R: 0, G: 0, B: 0xff, A: 0xff}
	forecastColor    = color.RGBA{R: 0, G: 0x30, B: 0xff, A: 0xff}
)

func Draw(w domain.Weather) image.Image {
	canvasHeight := boxTop + barHeight*float64(len(w.Observations)+len(w.Forecasts)) + boxBottom
	img := image.NewRGBA(image.Rect(0, 0, canvasWidth, int(canvasHeight)))
	gc := draw2dimg.NewGraphicContext(img)
	gc.SetFontData(regularFont)
	gc.SetFontSize(18)

	drawAxis(gc, canvasHeight-boxTop-boxBottom)

	gc.Translate(0, boxTop)
	drawRainfalls(gc, w.Observations, observationColor)
	gc.Translate(0, barHeight*float64(len(w.Observations)))
	drawRainfalls(gc, w.Forecasts, forecastColor)

	return img
}

func DrawPNG(w domain.Weather) ([]byte, error) {
	img := Draw(w)
	var b bytes.Buffer
	if err := png.Encode(&b, img); err != nil {
		return nil, errors.Wrapf(err, "error while encoding PNG")
	}
	return b.Bytes(), nil
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

func drawRainfalls(gc *draw2dimg.GraphicContext, events []domain.Event, barColor color.Color) {
	for i, event := range events {
		barWidth := float64(event.Rainfall) * barScale
		y := float64(i) * barHeight
		baseline := y + barHeight*3/4

		gc.SetFillColor(color.Black)
		gc.FillStringAt(event.Time.Format("15:04"), labelLeft, baseline)
		if event.Rainfall > 0 {
			x := math.Min(boxLeft+barWidth+10, boxLeft+boxWidth)
			gc.FillStringAt(fmt.Sprintf("%.2f mm/h", event.Rainfall), x, baseline)
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
