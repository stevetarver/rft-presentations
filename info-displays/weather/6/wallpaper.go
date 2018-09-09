package main

// Changes:
// - clean up forecast code
// - create a getJson function in anticipation of needing it to fetch news

// Code Hygiene:
// TODO: We could define a JSON struct holding only the information we need and unmarshal into that
//       - would simplify referencing the data
// TODO: Could create types for txt* consts to provide for compile time checking

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

	// text styling
	txtSizeLg      = 48
	txtSizeMd      = 28
	txtSizeSm      = 16
	txtSizeExSm    = 12
	txtAlignLeft   = "start"
	txtAlignMiddle = "middle"
	txtAlignRight  = "end"

	// Darksky.net provides a free API key for low volume use: https://darksky.net/dev
	// API form: https://api.darksky.net/forecast/[key]/[latitude],[longitude]
	// Bondale lat, long: 39.4022,-107.2112
	// Example data - paste into browser:
	//     https://api.darksky.net/forecast/071445d23ee7a1ed7b1392b2aec2726f/39.4022,-107.2112
	weatherUrl = "https://api.darksky.net/forecast/071445d23ee7a1ed7b1392b2aec2726f/39.4022,-107.2112"
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

// getJson fetches a url that returns JSON, decodes, and returns a general map.
// Any error is a panic
func getJson(endpoint string) map[string]interface{} {
	// Fetch json response
	response, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	// Ensure we close this resource
	defer response.Body.Close()

	// Read the whole response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// Parse json into a generic map
	result := map[string]interface{}{}
	if err := json.Unmarshal(body, &result); err != nil {
		panic(err)
	}
	return result
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
func drawForecast() {
	// Fetch the current day forecast
	info := getJson(weatherUrl)

	// Extract the information we want to use
	// The var.(type) is a type assertion - it says we know the variable is not
	// nil and is the paren'd type - and will panic if either are false.
	currently := info["currently"].(map[string]interface{})
	summary := currently["summary"].(string)
	temp := currently["temperature"].(float64)
	feelsLike := currently["apparentTemperature"].(float64)

	tempStr := fmt.Sprintf("%d°", int(temp))
	feelsLikeStr := fmt.Sprintf("(feels like %d°)", int(feelsLike))

	summaryIcon := currently["icon"].(string)
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
