// News API provides a free news source to non-commercial and open-source projects.
// Free use requires an attribution alongside any implementation or display of News API data.
// Preferrably, a hyperlink to https://newsapi.org with the text "Powered by News API"
// The above statement meets our attribution requirements.
package main

// Changes:
// - add news headlines
// - add News API attribution
// - change canvas size to match desktop
// - remove grid

// Code Hygiene:
// TODO: We could define a JSON struct holding only the information we need and unmarshal into that
//       - would simplify referencing the data
// TODO: Could create types for txt* consts to provide for compile time checking
// TODO: Add a news query 'q' query elt to focus results

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
	canvasWidth  = 2560
	canvasHeight = 1440
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

// drawNews prints current news headlines centered on the bottom 1/3 of screen
// Picked https://newsapi.org/ at random, aggregates many feeds, free to devels,
// open source, 15m delay. Requires attribution.
// Doc: https://newsapi.org/docs/get-started
// Api key: 4deee96f178748b6b0ba0a2987a6427c
// API Form: https://newsapi.org/v2/top-headlines?country=[country code]&apiKey=[key]
// Paste this into browser for example data:
//	https://newsapi.org/v2/top-headlines?country=us&apiKey=4deee96f178748b6b0ba0a2987a6427c
func drawNews() {
	url := "https://newsapi.org/v2/top-headlines?country=us&apiKey=4deee96f178748b6b0ba0a2987a6427c"
	info := getJson(url)

	// Draw 10 articles centered at the bottom of the screen
	articleCount := 10
	yOffset := 25
	x := int(canvasWidth / 2)
	y := int(canvasHeight - (articleCount * yOffset) - yMargin)

	// articles is a list of objects, we want:
	//  - source.name
	//  - title
	// default response is 20 items - see 'totalResults' in response
	//
	// NOTE:
	// Originally, I copy/pasted the following definition:
	//   articles := info["articles"].([]map[string]interface{})
	// On first run, go panic'd with:
	//   panic: interface conversion: interface {} is []interface {}, not []map[string]interface {}
	// Cool that it tells you how to change your code...
	articles := info["articles"].([]interface{})
	for index, item := range articles {
		article := item.(map[string]interface{})
		source := article["source"].(map[string]interface{})["name"].(string)
		title := article["title"].(string)
		headline := fmt.Sprintf("%s - %s", title, source)
		text(x, y, txtSizeSm, txtAlignMiddle, headline)
		y += yOffset
		// Stop if we have printed 10 articles
		if index >= articleCount {
			break
		}
	}
}

func main() {
	canvas.Start(canvasWidth, canvasHeight)
	defer canvas.End()

	background(0, 150, 255)

	drawForecast()
	drawDate()
	drawNews()
}
