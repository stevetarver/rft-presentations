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
// - Decided this will be a ski resort live photo page using On The Snow Images only:
//   - created a struct to hold standard image data
//   - created a struct to hold resort data
// - Made canvas writing functions return next x,y

const (
	canvasWidth  = 1200
	canvasHeight = 1000
	leftMargin   = 10
	topMargin    = 30

	// Images are 615x410
	imageWidth  = int(615 / 2)
	imageHeight = int(410 / 2)
	imageYSpace = 30
)

type ImageData struct {
	label            string
	url              string
	imageLinkPattern string
}

type ResortData struct {
	label  string
	images []ImageData
}

var (
	resorts = []ResortData{
		ResortData{
			"Aspen Snowmass",
			[]ImageData{
				ImageData{
					"Highlands",
					"https://www.onthesnow.com/colorado/aspen-snowmass/webcams.html?id=3778",
					"//img5.onthesnow.com/webcams/25/3778/([0-9_-]+)/la.jpg",
				},
				ImageData{
					"Sam's Knob",
					"https://www.onthesnow.com/colorado/aspen-snowmass/webcams.html",
					"//img3.onthesnow.com/webcams/25/1814/([0-9_-]+)/la.jpg",
				},
				ImageData{
					"Snowmass Base Village",
					"https://www.onthesnow.com/colorado/aspen-snowmass/webcams.html?id=3636",
					"//img1.onthesnow.com/webcams/25/3636/([0-9_-]+)/la.jpg",
				},
			},
		},
	}

	canvas = svg.New(os.Stdout)
)

// background is called after canvas.Start to... fill the background
func background(r, g, b int) {
	canvas.Rect(0, 0, canvasWidth, canvasHeight, canvas.RGB(r, g, b))
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
func image(x, y int, label, url string) (nextX, nextY int) {
	canvas.Link(url, label)
	canvas.Image(x, y, imageWidth, imageHeight, url)
	canvas.LinkEnd()
	return x, y + imageHeight + imageYSpace
}

// imageLabel writes label to the canvas with a common look and feel
func colHeading(x, y int, label string) (nextX, nextY int) {
	canvas.Text(x, y, label,
		"fill:black;font-size:24pt;font-family:serif;label-anchor:left")
	return x, y + 40
}

// imageLabel writes label to the canvas with a common look and feel
func imageLabel(x, y int, label string) (nextX, nextY int) {
	canvas.Text(x, y, label,
		"fill:black;font-size:18pt;font-family:serif;label-anchor:left")
	return x, y + 20
}

// labeledImage writes a label and an image to the global canvas.
// We use onTheSnow's standard image width/height and above functions to extract
// the embedded image url from the main web page.
// We return the new (x,y) writing coords
func labeledImage(x, y int, label, url, pattern string) (nextX, nextY int) {
	x, y = imageLabel(x, y, label)
	return image(x, y, label, getOnTheSnowLink(url, pattern))
}

func main() {
	canvas.Start(canvasWidth, canvasHeight)
	defer canvas.End()

	background(150, 98, 208)

	x := leftMargin
	y := topMargin

	for _, resort := range resorts {
		x, y = colHeading(x, y, resort.label)
		for _, image := range resort.images {
			x, y = labeledImage(x, y, image.label, image.url, image.imageLinkPattern)
		}
	}

	// Leave the grid in - it will help us position "widgets"
	canvas.Grid(0, 0, canvasWidth, canvasHeight, 10, "stroke:black;opacity:0.1")
}
