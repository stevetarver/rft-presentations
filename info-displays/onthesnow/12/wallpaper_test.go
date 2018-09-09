package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"testing"
)

// Playground for testing out a new regex
func TestGetOnTheSnowLinkRegex(t *testing.T) {

	filename := "arapahoe.html"
	html, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("Failed to read file: %s", filename)
	}

	//targetLine := `<div class="slide_cam"><img src="//img1.onthesnow.com/webcams/20/15648/2018-09-09_1636/la.jpg"`
	pattern := `<div class="slide_cam"><img src="//img\d.onthesnow.com/webcams/([/0-9_-]+)/la.jpg"`
	re := regexp.MustCompile(pattern)
	match := re.FindString(string(html))

	println(match)

	parts := strings.Split(match, "\"")
	fmt.Printf("%v\n", parts)
	println(parts[3])

	//if len(match) == 0 {
	//	t.Errorf("Could not find embedded image in %s using pattern '%s", data.filename, data.pattern)
	//}
	//
	//imageUrl := fmt.Sprintf("https:%s", match)
	//println(imageUrl)
	//if imageUrl != data.result {
	//	t.Errorf("Wrong image url. expect: '%s', actual: '%s'", data.result, imageUrl)
	//}
}
