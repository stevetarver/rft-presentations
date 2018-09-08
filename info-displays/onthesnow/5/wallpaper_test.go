package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"
)

var regexTests = []struct {
	filename string
	pattern  string
	result   string
}{
	{"sams_knob.html",
		"//img3.onthesnow.com/webcams/25/1814/([0-9_-]+)/la.jpg",
		"https://img3.onthesnow.com/webcams/25/1814/2018-09-06_1357/la.jpg"},
	{"highlands.html",
		"//img5.onthesnow.com/webcams/25/3778/([0-9_-]+)/la.jpg",
		"https://img5.onthesnow.com/webcams/25/3778/2018-06-22_1537/la.jpg"},
	{"base_village.html",
		"//img1.onthesnow.com/webcams/25/3636/([0-9_-]+)/la.jpg",
		"https://img1.onthesnow.com/webcams/25/3636/2018-09-07_2339/la.jpg"},
}

func TestGetOnTheSnowLinkRegex(t *testing.T) {

	for _, data := range regexTests {
		html, err := ioutil.ReadFile(data.filename)
		if err != nil {
			t.Errorf("Failed to read file: %s", data.filename)
		}

		re := regexp.MustCompile(data.pattern) // will panic on compile failure
		match := re.FindString(string(html))
		if len(match) == 0 {
			t.Errorf("Could not find embedded image in %s using pattern '%s", data.filename, data.pattern)
		}

		imageUrl := fmt.Sprintf("https:%s", match)
		//println(imageUrl)
		if imageUrl != data.result {
			t.Errorf("Wrong image url. expect: '%s', actual: '%s'", data.result, imageUrl)
		}
	}
}
