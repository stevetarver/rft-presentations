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
// - Noticed that we are passing ImageData around a lot, shouldn't an ImageData know
//   how to draw itself
//   - attach ImageData related functions to ImageData struct
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
	label       string
	homepage    string
	imageLinkRe string
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

// getImageLink extracts a link from a still photo web page from On The Snow website.
func (i ImageData) getImageLink() string {
	response, err := http.Get(i.homepage)
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
	re := regexp.MustCompile(i.imageLinkRe) // will panic on compile failure
	match := re.FindString(string(html))
	if len(match) == 0 {
		panic(fmt.Sprintf("Could not find embedded i in %s using pattern '%s", i.homepage, i.imageLinkRe))
	}

	return fmt.Sprintf("https:%s", match)
}

// drawImage adds an drawImage element to the canvas
func (i ImageData) drawImage(x, y int) (nextX, nextY int) {
	imageUrl := i.getImageLink()
	canvas.Link(imageUrl, i.label)
	canvas.Image(x, y, imageWidth, imageHeight, imageUrl)
	canvas.LinkEnd()
	return x, y + imageHeight + imageYSpace
}

// colHeading writes label to the canvas with a common look and feel
func colHeading(x, y int, label string) (nextX, nextY int) {
	canvas.Text(x, y, label,
		"fill:black;font-size:24pt;font-family:serif;label-anchor:left")
	return x, y + 40
}

// drawLabel writes label to the canvas with a common look and feel
func (i ImageData) drawLabel(x, y int) (nextX, nextY int) {
	canvas.Text(x, y, i.label,
		"fill:black;font-size:18pt;font-family:serif;label-anchor:left")
	return x, y + 20
}

// draw writes a label and an image to the global canvas.
// We use onTheSnow's standard image width/height and above functions to extract
// the embedded image homepage from the main web page.
// We return the new (x,y) writing coords
func (i ImageData) draw(x, y int) (nextX, nextY int) {
	x, y = i.drawLabel(x, y)
	return i.drawImage(x, y)
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
			x, y = image.draw(x, y)
		}
	}

	// Leave the grid in - it will help us position "widgets"
	canvas.Grid(0, 0, canvasWidth, canvasHeight, 10, "stroke:black;opacity:0.1")
}
