package main

// Changes:
// - get live weather data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/ajstarks/svgo"
)

const (
	canvasWidth  = 1000
	canvasHeight = 700
	xMargin      = 50
	yMargin      = 100

	txtSizeLg      = 48
	txtSizeMd      = 28
	txtSizeSm      = 16
	txtSizeExSm    = 12
	txtAlignLeft   = "start"
	txtAlignMiddle = "middle"
	txtAlignRight  = "end"
)

var (
	canvas = svg.New(os.Stdout)
)

// background is called after canvas.Start to... fill the background
func background(r, g, b int) {
	canvas.Rect(0, 0, canvasWidth, canvasHeight, canvas.RGB(r, g, b))
}

func text(x, y int, textSize int, textAlign string, msg string) {
	style := fmt.Sprintf("fill:white;font-size:%dpt;font-family:Calibri;text-anchor:%s", textSize, textAlign)
	canvas.Text(x, y, msg, style)
}

// drawDate draws the date / time in the top-right corner of the screen
func drawDate() {
	// The Go format characters are defined here: https://golang.org/src/time/format.go
	// Caveat: you can't just put arbitrary dates / times for formatting, you must use
	//         the ones listed above.
	t := time.Now()
	text(canvasWidth-xMargin, yMargin, txtSizeMd, txtAlignRight, t.Format("Monday, Jan 8"))
	text(canvasWidth-xMargin, yMargin+70, txtSizeLg, txtAlignRight, t.Format("3:04 pm"))
}

// drawForecast prints the current forecast in the upper left
// Turns out that there is such a demand for weather data, that it is now "difficult" to
// find an open provider - but you can get one with a free API key for personal user.
// I picked Darksky.net cause I really liked their web page: https://darksky.net/.
// Get a trial account here: https://darksky.net/dev
// API: https://api.darksky.net/forecast/[key]/[latitude],[longitude]
// Bondale lat, long: 39.4022,-107.2112
// Paste this in the browser to see what you get back:
//     https://api.darksky.net/forecast/071445d23ee7a1ed7b1392b2aec2726f/39.4022,-107.2112
func drawForecast() {
	// Fetch the current day forecast
	url := "https://api.darksky.net/forecast/071445d23ee7a1ed7b1392b2aec2726f/39.4022,-107.2112"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	// Let Go close this resource when we are through with it
	defer response.Body.Close()

	// Read the whole response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// Parse the json into something we can use
	info := map[string]interface{}{}
	if err := json.Unmarshal(body, &info); err != nil {
		panic(err)
	}

	// Extract the information we want to use
	// The var.(string) is a type assertion - it says we know the variable is not
	// nil and is the paren'd type. We can string those together... Using the example
	// response as a guide.
	currently := info["currently"].(map[string]interface{})
	summary := currently["summary"].(string)
	temp := currently["temperature"].(float64)
	feelsLike := currently["apparentTemperature"].(float64)

	tempStr := fmt.Sprintf("%d°", int(temp))
	feelsLikeStr := fmt.Sprintf("(feels like %d°)", int(feelsLike))

	summaryIcon := info["currently"].(map[string]interface{})["icon"].(string)
	summaryIconUrl := fmt.Sprintf("https://darksky.net/images/weather-icons/%s.png", summaryIcon)

	hourly := info["hourly"].(map[string]interface{})
	hourlySummary := hourly["summary"].(string)

	y := yMargin
	text(xMargin, y, txtSizeMd, txtAlignLeft, summary)
	y += 70
	text(xMargin, y, txtSizeLg, txtAlignLeft, tempStr)

	// Add a link to the summary icon
	canvas.Link(summaryIconUrl, summary)
	canvas.Image(xMargin+80, 100, 100, 100, summaryIconUrl)
	canvas.LinkEnd()

	y += 40
	text(xMargin, y, txtSizeSm, txtAlignLeft, feelsLikeStr)
	y += 20
	text(xMargin, y, txtSizeExSm, txtAlignLeft, hourlySummary)
}

func main() {
	canvas.Start(canvasWidth, canvasHeight)
	defer canvas.End()

	background(0, 150, 255)

	drawForecast()
	drawDate()

	canvas.Grid(0, 0, canvasWidth, canvasHeight, 10, "stroke:black;opacity:0.1")
}
