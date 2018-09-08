package main

// Changes:
// - change background() to take r,g,b for more flexibility
// - change some variable names
// - set the background color and write some text

import (
	"github.com/ajstarks/svgo"
	"os"
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
		"fill:white;font-size:36pt;font-family:Calibri;text-anchor:left")

	canvas.Grid(0, 0, canvasWidth, canvasHeight, 10, "stroke:black;opacity:0.1")
	canvas.End()
}
