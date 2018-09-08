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
// - Add more images, grouped by area

// TODO: Finding the image link in the web page is taking a bit - can we stop at first match to speed this up
// TODO: we could also split the web page fetch out into go routines... and fill in the image link data
// TODO: create custom cropping and constrained resize to handle image variations
// TODO: make into a proper photo gallery - hsize to fit desktop, but still ratio constrained
// TODO: Can I add clickable links to the canvas

type ImageData struct {
	label       string
	homepage    string
	imageLinkRe string
}

type ResortData struct {
	label  string
	images []ImageData
}

const (
	// My desktop is 2560x1440
	canvasWidth  = 2560
	canvasHeight = 1440
	leftMargin   = 10
	topMargin    = 30

	// Aspen Images are 615x410
	// Other resorts may vary a bit, but these are good enough for now
	imageWidth  = int(615 / 2)
	imageHeight = int(410 / 2)
	imageYSpace = 30
)

var (
	// Choose images from https://www.onthesnow.com/colorado/webcams.html
	resorts = []ResortData{
		ResortData{
			"Powderhorn, Crested Butte",
			[]ImageData{
				ImageData{
					"Powderhorn - Phreshies",
					"https://www.onthesnow.com/colorado/powderhorn/webcams.html",
					"//img4.onthesnow.com/webcams/329/4233/([0-9_-]+)/la.jpg",
				},
				ImageData{
					"Crested Butte - snow stake",
					"https://www.onthesnow.com/colorado/crested-butte-mountain-resort/webcams.html",
					"//img1.onthesnow.com/webcams/120/14628/([0-9_-]+)/la.jpg",
				},
			},
		},
		ResortData{
			"Aspen Snowmass",
			[]ImageData{
				ImageData{
					"Highlands",
					"https://www.onthesnow.com/colorado/aspen-snowmass/webcams.html?id=3778",
					"//img5.onthesnow.com/webcams/25/3778/([0-9_-]+)/la.jpg",
				},
				ImageData{
					"Snowmass - Sam's Knob",
					"https://www.onthesnow.com/colorado/aspen-snowmass/webcams.html",
					"//img3.onthesnow.com/webcams/25/1814/([0-9_-]+)/la.jpg",
				},
				ImageData{
					"Snowmass - Base Village",
					"https://www.onthesnow.com/colorado/aspen-snowmass/webcams.html?id=3636",
					"//img1.onthesnow.com/webcams/25/3636/([0-9_-]+)/la.jpg",
				},
				ImageData{
					"Sunlight Mountain - snow stake",
					"https://www.onthesnow.com/colorado/sunlight-mountain-resort/webcams.html",
					"//img5.onthesnow.com/webcams/445/3166/([0-9_-]+)/la.jpg",
				},
			},
		},
		ResortData{
			"Steamboat, Howelsen Hill",
			[]ImageData{
				ImageData{
					"Steamboat - Christie",
					"https://www.onthesnow.com/colorado/steamboat/webcams.html",
					"//img3.onthesnow.com/webcams/425/3524/([0-9_-]+)/la.jpg",
				},
			},
		},
		ResortData{
			"Vail, Beaver Creek",
			[]ImageData{
				ImageData{
					"Blue Sky",
					"https://www.onthesnow.com/colorado/vail/webcams.html?id=769",
					"//img2.onthesnow.com/webcams/482/769/([0-9_-]+)/la.jpg",
				},
				ImageData{
					"Lion's Head",
					"https://www.onthesnow.com/colorado/vail/webcams.html?id=7669",
					"//img2.onthesnow.com/webcams/482/7669/([0-9_-]+)/la.jpg",
				},
				ImageData{
					"Covered Bridge",
					"https://www.onthesnow.com/colorado/vail/webcams.html?id=7667",
					"//img6.onthesnow.com/webcams/482/7667/([0-9_-]+)/la.jpg",
				},
				//ImageData{
				//	"Buffalo's",
				//	"https://www.onthesnow.com/colorado/vail/webcams.html?id=765",
				//	"//img4.onthesnow.com/webcams/482/765/([0-9_-]+)/la.jpg",
				//},
				//ImageData{
				//	"Two Elk Lodge",
				//	"https://www.onthesnow.com/colorado/vail/webcams.html?id=854",
				//	"//img3.onthesnow.com/webcams/482/854/([0-9_-]+)/la.jpg",
				//},
				ImageData{
					"Wildwood",
					"https://www.onthesnow.com/colorado/vail/webcams.html?id=767",
					"//img6.onthesnow.com/webcams/482/767/([0-9_-]+)/la.jpg",
				},
				//ImageData{
				//	"Vail Village",
				//	"https://www.onthesnow.com/colorado/vail/webcams.html?id=3376",
				//	"//img5.onthesnow.com/webcams/482/3376/([0-9_-]+)/la.jpg",
				//},
				ImageData{
					"Beaver Creek - Chair 8",
					"https://www.onthesnow.com/colorado/beaver-creek/webcams.html",
					"//img1.onthesnow.com/webcams/36/3330/([0-9_-]+)/la.jpg",
				},
			},
		},
		ResortData{
			"Copper Mountain, Breck",
			[]ImageData{
				ImageData{
					"Copper Mountain - Super Bee",
					"https://www.onthesnow.com/colorado/copper-mountain-resort/webcams.html",
					"//img5.onthesnow.com/webcams/113/6466/([0-9_-]+)/la.jpg",
				},
				ImageData{
					"Copper Mountain - Resort",
					"https://www.onthesnow.com/colorado/copper-mountain-resort/webcams.html?id=6467",
					"//img6.onthesnow.com/webcams/113/6467/([0-9_-]+)/la.jpg",
				},
				ImageData{
					// See also: Base Peak 8
					"Breck - High Alpine Bowls",
					"https://www.onthesnow.com/colorado/breckenridge/webcams.html",
					"//img3.onthesnow.com/webcams/77/740/([0-9_-]+)/la.jpg",
				},
				ImageData{
					"Breck - Springmeier",
					"https://www.onthesnow.com/colorado/breckenridge/webcams.html?id=3247",
					"//img2.onthesnow.com/webcams/77/3247/([0-9_-]+)/la.jpg",
				},
				ImageData{
					"Breck - Freeway Terrain Park",
					"https://www.onthesnow.com/colorado/breckenridge/webcams.html?id=738",
					"//img1.onthesnow.com/webcams/77/738/([0-9_-]+)/la.jpg",
				},
			},
		},
		ResortData{
			"Keystone, loveland, Arapahoe",
			[]ImageData{
				ImageData{
					// See also: Pall, Mid-Mountain
					"Arapahoe Basin - Snow Plume",
					"https://www.onthesnow.com/colorado/arapahoe-basin-ski-area/webcams.html",
					"//img3.onthesnow.com/webcams/20/15650/([0-9_-]+)/la.jpg",
				},
				ImageData{
					"Arapahoe Basin - Base Area",
					"https://www.onthesnow.com/colorado/arapahoe-basin-ski-area/webcams.html?id=5480",
					"//img3.onthesnow.com/webcams/20/5480/([0-9_-]+)/la.jpg",
				},
				ImageData{
					// See also: North Peak, Snow Stake
					"Keystone - A51 Terrain Park",
					"https://www.onthesnow.com/colorado/keystone/webcams.html?id=855",
					"//img4.onthesnow.com/webcams/197/855/([0-9_-]+)/la.jpg",
				},
				ImageData{
					"Keystone - River Run",
					"https://www.onthesnow.com/colorado/keystone/webcams.html",
					"//img5.onthesnow.com/webcams/197/2818/([0-9_-]+)/la.jpg",
				},
				ImageData{
					"Loveland - snow stake",
					"https://www.onthesnow.com/colorado/loveland/webcams.html",
					"//img2.onthesnow.com/webcams/220/6355/([0-9_-]+)/la.jpg",
				},
			},
		},
		ResortData{
			"Winter Park, Granby, Eldora",
			[]ImageData{
				ImageData{
					// See also: Sunspot, Resort base, etc.
					"Winter Park - Snoasis",
					"https://www.onthesnow.com/colorado/winter-park-resort/webcams.html?id=5381",
					"//img6.onthesnow.com/webcams/507/5381/([0-9_-]+)/la.jpg",
				},
				ImageData{
					"Winter Park - Resort base",
					"https://www.onthesnow.com/colorado/winter-park-resort/webcams.html?id=3125",
					"//img6.onthesnow.com/webcams/507/3125/([0-9_-]+)/la.jpg",
				},
				ImageData{
					"Winter Park - snow stake",
					"https://www.onthesnow.com/colorado/winter-park-resort/webcams.html?id=6300",
					"//img1.onthesnow.com/webcams/507/6300/([0-9_-]+)/la.jpg",
				},
			},
		},
	}

	canvas = svg.New(os.Stdout)
)

// draw writes a label and an image to the global canvas.
// We use onTheSnow's standard image width/height and above functions to extract
// the embedded image homepage from the main web page.
// We return the new (x,y) writing coords
func (i ImageData) draw(x, y int) (nextX, nextY int) {
	x, y = i.drawLabel(x, y)
	return i.drawImage(x, y)
}

// drawImage adds an drawImage element to the canvas
func (i ImageData) drawImage(x, y int) (nextX, nextY int) {
	imageUrl := i.getImageLink()
	canvas.Link(imageUrl, i.label)
	canvas.Image(x, y, imageWidth, imageHeight, imageUrl)
	canvas.LinkEnd()
	return x, y + imageHeight + imageYSpace
}

// drawLabel writes label to the canvas with a common look and feel
func (i ImageData) drawLabel(x, y int) (nextX, nextY int) {
	canvas.Text(x, y, i.label,
		"fill:black;font-size:18pt;font-family:serif;label-anchor:left")
	return x, y + 20
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

// colHeading writes label to the canvas with a common look and feel
func (r ResortData) draw(x, y int) (nextX, nextY int) {
	canvas.Text(x, y, r.label,
		"fill:black;font-size:24pt;font-family:serif;label-anchor:left")
	return x, y + 40
}

// background is called after canvas.Start to... fill the background
func background(r, g, b int) {
	canvas.Rect(0, 0, canvasWidth, canvasHeight, canvas.RGB(r, g, b))
}

func main() {
	canvas.Start(canvasWidth, canvasHeight)
	defer canvas.End()

	background(150, 98, 208)

	x := leftMargin
	y := topMargin

	for _, resort := range resorts {
		x, y = resort.draw(x, y)
		for _, image := range resort.images {
			x, y = image.draw(x, y)
		}
		x += imageWidth + 30
		y = topMargin
	}

	// Leave the grid in - it will help us position "widgets"
	canvas.Grid(0, 0, canvasWidth, canvasHeight, 10, "stroke:black;opacity:0.1")
}
