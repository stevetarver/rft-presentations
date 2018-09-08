package main

import (
	"os"

	"github.com/ajstarks/svgo"
)

func main() {
	width := 600
	height := 400
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height, "fill:rgb(150,98,208)")
	canvas.Text(width/2, height/2, "Hello gophers!",
		"fill:white;font-size:60pt;font-family:serif;text-anchor:middle")
	canvas.End()
}
