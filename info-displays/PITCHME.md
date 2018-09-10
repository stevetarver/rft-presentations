---?image=info-displays/pitch/images/go-pencil/gopherswim.jpg&size=contain&color=white&position=right

@title[Intro]

<br/>
**_Automating Digital Asset_**
<br/>
**_ Workflows... with Go_**

<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
_An experiment in command_
<br/>
_line tools and drawing in Go_

+++

### ***What we'll do***

* Introduce Go
* Introduce "a" tool to "draw" SVGs
* Code up some SVGs
* Generate some SVGs and PNGs


+++

Follow along with this pitch at: 

<br />

[https://gitpitch.com/stevetarver/rft-presentations/master?p=info-displays](https://gitpitch.com/stevetarver/rft-presentations/master?p=info-displays)

---


### ***What are SVGs***

> Scalable Vector Graphics (SVG) is an XML-based vector image format for two-dimensional graphics with support for interactivity and animation.

+++

### ***Why use them?***

* Scale well compared to raster graphics
* All browsers render them
* Very small
* Export losslessly to `png` at target resolutions for devices that can't render `svg`
* You can generate them programmatically

+++

@snap[north-east]
<img style="border:0px; box-shadow: 0px 0px 0px rgba(0, 0, 0, .0);" height="60px" src="info-displays/pitch/images/pale-violet-square.png">
@snapend

#### SVGs are just XML

```xml
<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" 
    "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg" version="1.1" 
    viewBox="61 61 142 142" width="142pt" height="142pt"
     xmlns:dc="http://purl.org/dc/elements/1.1/">
    <metadata>
        Produced by OmniGraffle 6.6.2
        <dc:date>2018-09-05 19:20:43 +0000</dc:date>
    </metadata>
    <defs/>
    <g stroke="none" stroke-opacity="1" stroke-dasharray="none" 
        fill="none" fill-opacity="1">
        <title>Canvas 1</title>
        <g>
            <title>Layer 1</title>
            <rect x="72" y="72" width="120" height="120" fill="#9662d0"/>
            <rect x="72" y="72" width="120" height="120" stroke="black" 
                stroke-linecap="round" stroke-linejoin="round" stroke-width="1"/>
        </g>
    </g>
</svg>
```

---

### ***When _can_ you use them?***

Classically:

- Web page icons: [Font Awesome](https://fontawesome.com/?utm_source=v4_homepage&utm_medium=display&utm_campaign=fa5_released&utm_content=banner)
- Custom icon sets: [CTL Monocle](https://www.ctl.io/developers/blog/post/monocle-pattern-library)

<br/>

SVGs are great for iconic drawings, but you can do more...

+++

You can create charts:

<img style="border:0px; box-shadow: 0px 0px 0px rgba(0, 0, 0, .0);" height="500px" src="info-displays/pitch/images/starks_examples/stocks.png">

+++

Imagine snapshots from our four mountains:

<img style="border:0px; box-shadow: 0px 0px 0px rgba(0, 0, 0, .0);" height="450px" src="info-displays/pitch/images/starks_examples/jtree_thumbs.png">

+++

And then putting that on a large format TV:

<img style="border:0px; box-shadow: 0px 0px 0px rgba(0, 0, 0, .0);" height="450px" src="info-displays/pitch/images/starks_examples/IMG_4411.png">

---

### ***Design Asset Workflow***

* Concept
* Design
* Draw
* Export
* Publish
* Done

---

### ***How to handle "data"?***

If you want to show "live" data

1. Concept
2. Design
3. Fetch data
4. Draw
5. Export
6. Publish
7. When data changes, go back to #3...

---

### ***Why Go?***

_There are many automation tools... Why go?_

@ul

- Different: Better? New Tricks? Does this prove my toolset is the right choice?
- Runs anywhere: Is my toolset OS specific? Costly? Have license restrictions?
- Has SVG tooling: Nothing we have to write - just use the library.
- Version control: Revision history, merge changes...

@ulend


@snap[south-west]
@fa[git fa-2x]
@snapend

@snap[south]
<img style="border:0px; box-shadow: 0px 0px 0px rgba(0, 0, 0, .0);"  height="50" src="info-displays/pitch/images/Go-Logo/PNG/Go-Logo_Aqua_sm.png">
@snapend

@snap[south-east]
@fa[github fa-2x]
@snapend

---

### ***Install Go***

1. Go to the install page: [https://golang.org/doc/install](https://golang.org/doc/install)
1. Click the "Download Go" button.
1. Run the installer.
1. Check the version on the command line.

```
ᐅ go version
go version go1.11 darwin/amd64
```

##### Troubleshooting

* Restart the shell to pickup paths changes.
* Check paths setup properly.

+++

### ***Create a new project***

```
ᐅ go mod init github.com/stevetarver/rft-presentations/info-displays
go: creating new go.mod: module github.com/stevetarver/rft-presentations/info-displays

ᐅ cat go.mod
module github.com/stevetarver/rft-presentations/info-displays
```

+++

### ***Create your first app***

Create file `hello.go`:

```go
package main

func main() {
    println("Hello baby gophers!")
}
```

+++ 

Run it:

```go
ᐅ go run hello.go
Hello baby gophers!
```

---

### ***`svgo` overview***

[AJ Starks' `svgo`](https://github.com/ajstarks/svgo) is a library that simplifies making XML entries that form an SVG image.

Starks presented "Go for Information Displays" at Gophercon 2018 - _I wanted to see what I could do with it._

<br />

There are similar implementations in other languages... _Should one of them be in your toolbox?_

+++

<img style="align:center; border:0px; box-shadow: 0px 0px 0px rgba(0, 0, 0, .0);" height="550px" src="info-displays/pitch/images/starks_examples/svgdef.png">

---

### ***Hello SVG***

Install the library:

```bash
ᐅ go get github.com/ajstarks/svgo
go: finding github.com/ajstarks/svgo latest
go: downloading github.com/ajstarks/svgo v0.0.0-20180830174826-7338bd80e790

ᐅ cat go.mod
module github.com/stevetarver/rft-presentations/info-displays

require github.com/ajstarks/svgo v0.0.0-20180830174826-7338bd80e790 // indirect
```

+++

### ***Write the code***

```go
package main

import (
	"os"

	"github.com/ajstarks/svgo"
)

func main() {
	width := 600
	height := 400
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height, "fill:rgb(150,98,208)")
	canvas.Text(width, height, "Hello gophers!",
		"fill:white;font-size:60pt;font-family:serif;text-anchor:middle")
	canvas.End()
}
```

+++

### ***Run the code***

```shell
#!/bin/sh -e
#
# Run hello.go
#

go run hello_svg.go > hello_svg.svg

# Convert to png for easy display - gotta pay for this in GitPitch
convert hello_svg.svg hello_svg.png
```

+++

### ***The image***

<img style="border:0px; box-shadow: 0px 0px 0px rgba(0, 0, 0, .0);"  src="info-displays/hello_svg/hello_svg.png">

---

### ***A laptop wallpaper***

Let's get to the interesting part - live data!

Let's make some desktop wallpaper:

* Time, weather, news
* Photo gallery of area ski resort webcams - where should we ski today?

+++

### ***Demo***

---

### Reverse engineering an SVG

Some drawing is better done in a dedicated tool.

1. Develop complicated elements (beziers?) in your favorite tool
1. Identify live data sections
1. Migrate those xml elements to your Go code
1. Integrate live data fetch

---

### Use on website

Since all browsers render SVG, you can easily embed them. You can generate them:

* On your webhost, into your asset directory
* Create a dedicated server that just generates and serves the images

---

### Install as mac wallpaper

A little shell script can make this simple:

```bash
screen_resolution() {
    TEXT=$(system_profiler SPDisplaysDataType | grep Resolution | awk '{$1=$1};1')
    WIDTH=$(echo ${TEXT} | cut -f2 -d' ')
    HEIGHT=$(echo ${TEXT} | cut -f4 -d' ')
    printf "%s ${WIDTH} %s ${HEIGHT}" '-w' '-h'
}

go run wallpaper.go $(screen_resolution) > wallpaper.svg
convert msvg:wallpaper.svg wallpaper.png

osascript -e 'tell application "Finder" to set desktop picture to POSIX file "'${FILE_PATH}'"'
```

---

### Install on appletv

Strategies:

Many large TVs have a store mode for just this type of thing - my LG for example. This retail mode may jack with home user mode.

You can set up a local media server and many tvs can connect to it (webOS)

Use an AppleTV

* Use screen saver: Set to a photo or album and use as a screen saver
* Use Photos: Create a photos album holding a slide show and view that on the TV.
* Use Home Sharing: 
* Get an app to show content from dropbox, google drive, etc.

---

# References

## AJ Starks

* [svgo Github page: https://github.com/ajstarks/svgo](https://github.com/ajstarks/svgo)
* [svgo info displays: https://github.com/ajstarks/go-info-displays](https://github.com/ajstarks/go-info-displays)
* [Deck Github page: https://github.com/ajstarks/deck](https://github.com/ajstarks/deck)

Tons of examples in go-info-displays...

---

### ***Questions?***

- What tools are you using? Contrast...
- What are you doing manually today, that you could automate?

---

# TODO

