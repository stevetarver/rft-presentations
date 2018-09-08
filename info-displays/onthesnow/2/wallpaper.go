package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/ajstarks/svgo"
)

// Changes:
// - add today's photo of Sam's Knob
//   - add a photo label
//   - fetch web page
//   - extract link with regular expression

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
	canvas.Text(10, 20, "Sam's Knob",
		"fill:black;font-size:18pt;font-family:serif;text-anchor:left")

	// Snowmass Sam's Knob still photo webpage
	samsKnobUrl := "https://www.onthesnow.com/colorado/aspen-snowmass/webcams.html"
	response, err := http.Get(samsKnobUrl)
	if err != nil {
		// If this fails for any reason, bail out and tell why
		panic(err)
	}
	// Close the link when go think's it's the right time
	defer response.Body.Close()

	// Read the whole webpage into a variable
	html, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// We can view the html page source
	//fmt.Printf("%s\n", html)

	// And find that the displayed photo will look like
	//   //img3.onthesnow.com/webcams/25/1814/2018-09-06_1357/la.jpg
	// Search for the common part to locate the url for today's image
	re, err := regexp.Compile(`"//img3.onthesnow.com/webcams/25/1814/(.*).jpg" `)
	match := re.FindString(string(html))

	// This is what was found..
	//fmt.Printf("%s\n", match)

	// Strip off the doublequotes and turn it into a url
	//fmt.Printf("https:%s\n", match[1:len(match)-2])

	// Put the image in the SVG
	link := fmt.Sprintf("https:%s", match[1:len(match)-2])
	canvas.Link(link, "Sam's Knob")
	canvas.Image(10, 30, 615/2, 410/2, link)
	canvas.LinkEnd()

	// Leave the grid in - it will help us position "widgets"
	//canvas.Grid(0, 0, width, height, 10, "stroke:black;opacity:0.1")
	canvas.End()
}
