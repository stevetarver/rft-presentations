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
// - add constants for spacing
// - add snowmass base village
// - add aspen highlands photo
// - increase the canvas size

const (
	xMargin      = 10
	textYSpacing = 10
	imageYSpace  = 20
	samsKnobUrl  = "https://www.onthesnow.com/colorado/aspen-snowmass/webcams.html"
	// e.g.: //img3.onthesnow.com/webcams/25/1814/2018-09-06_1357/la.jpg
	samsKnobPhotoPattern = `"//img3.onthesnow.com/webcams/25/1814/(.*).jpg" `
	highlandsUrl         = "https://www.onthesnow.com/colorado/aspen-snowmass/webcams.html?id=3778"
	// e.g.: //img5.onthesnow.com/webcams/25/3778/2018-06-22_1537/la.jpg"
	highlandsPhotoPattern = `"//img5.onthesnow.com/webcams/25/3778/(.*).jpg" `
	baseVillageUrl        = "https://www.onthesnow.com/colorado/aspen-snowmass/webcams.html?id=3636"
	// e.g. //img1.onthesnow.com/webcams/25/3636/2018-09-06_1358/la.jpg
	baseVillagePhotoPattern = `"//img1.onthesnow.com/webcams/25/3636/(.*).jpg" `
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
	re, err := regexp.Compile(linkPattern)
	match := re.FindString(string(html))

	return fmt.Sprintf("https:%s", match[1:len(match)-2])
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

func main() {
	canvas.Start(width, height)
	defer canvas.End()

	background(150, 98, 208)

	label := "Sam's Knob"
	currentY := 20
	text(xMargin, currentY, label)
	currentY += textYSpacing
	image(xMargin, currentY, 615/2, 410/2, label, getOnTheSnowLink(samsKnobUrl, samsKnobPhotoPattern))

	label = "Aspen Highlands"
	currentY += 410/2 + imageYSpace
	text(xMargin, currentY, label)
	currentY += textYSpacing
	image(xMargin, currentY, 615/2, 410/2, label, getOnTheSnowLink(highlandsUrl, highlandsPhotoPattern))

	label = "Snowmass Base Village"
	currentY += 410/2 + imageYSpace
	text(xMargin, currentY, label)
	currentY += textYSpacing
	image(xMargin, currentY, 615/2, 410/2, label, getOnTheSnowLink(baseVillageUrl, baseVillagePhotoPattern))

	// Leave the grid in - it will help us position "widgets"
	canvas.Grid(0, 0, width, height, 10, "stroke:black;opacity:0.1")
}
