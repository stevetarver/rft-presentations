package main

// Changes:
// - add formatted date / time
// - right-align on upper right side
// -

import (
	"os"
	"time"

	"github.com/ajstarks/svgo"
)

var (
	canvasWidth  = 500
	canvasHeight = 500
	canvas       = svg.New(os.Stdout)
)

// background is called after canvas.Start to... fill the background
func background(r, g, b int) {
	canvas.Rect(0, 0, canvasWidth, canvasHeight, canvas.RGB(r, g, b))
}

func main() {
	canvas.Start(canvasWidth, canvasHeight)
	background(0, 150, 255)

	canvas.Text(20, 40, "Forecast",
		"fill:white;font-size:36pt;font-family:Calibri;text-anchor:start")

	t := time.Now()

	// Align date / time with right edge of screen
	// The Go format characters are defined here: https://golang.org/src/time/format.go
	// Caveat: you can't just put arbitrary dates / times for formatting, you must use
	//         the ones listed above.
	canvas.Text(canvasWidth, 100, t.Format("Monday January 8"),
		"fill:white;font-size:36pt;font-family:Calibri;text-anchor:end")

	canvas.Text(canvasWidth, 140, t.Format("3:04 pm"),
		"fill:white;font-size:36pt;font-family:Calibri;text-anchor:end")

	canvas.Grid(0, 0, canvasWidth, canvasHeight, 10, "stroke:black;opacity:0.1")
	canvas.End()
}
