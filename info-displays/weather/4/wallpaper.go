package main

// Changes:
// - move date / time printing into its own func
// - move forecast into its own func
// - defer canvas.End() - let go ensure that the canvas is ended at the right time

import (
	"fmt"
	"os"
	"time"

	"github.com/ajstarks/svgo"
)

const (
	canvasWidth  = 700
	canvasHeight = 700
	xMargin      = 50
	yMargin      = 100

	txtSizeLg      = 48
	txtSizeMd      = 28
	txtSizeSm      = 16
	txtAlignLeft   = "start"
	txtAlignMiddle = "middle"
	txtAlignRight  = "end"
)

var (
	canvas = svg.New(os.Stdout)
)

// background is called after canvas.Start to... fill the background
func background(r, g, b int) {
	canvas.Rect(0, 0, canvasWidth, canvasHeight, canvas.RGB(r, g, b))
}

func text(x, y int, textSize int, textAlign string, msg string) {
	style := fmt.Sprintf("fill:white;font-size:%dpt;font-family:Calibri;text-anchor:%s", textSize, textAlign)
	canvas.Text(x, y, msg, style)
}

// drawDate draws the date / time in the top-right corner of the screen
func drawDate() {
	// The Go format characters are defined here: https://golang.org/src/time/format.go
	// Caveat: you can't just put arbitrary dates / times for formatting, you must use
	//         the ones listed above.
	t := time.Now()
	text(canvasWidth-xMargin, yMargin, txtSizeMd, txtAlignRight, t.Format("Monday, Jan 8"))
	text(canvasWidth-xMargin, yMargin+70, txtSizeLg, txtAlignRight, t.Format("3:04 pm"))
}

// drawForecast prints the current forecast in the upper left
func drawForecast() {
	text(xMargin, yMargin, txtSizeLg, txtAlignLeft, "Forecast")
}

func main() {
	canvas.Start(canvasWidth, canvasHeight)
	defer canvas.End()

	background(0, 150, 255)

	drawForecast()
	drawDate()

	canvas.Grid(0, 0, canvasWidth, canvasHeight, 10, "stroke:black;opacity:0.1")
}
