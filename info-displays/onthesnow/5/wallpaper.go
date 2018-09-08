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
// - refactor repetitive code in main into a function that prints an onTheSnowLabeledImage
// - found regex wasn't resilient, made a test to improve it

const (
	xMargin      = 10
	textYSpacing = 10
	imageYSpace  = 20
	samsKnobUrl  = "https://www.onthesnow.com/colorado/aspen-snowmass/webcams.html"
	// e.g.: //img3.onthesnow.com/webcams/25/1814/2018-09-06_1357/la.jpg
	samsKnobPhotoPattern = "//img3.onthesnow.com/webcams/25/1814/([0-9_-]+)/la.jpg"
	highlandsUrl         = "https://www.onthesnow.com/colorado/aspen-snowmass/webcams.html?id=3778"
	// e.g.: //img5.onthesnow.com/webcams/25/3778/2018-06-22_1537/la.jpg"
	highlandsPhotoPattern = "//img5.onthesnow.com/webcams/25/3778/([0-9_-]+)/la.jpg"
	baseVillageUrl        = "https://www.onthesnow.com/colorado/aspen-snowmass/webcams.html?id=3636"
	// e.g. //img1.onthesnow.com/webcams/25/3636/2018-09-06_1358/la.jpg
	baseVillagePhotoPattern = "//img1.onthesnow.com/webcams/25/3636/([0-9_-]+)/la.jpg"
)

var (
	width  = 500
	height = 1000
	canvas = svg.New(os.Stdout)
)

// background is called after canvas.Start to... fill the background
func background(r, g, b int) {
	canvas.Rect(0, 0, width, height, canvas.RGB(r, g, b))
}

// getOnTheSnowLink extracts a link from a still photo web page from On The Snow website.
func getOnTheSnowLink(url, linkPattern string) string {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	html, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// The page should contain a link to the embedded photo that looks like
	//   //img3.onthesnow.com/webcams/25/1814/2018-09-06_1357/la.jpg
	re := regexp.MustCompile(linkPattern) // will panic on compile failure
	match := re.FindString(string(html))
	if len(match) == 0 {
		panic(fmt.Sprintf("Could not find embedded image in %s using pattern '%s", url, linkPattern))
	}

	return fmt.Sprintf("https:%s", match)
}

// image adds an image element to the canvas
func image(x, y, w, h int, label, url string) {
	canvas.Link(url, label)
	canvas.Image(x, y, w, h, url)
	canvas.LinkEnd()
}

// label writes text to the canvas with a common look and feel
func text(x, y int, label string) {
	canvas.Text(x, y, label,
		"fill:black;font-size:18pt;font-family:serif;label-anchor:left")
}

// onTheSnowLabeledImage writes a label and an image to the global canvas.
// We use onTheSnow's standard image width/height and above functions to extract
// the embedded image url from the main web page.
// We return the new (x,y) writing coords
func onTheSnowLabeledImage(x, y int, label, url, pattern string) (nextX, nextY int) {
	println(x, y)
	width := int(615 / 2)
	height := int(410 / 2)
	currentY := y
	text(x, currentY, label)
	currentY += textYSpacing
	image(x, currentY, width, height, label, getOnTheSnowLink(url, pattern))
	currentY += height + imageYSpace
	return x, currentY
}

func main() {
	canvas.Start(width, height)
	defer canvas.End()

	background(150, 98, 208)

	x, y := onTheSnowLabeledImage(xMargin, 20, "Aspen Highlands", highlandsUrl, highlandsPhotoPattern)
	x, y = onTheSnowLabeledImage(x, y, "Sam's Knob", samsKnobUrl, samsKnobPhotoPattern)
	x, y = onTheSnowLabeledImage(x, y, "Snowmass Base Village", baseVillageUrl, baseVillagePhotoPattern)

	// Leave the grid in - it will help us position "widgets"
	canvas.Grid(0, 0, width, height, 10, "stroke:black;opacity:0.1")
}
