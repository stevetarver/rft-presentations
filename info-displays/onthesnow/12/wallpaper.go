package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/ajstarks/svgo"
)

// Changes:
// - Improve regex. Found that onthesnow significantly changes the embedded link, but they don't have
//   easy access to a developer account - for open source or demo use.
// - Add a generated date

// TODO: Sometimes, a web cam will just disappear - need to handle that case
// TODO: Finding the image link in the web page is taking a bit - can we stop at first match to speed this up
// TODO: we could also split the web page fetch out into go routines... and fill in the image link data
// TODO: create custom cropping and constrained resize to handle image variations
// TODO: Can I add clickable links to the canvas

type ImageData struct {
	label    string
	homepage string
}

type ResortData struct {
	label  string
	images []ImageData
}

const (
	xMargin = 30
	yMargin = 60

	imageYSpace = 30
	colXSpace   = 30
	colCount    = 7
)

var (
	canvasWidth  = 2560
	canvasHeight = 1440
	// Aspen Images are 615wx410h - resize but constrain proportions
	colWidth    = int((canvasWidth - (2 * xMargin)) / colCount)
	imageWidth  = colWidth - imageYSpace
	imageHeight = int((imageWidth * 410) / 615)

	// Choose images from https://www.onthesnow.com/colorado/webcams.html
	resorts = []ResortData{
		ResortData{
			// Powderhorn, Crested Butte
			"South, West",
			[]ImageData{
				ImageData{
					"Powderhorn - Phreshies",
					"https://www.onthesnow.com/colorado/powderhorn/webcams.html",
				},
				ImageData{
					"Crested Butte - snow stake",
					"https://www.onthesnow.com/colorado/crested-butte-mountain-resort/webcams.html",
				},
			},
		},
		ResortData{
			"Aspen Snowmass",
			[]ImageData{
				ImageData{
					"Highlands",
					"https://www.onthesnow.com/colorado/aspen-snowmass/webcams.html?id=3778",
				},
				ImageData{
					"Snowmass - Sam's Knob",
					"https://www.onthesnow.com/colorado/aspen-snowmass/webcams.html",
				},
				ImageData{
					"Snowmass - Base Village",
					"https://www.onthesnow.com/colorado/aspen-snowmass/webcams.html?id=3636",
				},
				ImageData{
					"Sunlight Mountain - snow stake",
					"https://www.onthesnow.com/colorado/sunlight-mountain-resort/webcams.html",
				},
			},
		},
		ResortData{
			/// Steamboat, Howelsen Hill
			"Steamboat",
			[]ImageData{
				ImageData{
					"Steamboat - Christie",
					"https://www.onthesnow.com/colorado/steamboat/webcams.html",
				},
			},
		},
		ResortData{
			"Vail, Beaver Creek",
			[]ImageData{
				ImageData{
					"Blue Sky",
					"https://www.onthesnow.com/colorado/vail/webcams.html?id=769",
				},
				ImageData{
					"Lions Head",
					"https://www.onthesnow.com/colorado/vail/webcams.html?id=7669",
				},
				ImageData{
					"Covered Bridge",
					"https://www.onthesnow.com/colorado/vail/webcams.html?id=7667",
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
				},
				//ImageData{
				//	"Vail Village",
				//	"https://www.onthesnow.com/colorado/vail/webcams.html?id=3376",
				//	"//img5.onthesnow.com/webcams/482/3376/([0-9_-]+)/la.jpg",
				//},
				ImageData{
					"Beaver Creek - Chair 8",
					"https://www.onthesnow.com/colorado/beaver-creek/webcams.html",
				},
			},
		},
		ResortData{
			"Copper Mtn, Breck",
			[]ImageData{
				ImageData{
					"Copper Mountain - Super Bee",
					"https://www.onthesnow.com/colorado/copper-mountain-resort/webcams.html",
				},
				ImageData{
					"Copper Mountain - Resort",
					"https://www.onthesnow.com/colorado/copper-mountain-resort/webcams.html?id=6467",
				},
				ImageData{
					// See also: Base Peak 8
					"Breck - High Alpine Bowls",
					"https://www.onthesnow.com/colorado/breckenridge/webcams.html",
				},
				ImageData{
					"Breck - Springmeier",
					"https://www.onthesnow.com/colorado/breckenridge/webcams.html?id=3247",
				},
				ImageData{
					"Breck - Freeway Terrain Park",
					"https://www.onthesnow.com/colorado/breckenridge/webcams.html?id=738",
				},
			},
		},
		ResortData{
			// Keystone, loveland, Arapahoe
			"Keystone, Arapahoe",
			[]ImageData{
				ImageData{
					// See also: Pall, Mid-Mountain
					"Arapahoe Basin - Snow Plume",
					"https://www.onthesnow.com/colorado/arapahoe-basin-ski-area/webcams.html",
				},
				ImageData{
					"Arapahoe Basin - Base Area",
					"https://www.onthesnow.com/colorado/arapahoe-basin-ski-area/webcams.html?id=5480",
				},
				ImageData{
					// See also: North Peak, Snow Stake
					"Keystone - A51 Terrain Park",
					"https://www.onthesnow.com/colorado/keystone/webcams.html?id=855",
				},
				ImageData{
					"Keystone - River Run",
					"https://www.onthesnow.com/colorado/keystone/webcams.html",
				},
				ImageData{
					"Loveland - snow stake",
					"https://www.onthesnow.com/colorado/loveland/webcams.html",
				},
			},
		},
		ResortData{
			// Winter Park, Granby, Eldora
			"Winter Park",
			[]ImageData{
				ImageData{
					// See also: Sunspot, Resort base, etc.
					"Winter Park - Snoasis",
					"https://www.onthesnow.com/colorado/winter-park-resort/webcams.html?id=5381",
				},
				ImageData{
					"Winter Park - Village",
					"https://www.onthesnow.com/colorado/winter-park-resort/webcams.html?id=5382",
				},
				ImageData{
					"Winter Park - snow stake",
					"https://www.onthesnow.com/colorado/winter-park-resort/webcams.html?id=6300",
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
	// The host and numbers change frequently, but there is exactly one 'slide_cam' div, so
	// we will look for a string like:
	//   <div class="slide_cam"><img src="//img1.onthesnow.com/webcams/20/15648/2018-09-09_1636/la.jpg"
	// and form the url from that.
	pattern := `<div class="slide_cam"><img src="//img\d.onthesnow.com/webcams/([/0-9_-]+)/la.jpg"`
	re := regexp.MustCompile(pattern)
	match := re.FindString(string(html))
	if len(match) == 0 {
		panic(fmt.Sprintf("Could not find embedded i in %s using pattern '%s'", i.homepage, pattern))
	}

	// Split the match by doublequotes for easy access to the URI
	parts := strings.Split(match, "\"")
	if len(parts) < 4 {
		panic(fmt.Sprintf("Can't locate the embedded image link in '%s'", match))
	}

	return fmt.Sprintf("https:%s", parts[3])
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

// parseCmdLineArgs updates canvas width and or height if the user specified them.
// Shows help if args specified but not in correct format
// e.g. go run wallpaper.go -w 640 -h 400
func parseCmdLineArgs() {
	flag.IntVar(&canvasWidth, "w", canvasWidth, "Canvas width in pixels")
	flag.IntVar(&canvasHeight, "h", canvasHeight, "Canvas height in pixels")
	flag.Parse()
}

func main() {
	parseCmdLineArgs()

	canvas.Start(canvasWidth, canvasHeight)
	defer canvas.End()

	background(122, 129, 255)

	x := xMargin
	y := yMargin

	for _, resort := range resorts {
		x, y = resort.draw(x, y)
		for _, image := range resort.images {
			x, y = image.draw(x, y)
		}
		// Prepare to draw next column
		x += imageWidth + colXSpace
		y = yMargin
	}

	t := time.Now()
	canvas.Text(xMargin, canvasHeight-yMargin, t.Format("Monday, Jan 8, 3:04 pm"),
		"fill:black;font-size:24pt;font-family:serif;label-anchor:left")
}
