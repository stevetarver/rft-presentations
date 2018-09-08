package main

import (
	"os"

	"github.com/ajstarks/svgo"
)

// Changes:
// - 'background' takes r,g,b values
// - add background color
// - add a random image

var (
	width  = 500
	height = 500
	canvas = svg.New(os.Stdout)
)

// background is called after canvas.Start to... fill the background
func background(r, g, b int) {
	canvas.Rect(0, 0, width, height, canvas.RGB(r, g, b))
}

func main() {
	canvas.Start(width, height)

	background(150, 98, 208)

	// An external image has an html "anchor" and positioning info
	link := "https://github.com/ajstarks/go-info-displays/raw/master/images/Dr-Alan-Turing-2956483.jpg"
	canvas.Link(link, "random")
	canvas.Image(10, 10, 100, 100, link)
	canvas.LinkEnd()

	// Leave the grid in - it will help us position "widgets"
	canvas.Grid(0, 0, width, height, 10, "stroke:black;opacity:0.1")
	canvas.End()
}
